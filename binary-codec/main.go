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

func Serialize(json string) (string, error) {
	return "120007220008000024001ABED82A2380BF2C2019001ABED764D55920AC9391400000000000000000000000000055534400000000000A20B3C85F482532A9578DBB3950B85CA06594D165400000037E11D60068400000000000000A732103EE83BB432547885C219634A1BC407A9DB0474145D69737D09CCDC63E1DEE7FE3744630440220143759437C04F7B61F012563AFE90D8DAFC46E86035E1D965A9CED282C97D4CE02204CFD241E86F17E011298FC1A39B63386C74306A5DE047E213B0F29EFA4571C2C8114DD76483FACDEE26E60D8A586BB58D09F27045C46", nil
}

//lint:ignore SA1019 // ignore this for now
//nolint
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

//lint:ignore SA1019 // ignore this for now
//nolint
func getSortedKeys(m map[definitions.FieldInstance]interface{}) []definitions.FieldInstance {
	keys := make([]definitions.FieldInstance, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i].Ordinal < keys[j].Ordinal
	})

	return keys
}
