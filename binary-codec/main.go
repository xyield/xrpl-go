package binarycodec

import (
	"encoding/hex"
	"errors"
	"strings"

	"github.com/xyield/xrpl-go/binary-codec/definitions"
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

	channel, err := types.NewHash256().SerializeJson(json["Channel"])

	if err != nil {
		return "", err
	}

	st := &types.UInt64{}
	amount, err := st.SerializeJson(json["Amount"])

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
