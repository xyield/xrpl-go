package binarycodec

import (
	"bytes"
	"encoding/hex"
	"errors"
	"reflect"
	"strings"

	"github.com/CreatureDev/xrpl-go/model/transactions"

	"github.com/CreatureDev/xrpl-go/binary-codec/serdes"
	"github.com/CreatureDev/xrpl-go/binary-codec/types"
)

var ErrSigningClaimFieldNotFound = errors.New("'Channel' & 'Amount' fields are both required, but were not found")

const (
	txMultiSigPrefix          = "534D5400"
	paymentChannelClaimPrefix = "434C4D00"
	txSigPrefix               = "53545800"
)

func encode(tx transactions.Tx, onlySigning bool, mutations map[string]types.FieldMutation) (string, error) {
	st := &types.STObject{
		OnlySigning: onlySigning,
		Mutations:   mutations,
	}
	b, err := st.FromJson(tx)
	if err != nil {
		return "", err
	}

	return strings.ToUpper(hex.EncodeToString(b)), nil
}

// Encode converts a JSON transaction object to a hex string in the canonical binary format.
// The binary format is defined in XRPL's core codebase.
func Encode(tx transactions.Tx) (string, error) {
	return encode(tx, false, nil)
}

// EncodeForMultiSign: encodes a transaction into binary format in preparation for providing one
// signature towards a multi-signed transaction.
// (Only encodes fields that are intended to be signed.)
func EncodeForMultisigning(tx transactions.Tx, xrpAccountID string) (string, error) {

	st := &types.AccountID{}

	suffix, err := st.FromJson(xrpAccountID)
	if err != nil {
		return "", err
	}

	// SigningPubKey is required for multi-signing but should be set to empty string.
	err = setFieldFromTx(tx, "SigningPubKey", "placeholder", func(v any) bool {
		return v.(string) == ""
	})
	if err != nil {
		return "", err
	}
	encoded, err := encode(tx, true, map[string]types.FieldMutation{
		"SigningPubKey": types.Zero(),
	})

	if err != nil {
		return "", err
	}

	return strings.ToUpper(txMultiSigPrefix + encoded + hex.EncodeToString(suffix)), nil
}

// Encodes a transaction into binary format in preparation for signing.
func EncodeForSigning(tx transactions.Tx) (string, error) {

	encoded, err := encode(tx, true, nil)

	if err != nil {
		return "", err
	}

	return strings.ToUpper(txSigPrefix + encoded), nil
}

// EncodeForPaymentChannelClaim: encodes a payment channel claim into binary format in preparation for signing.
func EncodeForSigningClaim(tx transactions.PaymentChannelClaim) (string, error) {

	if tx.Channel == "" || tx.Amount == 0 {
		return "", ErrSigningClaimFieldNotFound
	}

	channel, err := types.NewHash256().FromJson(tx.Channel)

	if err != nil {
		return "", err
	}

	t := &types.Amount{}
	amount, err := t.FromJson(tx.Amount)

	if err != nil {
		return "", err

	}

	if bytes.HasPrefix(amount, []byte{0x40}) {
		amount = bytes.Replace(amount, []byte{0x40}, []byte{0x00}, 1)
	}

	return strings.ToUpper(paymentChannelClaimPrefix + hex.EncodeToString(channel) + hex.EncodeToString(amount)), nil
}

// Decode decodes a hex string in the canonical binary format into a JSON transaction object.
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

// Overwrites a field in a transaction with a new value if condition is met.
func setFieldFromTx(tx transactions.Tx, fieldName string, value any, condition func(any) bool) error {
	rv := reflect.ValueOf(tx)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	} else {
		return errors.New("invalid transaction")
	}
	if !rv.FieldByName(fieldName).IsValid() {
		return errors.New("invalid field name")
	}
	if condition != nil && condition(rv.FieldByName(fieldName).Interface()) {
		rv.FieldByName(fieldName).Set(reflect.ValueOf(value))
		return nil
	}
	return nil
}
