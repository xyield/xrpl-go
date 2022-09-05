package types

import (
	"errors"
	"fmt"
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
	decimal, ok := decimal.SetString(value) // bigFloat for precision

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

	decimalValue, ok := new(big.Float).SetString(value) // bigFloat for precision

	if !ok {
		return errors.New("failed to convert string to big.Float")
	}

	if decimalValue.Sign() == 0 {
		return nil
	}

	fmt.Println("decimalValue raw:", decimalValue)
	precision := decimalValue.MinPrec()
	fmt.Println("precision raw", precision)
	exponent := decimalValue.MantExp(decimalValue)

	fmt.Println("decimalValue after:", decimalValue)
	fmt.Println("precision after:", precision)

	// Exponent must be between -96 and 80 - need further validation
	if precision > MAX_IOU_PRECISION || exponent < MIN_IOU_EXPONENT || exponent > MAX_IOU_EXPONENT {
		return errors.New("IOU value is an invalid IOU amount - precision is too large > 16")
	}

	// verify there is no decimal point after being multiplied by exponent
	if !containsDecimal(fmt.Sprint(exponent)) {
		return nil
	} else {
		return errors.New("IOU value is an invalid IOU amount - contains a decimal point")
	}

}

// Serializes the value field of an issued currency amount to its bytes representation
func SerializeIssuedCurrencyValue(value string) []byte {

	VerifyIOUValue(value)

	decimalValue, _ := new(big.Float).SetString(value) // bigFloat for precision
	if decimalValue.Sign() == 0 {
		x := new(big.Int).SetUint64(ZERO_CURRENCY_AMOUNT_HEX)
		fmt.Println("x:", x)
		return []byte(x.Bytes())
	}

	// convert components to integers

	return nil
}

// Returns true is this amount is a "native" XRP amount
func IsNative(value []byte) bool {
	return []byte(value)[0]&0x80 == 0
}

// Determines if this AmountType is positive
func IsPositive(value []byte) bool {
	// 2nd bit in 1st byte is set to 1 for positive amounts
	return []byte(value)[0]&0x40 > 0
}
