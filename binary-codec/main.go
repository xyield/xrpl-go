package binarycodec

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"sort"

	"github.com/xyield/xrpl-go/binary-codec/definitions"
	"github.com/xyield/xrpl-go/binary-codec/types"

	//nolint
	"golang.org/x/crypto/ripemd160" //lint:ignore SA1019 // ignore this for now
)

// Encode: encodes a transaction or other object from json to the canonical binary format as a hex string.
func Encode(json map[string]interface{}) (string, error) {

	fimap, err := createFieldInstanceMapFromJson(json)

	if err != nil {
		return "", err
	}

	sk := getSortedKeys(fimap)

	var sink []byte

	for _, v := range sk {

		h, err := EncodeFieldID(v.FieldName)

		if err != nil {
			return "", err
		}

		sink = append(sink, h...)
		// fmt.Println(hex.EncodeToString(sink))

		// need to write bytes to new buffers
		// amount, uint, hash all big endian
		st := types.GetSerializedType(v.Type)
		b, err := st.SerializeJson(fimap[v])
		if err != nil {
			return "", err
		}

		// fmt.Println(buf.Bytes())
		sink = append(sink, b...)
		// fmt.Println(hex.EncodeToString(sink))
	}

	// Loop through and create map of map[FieldInstance]interface{}
	// Sort by Ordinal
	// Start serializing
	//	optimize encode from field id codec, making same call twice

	// fmt.Println(string(sink))

	return hex.EncodeToString(sink), nil
}

// func Serialize(json string) (string, error) {
// 	return "120007220008000024001ABED82A2380BF2C2019001ABED764D55920AC9391400000000000000000000000000055534400000000000A20B3C85F482532A9578DBB3950B85CA06594D165400000037E11D60068400000000000000A732103EE83BB432547885C219634A1BC407A9DB0474145D69737D09CCDC63E1DEE7FE3744630440220143759437C04F7B61F012563AFE90D8DAFC46E86035E1D965A9CED282C97D4CE02204CFD241E86F17E011298FC1A39B63386C74306A5DE047E213B0F29EFA4571C2C8114DD76483FACDEE26E60D8A586BB58D09F27045C46", nil
// }

// nolint
//
//lint:ignore U1000 // ignore this for now
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

// nolint
//
//lint:ignore U1000 // ignore this for now
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

// func accountIDToBytes(address string) byte {}

// func amountToBytes(a string) byte {}

// func blobToBytes(hex string) byte {}

func calcHash(buf []byte, hasher hash.Hash) []byte {
	_, _ = hasher.Write(buf)
	return hasher.Sum(nil)
}

// func hashToBytes(hex string) byte {}

// func hash128(buf []byte) []byte {

// 	if len(buf) != 16 {
// 		panic("hash128 only supports 16 byte buffers")
// 	}

// 	return calcHash(buf, sha256.New())[:16]
// }

// func hash128ToBytes(hex string) byte {}

func Hash160(buf []byte) []byte {

	if len(buf) != 20 {
		panic("Hash160 only supports 20 byte buffers")
	}

	return calcHash(calcHash(buf, sha256.New()), ripemd160.New())
}

// func hash160ToBytes(hex string) (byte, error) {
// 	b := hashToBytes(hex)

// 	if len(b) != 20 {
// 		return 0, errors.New("Hash160 is not 160 bits long")
// 	}

// 	return b, nil
// }

// func Hash256(buf []byte) []byte {

// 	if len(buf) != 32 {
// 		panic("Hash256 only supports 32 byte buffers")
// 	}

// 	return calcHash(calcHash(buf, sha256.New()), sha256.New())
// }

// func hash256ToBytes(hex string) (byte, error) {
// 	b := hashToBytes(hex)

// 	if len(b) != 32 {
// 		return 0, errors.New("Hash256 is not 256 bits long")
// 	}

// 	return b, nil
// }

// func pathsetToBytes([][]string) (byte, error) {
// }

// func arrayToBytes(array []string) byte {

// }

// func objectToBytes(jsonObj string) byte {

// }

// func uint8ToBytes(i uint8) byte {
// 	return byte(i)
// }

// func uint16ToBytes(i uint16) byte {
// 	return byte(i)
// }

// func uint32ToBytes(i uint32) byte {
// 	return byte(i)
// }
