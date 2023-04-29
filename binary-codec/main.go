package binarycodec

import (
	"encoding/hex"
	"errors"
	"strings"

	"github.com/xyield/xrpl-go/binary-codec/definitions"

	"github.com/xyield/xrpl-go/binary-codec/serdes"
	"github.com/xyield/xrpl-go/binary-codec/types"
)

var ErrSigningClaimFieldNotFound = errors.New("'Channel' & 'Amount' fields are both required, but were not found")

const (
	txMultiSigPrefix          = "534D5400"
	paymentChannelClaimPrefix = "434C4D00"
	txSigPrefix               = "53545800"
)

// Encode: encodes a transaction or other object from json to the canonical binary format as a hex string.
func Encode(json map[string]any) (string, error) {

	st := &types.STObject{}
	b, err := st.FromJson(json)
	if err != nil {
		return "", err
	}

	return strings.ToUpper(hex.EncodeToString(b)), nil
}

// EncodeForMultiSign: encodes a transaction into binary format in preparation for providing one
// signature towards a multi-signed transaction.
// (Only encodes fields that are intended to be signed.)
func EncodeForMultisigning(json map[string]any, xrpAccountID string) (string, error) {

	st := &types.AccountID{}

	// SigningPubKey is required for multi-signing but should be set to empty string.

	json["SigningPubKey"] = ""

	suffix, err := st.FromJson(xrpAccountID)
	if err != nil {
		return "", err
	}

	encoded, err := Encode(removeNonSigningFields(json))

	if err != nil {
		return "", err
	}

	return strings.ToUpper(txMultiSigPrefix + encoded + hex.EncodeToString(suffix)), nil
}

// Encodes a transaction into binary format in preparation for signing.
func EncodeForSigning(json map[string]any) (string, error) {

	encoded, err := Encode(removeNonSigningFields(json))

	if err != nil {
		return "", err
	}

	return strings.ToUpper(txSigPrefix + encoded), nil
}

// EncodeForPaymentChannelClaim: encodes a payment channel claim into binary format in preparation for signing.
func EncodeForSigningClaim(json map[string]any) (string, error) {

	if json["Channel"] == nil || json["Amount"] == nil {
		return "", ErrSigningClaimFieldNotFound
	}

	channel, err := types.NewHash256().FromJson(json["Channel"])

	if err != nil {
		return "", err
	}

	t := &types.UInt64{}
	amount, err := t.FromJson(json["Amount"])

	if err != nil {
		return "", err
	}

	return strings.ToUpper(paymentChannelClaimPrefix + hex.EncodeToString(channel) + hex.EncodeToString(amount)), nil
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

func Decode(hexEncoded string) (map[string]any, error) {
	b, err := hex.DecodeString(hexEncoded)
	if err != nil {
		return nil, err
	}
	p := serdes.NewBinaryParser(b)
	st := &types.STObject{}
	m, err := st.ToJson(p)
	if err != nil {
		return nil, err
	}

	return m.(map[string]any), nil
}
