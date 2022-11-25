package binarycodec

import (
	"encoding/hex"
	"strings"

	"github.com/xyield/xrpl-go/binary-codec/definitions"
	"github.com/xyield/xrpl-go/binary-codec/types"
)

const (
	txMultiSigPrefix = "534D5400"
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

// EncodeForMultiSign: encodes a transaction into binary format in preparation for providing one
// signature towards a multi-signed transaction.
// (Only encodes fields that are intended to be signed.)
func EncodeForMultisigning(json map[string]any, xrpAccountID map[string]any) (string, error) {

	st := &types.STObject{}

	// remove the SigningPubKey field because any existing signing keys
	// shouldn't be signed over again.

	delete(json, "SigningPubKey")

	suffix, err := st.SerializeJson(xrpAccountID)
	if err != nil {
		return "", err
	}

	encoded, err := Encode(removeNonSigningFields(json))

	if err != nil {
		return "", err
	}

	return strings.ToUpper(txMultiSigPrefix + encoded + hex.EncodeToString(suffix)), nil
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
