package binarycodec

import (
	"encoding/hex"
	"strings"

	"github.com/xyield/xrpl-go/binary-codec/definitions"
	"github.com/xyield/xrpl-go/binary-codec/types"
)

const (
	txSigPrefix = "53545800"
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

// Encodes a transaction into binary format in preparation for signing.
func EncodeForSigning(json map[string]any) (string, error) {

	encoded, err := Encode(removeNonSigningFields(json))

	if err != nil {
		return "", err
	}

	return txSigPrefix + encoded, nil
}

func removeNonSigningFields(json map[string]any) map[string]any {

	for k := range json {
		fi, _ := definitions.Get().GetFieldInstanceByFieldName(k)

		if fi != nil && !fi.IsSigningField {
			delete(json, k)
		}
	}

	return json
}
