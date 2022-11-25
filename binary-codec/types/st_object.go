package types

import (
	"fmt"
	"sort"

	"github.com/xyield/xrpl-go/binary-codec/definitions"
	"github.com/xyield/xrpl-go/binary-codec/serdes"
)

type STObject struct{}

func (t *STObject) SerializeJson(json any) ([]byte, error) {
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
		b, err := st.SerializeJson(fimap[v])
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
