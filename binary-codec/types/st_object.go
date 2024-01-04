package types

import (
	"reflect"
	"sort"

	"github.com/CreatureDev/xrpl-go/binary-codec/definitions"
	"github.com/CreatureDev/xrpl-go/binary-codec/serdes"
	"github.com/mitchellh/mapstructure"
)

// FieldMutation allows values to mutated before being serialized.
type FieldMutation func(any) any

// Zero returns a FieldMutation that sets the value to its zero value.
func Zero() FieldMutation {
	return func(v any) any {
		return reflect.Zero(reflect.TypeOf(v)).Interface()
	}
}

// STObject represents a map of serialized field instances, where each key is a field name
// and the associated value is the field's value. This structure allows us to represent nested
// and complex structures of the Ripple protocol.
type STObject struct {
	OnlySigning bool
	Mutations   map[string]FieldMutation
}

// FromJson converts a JSON object into a serialized byte slice.
// It works by converting the JSON object into a map of field instances (which include the field definition
// and value), and then serializing each field instance.
// This method returns an error if the JSON input is not a valid object.
func (t *STObject) FromJson(json any) ([]byte, error) {
	s := serdes.NewSerializer()
	var m map[string]any
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: "json", Result: &m, Squash: true})
	if err != nil {
		return nil, err
	}
	err = dec.Decode(json)
	if err != nil {
		return nil, err
	}

	for k, v := range t.Mutations {
		if _, ok := m[k]; ok {
			m[k] = v(m[k])
		}
	}

	fimap, err := createFieldInstanceMapFromJson(m)

	if err != nil {
		return nil, err
	}

	sk := getSortedKeys(fimap)

	for _, v := range sk {
		if checkZero(fimap[v]) && !containsKey(t.Mutations, v.FieldName) {
			continue
		}

		if !v.IsSerialized {
			continue
		}

		if t.OnlySigning && !v.IsSigningField {
			continue
		}

		st := GetSerializedType(v.Type)
		b, err := st.FromJson(fimap[v])
		if err != nil {
			return nil, err
		}
		err = s.WriteFieldAndValue(v, b)
		if err != nil {
			return nil, err
		}
	}
	return s.GetSink(), nil
}

// ToJson takes a BinaryParser and optional parameters, and converts the serialized byte data
// back to a JSON value. It will continue parsing until it encounters an end marker for an object
// or an array, or until the parser has no more data.
func (t *STObject) ToJson(p *serdes.BinaryParser, opts ...int) (any, error) {
	m := make(map[string]any)

	for p.HasMore() {

		fi, err := p.ReadField()
		if err != nil {
			return nil, err
		}

		if fi.FieldName == "ObjectEndMarker" || fi.FieldName == "ArrayEndMarker" {
			break
		}

		st := GetSerializedType(fi.Type)

		var res any
		if fi.IsVLEncoded {
			vlen, err := p.ReadVariableLength()
			if err != nil {
				return nil, err
			}
			res, err = st.ToJson(p, vlen)
			if err != nil {
				return nil, err
			}

		} else {
			res, err = st.ToJson(p)
			if err != nil {
				return nil, err
			}
		}
		res, err = enumToStr(fi.FieldName, res)
		if err != nil {
			return nil, err
		}

		m[fi.FieldName] = res
	}
	return m, nil
}

// type fieldInstanceMap map[definitions.FieldInstance]any

// func (f fieldInstanceMap) addFieldInstanceFromMap(rv reflect.Value) error {
// 	if rv.Kind() != reflect.Map {
// 		return errors.New("not of type map")
// 	}

// 	iter := rv.MapRange()
// 	for iter.Next() {
// 		fi, err := definitions.Get().GetFieldInstanceByFieldName(iter.Key().String())
// 		if err != nil {
// 			return err
// 		}
// 		f[*fi] = iter.Value().Interface()
// 	}
// 	return nil
// }

// func (f fieldInstanceMap) addFieldInstanceFromStruct(rv reflect.Value) error {
// 	if rv.Kind() != reflect.Struct {
// 		return errors.New("not of type struct")
// 	}
// 	for i := 0; i < rv.NumField(); i++ {

// 		rvField := rv.Type().Field(i)
// 		if rvField.Name == "BaseTx" {
// 			continue
// 		}
// 		if rv.Field(i).IsZero() {
// 			continue
// 		}
// 		fi, err := definitions.Get().GetFieldInstanceByFieldName(rvField.Name)

// 		if err != nil {
// 			return err
// 		}

// 		f[*fi] = rv.Field(i).Interface()
// 	}
// 	return nil
// }

// nolint
// createFieldInstanceMapFromJson creates a map of field instances from a JSON object.
// Each key-value pair in the JSON object is converted into a field instance, where the key
// represents the field name and the value is the field's value.
//
//lint:ignore U1000 // ignore this for now
func createFieldInstanceMapFromJson(json map[string]any) (map[definitions.FieldInstance]any, error) {
	m := make(map[definitions.FieldInstance]any, len(json))

	for k, v := range json {
		fi, err := definitions.Get().GetFieldInstanceByFieldName(k)

		if err != nil {
			return nil, err
		}

		m[*fi] = v
	}
	return m, nil
	// rv := reflect.ValueOf(json)
	// m := make(fieldInstanceMap)
	// switch rv.Kind() {
	// case reflect.Map:
	// 	err := m.addFieldInstanceFromMap(rv)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// case reflect.Struct:
	// 	err := m.addFieldInstanceFromStruct(rv)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// default:
	// 	return nil, errors.New("not a valid json node")
	// }
	// return m, nil
}

// nolint
//
// getSortedKeys is a helper function to sort the keys of a map of field instances based on
// their ordinal values. This is used to ensure that the fields are serialized in the
// correct order.
//
//lint:ignore U1000 // ignore this for now
func getSortedKeys(m map[definitions.FieldInstance]any) []definitions.FieldInstance {
	keys := make([]definitions.FieldInstance, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i].Ordinal < keys[j].Ordinal
	})
	return keys
}

// enumToStr is a helper function that takes a field name and its associated value,
// and returns a string representation of the value if the field is an enumerated type
// (i.e., TransactionType, TransactionResult, LedgerEntryType).
// If the field is not an enumerated type, the original value is returned.
func enumToStr(fieldType string, value any) (any, error) {
	switch fieldType {
	case "TransactionType":
		return definitions.Get().GetTransactionTypeNameByTransactionTypeCode(int32(value.(int)))
	case "TransactionResult":
		return definitions.Get().GetTransactionResultNameByTransactionResultTypeCode(int32(value.(int)))
	case "LedgerEntryType":
		return definitions.Get().GetLedgerEntryTypeNameByLedgerEntryTypeCode(int32(value.(int)))
	default:
		return value, nil
	}
}

// check for zero value
func checkZero(v any) bool {
	rv := reflect.ValueOf(v)
	return rv.IsZero()
}

func containsKey[T any](m map[string]T, key string) bool {
	_, ok := m[key]
	return ok
}
