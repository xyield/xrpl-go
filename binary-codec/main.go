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
