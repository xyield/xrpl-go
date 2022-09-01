package types

import (
	"errors"
	"math/big"
	"strings"
)

type Amount struct{}

const (
	MIN_IOU_EXPONENT  = -96
	MAX_IOU_EXPONENT  = 80
	MAX_IOU_PRECISION = 16
	MIN_IOU_MANTISSA  = 10e15
	MAX_IOU_MANTISSA  = 10e16 - int(1)

	NOT_XRP_BIT_MASK            = 0x80
	POS_SIGN_BIT_MASK           = 0x4000000000000000
	ZERO_CURRENCY_AMOUNT_HEX    = 0x8000000000000000
	NATIVE_AMOUNT_BYTE_LENGTH   = 8
	CURRENCY_AMOUNT_BYTE_LENGTH = 48

	MIN_XRP   = 1e-6
	MAX_DROPS = 1e17
)

func (a *Amount) SerializeJson(value any) ([]byte, error) {

	return nil, nil
}

// returns true if the string contains a decimal point character
func containsDecimal(s string) bool {
	return strings.ContainsAny(s, ".")
}

// validates the format of an XRP amount
func VerifyXrpValue(value string) error {

	if !containsDecimal(value) {
		return errors.New("XRP value must contain a decimal")
	}

	decimal := new(big.Float)
	decimal, ok := decimal.SetString(value) // may need to be bigFloat for precision

	if !ok {
		return errors.New("failed to convert string to big.Float")
	}

	if decimal.Sign() == 0 {
		return nil
	}

	if decimal.Cmp(big.NewFloat(MIN_XRP)) == -1 || decimal.Cmp(big.NewFloat(MAX_DROPS)) == 1 {
		return errors.New("XRP value is an invalid XRP amount")
	}

	return nil
}

// validates the format of an issued currency amount
func VerifyIOUValue(value string) error {

	decimalValue := new(big.Float)

	decimalValue, ok := decimalValue.SetString(value)

	if !ok {
		return errors.New("failed to convert string to big.Float")
	}

	if decimalValue.Sign() == 0 {
		return nil
	}

	// var exponent = decimalValue.MantExp(nil)
	return nil
}
