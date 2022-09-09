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
	MIN_IOU_MANTISSA  = 1e16
	MAX_IOU_MANTISSA  = 1e17 - int(1)

	NOT_XRP_BIT_MASK            = 0x80
	POS_SIGN_BIT_MASK           = 0x4000000000000000
	ZERO_CURRENCY_AMOUNT_HEX    = 0x8000000000000000
	NATIVE_AMOUNT_BYTE_LENGTH   = 8
	CURRENCY_AMOUNT_BYTE_LENGTH = 48

	MIN_XRP   = 1e-6
	MAX_DROPS = 1e17 // 100 billion XRP in drops aka 10^17
)

func (a *Amount) SerializeJson(value any) ([]byte, error) {

	return nil, nil
}

// XRP values shouldn't contain a decimal point BECAUSE they are represented as integers as drops
// returns true if the string contains a SINGLE decimal point character
func containsDecimal(s string) (bool, error) {
	decCount := strings.Count(s, ".") // count the number of decimal points
	if decCount > 1 {
		return true, errors.New("invalid - string contains more than one decimal point")
	}
	return strings.Contains(s, "."), nil
}

// validates the format of an XRP amount value
// XRP values shouldn't contain a decimal point BECAUSE they are represented as integers as drops
func VerifyXrpValue(value string) error {

	containsDecimal, err := containsDecimal(value)

	if err != nil {
		return err
	} else if containsDecimal {
		return errors.New("XRP value must not contain a decimal")
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
func VerifyIOUValue(value string) (prec int, e error) {

	precision, exp, digits, sign, err := getSignificantDigits(value)

	fmt.Println("precision:", precision)
	fmt.Println("exp:", exp)
	fmt.Println("digits:", digits)
	fmt.Println("sign:", sign)

	if err != nil {
		return 0, err
	}

	decimalValue, ok := new(big.Float).SetString(value) // bigFloat for precision

	if !ok {
		return 0, errors.New("failed to convert string to big.Float")
	}

	if decimalValue.Sign() == 0 {
		return precision, nil
	}

	// Exponent must be between -96 and 80 - Exponents are being generated correctly - checked with debugging
	// Precision is more complex to calculate because of varying definitions and understanding of what it is
	if exp < MIN_IOU_EXPONENT || exp > MAX_IOU_EXPONENT {
		return 0, errors.New("IOU value is an invalid IOU amount - exponent is out of range")
	}

	// TODO: need to verify there is no decimal point after being multiplied by exponent

	return precision, nil

}

// XRPL definition of precision is number of significant digits:
// Tokens can represent a wide variety of assets, including those typically measured in very small or very large denominations.
// This format uses significant digits and a power-of-ten exponent in a similar way to scientific notation.
// The format supports positive and negative significant digits and exponents within the specified range.
// Unlike typical floating-point representations of non-whole numbers, this format uses integer math for all calculations,
// so it always maintains 15 decimal digits of precision. Multiplication and division have adjustments to compensate for
// over-rounding in the least significant digits.

// This function (getSigDigits) needs to be rewritten to return the following {
// 1. The sign
// 2. The number of significant digits in the value
// 3. The actual significant digits in the value
// 4. The exponent
// }
func getSignificantDigits(value string) (prec int, exp int, dig int, sign string, err error) {

	var prefix, suffix, sigDigits string

	if strings.HasPrefix(value, "-") {
		sign = "-"
		value = strings.TrimPrefix(value, "-")
	}

	containsDecimal, decimalErr := containsDecimal(value)

	if decimalErr != nil {
		return 0, 0, 0, "", decimalErr
	} else if containsDecimal {

		sigDigits = strings.Trim(strings.ReplaceAll(value, ".", ""), "0") // remove all leading and trailing zeros and decimal point
		if strings.Contains(value, "E") {
			sigDigits = strings.Split(sigDigits, "E")[0] // remove E from the string
		}
		if strings.Contains(value, "e") {
			sigDigits = strings.Split(sigDigits, "e")[0] // remove e from the string
		}

		prefix = strings.TrimLeft(strings.Split(value, ".")[0], "0")  // get the leading digits before the decimal point and trim any leading zeros
		suffix = strings.TrimRight(strings.Split(value, ".")[1], "0") // get the trailing digits after the decimal point and trim any trailing zeros

		if len(sigDigits) > MAX_IOU_PRECISION {
			return 0, 0, 0, "", errors.New("IOU value is an invalid IOU amount - precision is too large > 16")
		}

		if len(prefix) == 0 || len(suffix)&len(prefix) > 0 {
			exp = -len(suffix)
		} else if len(suffix) == 0 {
			exp = len(prefix)
		}

		if strings.Contains(value, "E") {
			expString := strings.Split(value, "E")
			fmt.Println("expString:", expString[1])
			expInt, _ := strconv.ParseInt(expString[1], 10, 64)
			exp = int(expInt)
		}

		if strings.Contains(value, "e") {
			expString := strings.Split(value, "e")
			fmt.Println("expString:", expString[1])
			expInt, _ := strconv.ParseInt(expString[1], 10, 64)
			exp = int(expInt)
		}

		digits, _ := strconv.ParseInt(sigDigits, 10, 64)
		return len(sigDigits), exp, int(digits), sign, nil

	} else { // if the value does not contain a decimal point then it is an integer, hence no suffix
		sigDigits = strings.Trim(value, "0")  // remove all leading and trailing zeros
		prefix = strings.TrimLeft(value, "0") // get the leading digits before the decimal point and trim any leading zeros
		exp = len(prefix) - len(sigDigits)

		if strings.Contains(value, "E") {
			sigDigits = strings.Split(sigDigits, "E")[0] // remove E from the string
		}
		if strings.Contains(value, "e") {
			sigDigits = strings.Split(sigDigits, "e")[0] // remove e from the string
		}

		if strings.Contains(value, "E") { // if the value contains an E then it is in scientific notation
			expString := strings.Split(value, "E")
			fmt.Println("expString:", expString[1])
			expInt, _ := strconv.ParseInt(expString[1], 10, 64)
			exp = int(expInt)
		}

		if strings.Contains(value, "e") { // if the value contains an e then it is in scientific notation
			expString := strings.Split(value, "e")
			fmt.Println("expString:", expString[1])
			expInt, _ := strconv.ParseInt(expString[1], 10, 64)
			exp = int(expInt)
		}

		digits, _ := strconv.ParseInt(sigDigits, 10, 64)
		return len(sigDigits), exp, int(digits), sign, nil
	}
}

// Serializes the value field of an issued currency amount to its bytes representation
func serializeIssuedCurrencyValue(value string) ([]byte, error) {

	prec, err := VerifyIOUValue(value)

	fmt.Println("prec in serializeIssuedCurrencyValue:", prec)
	if err != nil {
		return nil, err
	}

	decimalValue, _ := new(big.Float).SetString(value) // bigFloat for precision
	if decimalValue.Sign() == 0 {
		x := new(big.Int).SetUint64(ZERO_CURRENCY_AMOUNT_HEX)
		return []byte(x.Bytes()), nil // if the value is zero, then return the zero currency amount hex
	}

	// convert components to integers

	// x == mantissa * 2^exponent
	mantissa := decimalValue.MinPrec()
	fmt.Println("mantissa:", mantissa)

	return nil, nil
}

// Returns true if this amount is a "native" XRP amount - first bit in first byte set to 0 for native XRP
func isNative(value []byte) bool {
	fmt.Printf("%08b", value)
	x := []byte(value)[0]&NOT_XRP_BIT_MASK == 0 // & bitwise operator returns 1 if both first bits are 1, otherwise 0
	return x
}

// Determines if this AmountType is positive - 2nd bit in 1st byte is set to 1 for positive amounts
func isPositive(value []byte) bool {
	fmt.Printf("%08b", value)
	x := []byte(value)[0]&0x40 > 0
	return x
}
