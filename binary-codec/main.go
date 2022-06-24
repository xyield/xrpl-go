package binarycodec

import (
	"sort"

	"github.com/xyield/xrpl-go/binary-codec/definitions"
)

func Encode(json map[string]interface{}) (string, error) {

	// Loop through and create map of map[FieldInstance]interface{}
	// Sort by Ordinal
	// Start serializing
	return "", nil
}

func createFieldInstanceMapFromJson(json map[string]interface{}) (map[definitions.FieldInstance]interface{}, error) {

	m := make(map[definitions.FieldInstance]interface{}, len(json))

	for k, v := range json {
		fi, err := definitions.Get().GetFieldInstanceByFieldName(k)

		if err != nil {
			return nil, err
		}

		m[*fi] = v

	}

	return m, nil
}

func getSortedKeys(m map[definitions.FieldInstance]interface{}) []definitions.FieldInstance {
	keys := make([]definitions.FieldInstance, 0, len(m))

	for k, _ := range m {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i].Ordinal < keys[j].Ordinal
	})

	return keys
}
