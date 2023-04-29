package types

import (
	"fmt"
	"sort"

	"github.com/xyield/xrpl-go/binary-codec/definitions"
	"github.com/xyield/xrpl-go/binary-codec/serdes"
)

type STObject struct{}

func (t *STObject) FromJson(json any) ([]byte, error) {
	s := serdes.NewSerializer()
	if _, ok := json.(map[string]any); !ok {
		return nil, fmt.Errorf("not a valid json node")
	}
	fimap, err := createFieldInstanceMapFromJson(json.(map[string]any))

	if err != nil {
		return nil, err
	}

	sk := getSortedKeys(fimap)

	for _, v := range sk {
		if !v.IsSerialized {
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

func (t *STObject) FromParser(p *serdes.BinaryParser, opts ...int) (any, error) {
	m := make(map[string]any)
	for p.HasMore() {
		f, err := p.ReadField()
		if err != nil {
			return nil, err
		}
		st := GetSerializedType(f.Type)
		var res any
		if f.IsVLEncoded {
			size, err := p.ReadVariableLength()
			if err != nil {
				return nil, err
			}
			res, err = st.FromParser(p, size)
			if err != nil {
				return nil, err
			}
		} else {
			res, err = st.FromParser(p)
		}
		if err != nil {
			return nil, err
		}
		res, err = enumToStr(f.FieldName, res)
		if err != nil {
			return nil, err
		}
		m[f.FieldName] = res
	}
	// fmt.Println(f)
	return m, nil
}

// nolint
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
}

// nolint
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
