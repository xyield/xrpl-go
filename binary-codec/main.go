package binarycodec

import (
	"encoding/hex"
	"strings"

	"github.com/xyield/xrpl-go/binary-codec/types"
)

// Encode: encodes a transaction or other object from json to the canonical binary format as a hex string.
func Encode(json map[string]any) (string, error) {

	st := &types.STObject{}
	b, err := st.SerializeJson(json)
	if err != nil {
		return "", err
	}

	return strings.ToUpper(hex.EncodeToString(b)), nil
}

// nolint
//
//lint:ignore U1000 // ignore this for now
// func createFieldInstanceMapFromJson(json map[string]any) (map[definitions.FieldInstance]any, error) {

// 	m := make(map[definitions.FieldInstance]any, len(json))

// 	for k, v := range json {
// 		fi, err := definitions.Get().GetFieldInstanceByFieldName(k)

// 		if err != nil {
// 			return nil, err
// 		}

// 		m[*fi] = v

// 	}

// 	return m, nil
// }

// // nolint
// //
// //lint:ignore U1000 // ignore this for now
// func getSortedKeys(m map[definitions.FieldInstance]any) []definitions.FieldInstance {
// 	keys := make([]definitions.FieldInstance, 0, len(m))

// 	for k := range m {
// 		keys = append(keys, k)
// 	}

// 	sort.SliceStable(keys, func(i, j int) bool {
// 		return keys[i].Ordinal < keys[j].Ordinal
// 	})

// 	return keys
// }

// func encodeVariableLength(len int) ([]byte, error) {
// 	if len <= 192 {
// 		return []byte{byte(len)}, nil
// 	}
// 	if len < 12480 {
// 		len -= 193
// 		b1 := byte((len >> 8) + 193)
// 		b2 := byte((len & 0xFF))
// 		return []byte{b1, b2}, nil
// 	}
// 	if len <= 918744 {
// 		len -= 12481
// 		b1 := byte((len >> 16) + 241)
// 		b2 := byte((len >> 8) & 0xFF)
// 		b3 := byte(len & 0xFF)
// 		return []byte{b1, b2, b3}, nil
// 	}
// 	return nil, ErrLengthPrefixTooLong
// }
