package types

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
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

// validates the format of an XRP amount value
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

// validates the format of an issued currency amount value
func VerifyIOUValue(value string) error {

	precision, err := getSignificantDigits(value)
	if err != nil {
		return err
	}

	decimalValue, ok := new(big.Float).SetString(value) // bigFloat for precision

	if !ok {
		return errors.New("failed to convert string to big.Float")
	}

	if decimalValue.Sign() == 0 {
		return nil
	}

	fmt.Println("decimalValue:", decimalValue)

	fmt.Println("precision calculated by getSignificantDigits:", precision)
	t := decimalValue.Text('f', -1)
	fmt.Println("t:", t)
	exponent := decimalValue.MantExp(decimalValue)

	fmt.Println("exponent calculated by MantExp:", exponent)

	// Exponent must be between -96 and 80 - Exponents are being generated correctly - checked with debugging
	// Precision is more complex to calculate because of varying definitions and understanding of what it is
	if exponent < MIN_IOU_EXPONENT || exponent > MAX_IOU_EXPONENT {
		return errors.New("IOU value is an invalid IOU amount - exponent is out of range")
	}

	// verify there is no decimal point after being multiplied by exponent
	if !containsDecimal(fmt.Sprint(exponent)) {
		return nil
	} else {
		return errors.New("IOU value is an invalid IOU amount - contains a decimal point")
	}

}

// XRPL definition of precision is number of significant digits:
// Tokens can represent a wide variety of assets, including those typically measured in very small or very large denominations.
// This format uses significant digits and a power-of-ten exponent in a similar way to scientific notation.
// The format supports positive and negative significant digits and exponents within the specified range.
// Unlike typical floating-point representations of non-whole numbers, this format uses integer math for all calculations,
// so it always maintains 15 decimal digits of precision. Multiplication and division have adjustments to compensate for
// over-rounding in the least significant digits.

func getSignificantDigits(value string) (int, error) {

	var prefix, suffix string

	if containsDecimal(value) {
		prefix = strings.TrimRight(strings.Split(value, ".")[0], "0") // get the leading digits before the decimal point and trim any leading zeros
		suffix = strings.TrimRight(strings.Split(value, ".")[1], "0") // get the trailing digits after the decimal point and trim any trailing zeros
	}

	prefixInt64, _ := strconv.ParseInt(prefix, 10, 64) // convert to int64 to allow zero comparison
	suffixInt64, _ := strconv.ParseInt(suffix, 10, 64) // convert to int64 to allow zero comparison

	if prefixInt64 == 0 { // if the prefix is zero, then the significant digits are the length of the suffix
		if len(fmt.Sprint(suffixInt64)) > MAX_IOU_PRECISION { // if the length of the suffix is greater than 16 digits, then it is an invalid amount
			return 0, errors.New("IOU value is an invalid IOU amount - precision is too large > 16")
		}
		suffix = strings.TrimLeft(suffix, "0") // trim any leading zeros from the suffix
		return len(suffix), nil                // return the length of the suffix
	} else if suffixInt64 == 0 { // if the suffix is zero, then the significant digits are the length of the prefix
		if len(fmt.Sprint(prefix)) > MAX_IOU_PRECISION { // if the length of the prefix is greater than 16 digits, then it is an invalid amount
			return 0, errors.New("IOU value is an invalid IOU amount - precision is too large > 16")
		}
		prefix = strings.TrimLeft(prefix, "0") // trim any leading zeros from the prefix
		return len(prefix), nil                // return the length of the prefix
	} else { // if both the prefix and suffix are not zero, then the significant digits are the length of the prefix + the length of the suffix
		if (len(prefix) + len(suffix)) > MAX_IOU_PRECISION { // if the length of the prefix + suffix is greater than 16 digits, then it is an invalid amount
			return 0, errors.New("IOU value is an invalid IOU amount - precision is too large > 16")
		}
		return len(prefix) + len(suffix), nil // return the length of the prefix + suffix
	}

	// decimalValue, _ := new(big.Float).SetString(value)

	// float, _ := strconv.ParseFloat(decimalValue.String(), 64) // 1)
	// Ffloat := strconv.FormatFloat(float, 'f', -1, 64)

	// fmt.Println("value:", value)
	// fmt.Println("float:", float)
	// fmt.Println("Ffloat:", Ffloat)

	// integerDigits := len(Ffloat) - 1
	// if !containsDecimal(Ffloat) {
	// 	integerDigits = len(Ffloat)
	// }
	// fmt.Println("Digits minus decimal point:", integerDigits)
	// fractionalDigits := len(Ffloat) - integerDigits
	// fmt.Println("digits after decimal point:", fractionalDigits)

	// precision := int(math.Ceil(float64(fractionalDigits)*math.Log2(10.0) + float64(integerDigits)*math.Log2(10.0)))
	// fmt.Println("calculated precision:", precision)
	// decimalValue, _ = decimalValue.SetPrec(uint(precision)).SetString(Ffloat)
	// fmt.Println("dv:", decimalValue)

	// // integerDigits - fractionalDigits gives the number of digits after the decimal point
	// // integerDigits gives the total number of digits
	// // if the number of digits after the decimal point is greater than 15, then the value should be invalid

	// if integerDigits > MAX_IOU_PRECISION || fractionalDigits >= MAX_IOU_PRECISION {
	// 	return 0, errors.New("IOU value is an invalid IOU amount - precision is too large > 16")
	// }

	// fmt.Println("dv precision:", decimalValue.Prec())
	// return integerDigits, nil
}

// Serializes the value field of an issued currency amount to its bytes representation
func serializeIssuedCurrencyValue(value string) []byte {

	VerifyIOUValue(value)

	decimalValue, _ := new(big.Float).SetString(value) // bigFloat for precision
	if decimalValue.Sign() == 0 {
		x := new(big.Int).SetUint64(ZERO_CURRENCY_AMOUNT_HEX)
		fmt.Println("x:", x)
		return []byte(x.Bytes())
	}

	// convert components to integers

	// x == mantissa * 2^exponent
	mantissa := decimalValue.MinPrec()
	fmt.Println("mantissa:", mantissa)

	return nil
}

// Returns true if this amount is a "native" XRP amount - first bit in first byte set to 0 for native XRP
func isNative(value []byte) bool {
	fmt.Printf("%08b", value)
	x := []byte(value)[0]&NOT_XRP_BIT_MASK == 0
	return x
}

// Determines if this AmountType is positive - 2nd bit in 1st byte is set to 1 for positive amounts
func isPositive(value []byte) bool {
	fmt.Printf("%08b", value)
	x := []byte(value)[0]&0x40 > 0
	return x
}
