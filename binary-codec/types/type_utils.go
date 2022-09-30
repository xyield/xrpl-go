package types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	MinIOUExponent  = -96
	MaxIOUExponent  = 80
	MaxIOUPrecision = 16
	MinIOUMantissa  = 1000000000000000
	MaxIOUMantissa  = 9999999999999999

	NotXRPBitMask            = 0x80
	PosSignBitMask           = 0x4000000000000000
	ZeroCurrencyAmountHex    = 0x8000000000000000
	NativeAmountByteLength   = 8
	CurrencyAmountByteLength = 48

	MinXRP   = 1e-6
	MaxDrops = 1e17 // 100 billion XRP in drops aka 10^17

	AllowedIOUValueCharacters = "0123456789.-eE"                                                                   // going to do this with Regex
	AllowedIOUCodeCharacters  = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz?!@#$%^&*<>(){}[]|" // going to do this with Regex
)

type BigDecimal struct {
	Scale         int
	Precision     int
	UnscaledValue string
	Sign          int // 1 for negative, 0 for positive
}

type InvalidCharacterError struct {
	AllowedChars string
}

func (e *InvalidCharacterError) Error() string {
	return fmt.Sprintf("value contains invalid characters. Only the following are allowed: '%s'", e.AllowedChars)
}

type InvalidNativeCharacterError struct {
	InvalidCharacter string
}

func (e *InvalidNativeCharacterError) Error() string {
	return fmt.Sprintf("value contains invalid character '%s'", e.InvalidCharacter)
}

type InvalidAmountError struct {
	Amount string
}

func (e *InvalidAmountError) Error() string {
	return fmt.Sprintf("value '%s' is an invalid amount", e.Amount)
}

type OutOfRangeError struct {
	Type string
}

func (e *OutOfRangeError) Error() string {
	return fmt.Sprintf("%s is out of range", e.Type)
}

type InvalidCodeError struct {
	Disallowed string
}

func (e *InvalidCodeError) Error() string {
	return fmt.Sprintf("'%s' is/are disallowed or invalid", e.Disallowed)
}

// Creates a new custom BigDecimal object from a value string
func NewBigDecimal(value string) (*BigDecimal, error) {

	bigDecimal := new(BigDecimal)
	var err error

	if containsInvalidIOUValueCharacters(value) {
		return nil, &InvalidCharacterError{AllowedChars: AllowedIOUValueCharacters}
	}
	if strings.HasPrefix(value, "-") { // if the value is negative, set the sign to negative
		bigDecimal.Sign = 1
	}

	value = strings.ToLower(strings.TrimPrefix(value, "-")) // remove the sign from the value

	prefix, suffix, ePresent := strings.Cut(value, "e") // split the value into prefix and suffix at the 'e' character, if present

	if strings.Contains(prefix, "-") { // if the value still has a minus sign in the prefix, return an error
		return nil, errors.New("value contains multiple '-' characters, excluding the exponent sign")
	}
	if strings.Contains(prefix, "e") || strings.Contains(suffix, "e") { // if the prefix or suffix still contains an 'e' character, return an error
		return nil, errors.New("value contains multiple 'e' or 'E' characters")
	}

	emptyCheck := strings.Trim(prefix, "0") // remove all zeros from the prefix

	if emptyCheck == "" || emptyCheck == "." { // if the prefix is empty or just a decimal point, set everything to 0 or "" and return
		bigDecimal.Scale = 0
		bigDecimal.Precision = 0
		bigDecimal.UnscaledValue = ""
		bigDecimal.Sign = 0
		return bigDecimal, nil
	}

	containsDecimal, decimalErr := containsDecimal(value) // check if the value contains a SINGLE decimal point

	if decimalErr != nil {
		return nil, decimalErr
	}

	decimalPrefix, decimalSuffix, _ := strings.Cut(prefix, ".") // split the prefix into decimal prefix and decimal suffix at the '.' character, if present

	if ePresent { // if the value contains an 'e' character
		bigDecimal.Scale, err = strconv.Atoi(suffix) // convert the suffix to an int, which is the scale
		if err != nil {
			return nil, err
		}

		if containsDecimal { // if the value contains a SINGLE decimal point
			decimalSuffixNoTrailingZeros := strings.TrimRight(decimalSuffix, "0")         // remove trailing zeros from the decimal suffix
			decimalPrefixNoLeadingZeros := strings.TrimLeft(decimalPrefix, "0")           // remove leading zeros from the decimal prefix
			bigDecimal.Scale = bigDecimal.Scale - len(decimalSuffixNoTrailingZeros)       // subtract the length of the decimal suffix from the scale
			bigDecimal.UnscaledValue = strings.Trim((decimalPrefix + decimalSuffix), "0") // remove leading and trailing zeros from the concatenated decimal prefix and decimal suffix to get the unscaled value

			if decimalSuffixNoTrailingZeros == "" { // if the decimal suffix is empty because it only contained trailing zeros
				bigDecimal.Scale = bigDecimal.Scale + (len(decimalPrefixNoLeadingZeros) - len(bigDecimal.UnscaledValue)) // add the difference between the length of the decimal prefix and the length of the unscaled value, to the scale
			}

		} else if !containsDecimal { // if the value does not contain a SINGLE decimal point
			prefixNoTrailingZeros := strings.Trim(prefix, "0")                                             // remove trailing zeros from the prefix
			prefixNoLeadingZeros := strings.TrimLeft(prefix, "0")                                          // remove leading zeros from the prefix
			bigDecimal.Scale = bigDecimal.Scale + (len(prefixNoLeadingZeros) - len(prefixNoTrailingZeros)) // add the difference between the length of the prefix and the length of the prefix with trailing zeros removed, to the scale
			bigDecimal.UnscaledValue = prefixNoTrailingZeros                                               // set the unscaled value to the prefix with trailing zeros removed
		} else {
			return nil, err // Do I need this?
		}
	}

	if !ePresent { // if the value does not contain an 'e' character

		if containsDecimal && decimalErr == nil { // if the value contains a SINGLE decimal point
			decimalSuffixNoTrailingZeros := strings.TrimRight(decimalSuffix, "0")                        // remove trailing zeros from the decimal suffix
			decimalPrefixNoLeadingZeros := strings.TrimLeft(decimalPrefix, "0")                          // remove leading zeros from the decimal prefix
			bigDecimal.Scale = -len(decimalSuffixNoTrailingZeros)                                        // set the scale to the negative of the length of the decimal suffix with trailing zeros removed
			bigDecimal.UnscaledValue = strings.Trim((decimalPrefix + decimalSuffixNoTrailingZeros), "0") // remove leading and trailing zeros from the concatenated decimal prefix and decimal suffix with trailing zeros removed to get the unscaled value

			if decimalSuffixNoTrailingZeros == "" { // if the decimal suffix is empty because it only contained trailing zeros
				bigDecimal.Scale = len(decimalPrefixNoLeadingZeros) - len(bigDecimal.UnscaledValue) // set the scale to the difference between the length of the decimal prefix and the length of the unscaled value
			}

		} else if !containsDecimal && decimalErr == nil { // if the value does not contain a SINGLE decimal point
			decimalPrefixNoTrailingZeros := strings.TrimRight(prefix, "0")             // remove trailing zeros from the prefix
			bigDecimal.Scale = len(prefix) - len(decimalPrefixNoTrailingZeros)         // set the scale to the difference between the length of the prefix and the length of the prefix with trailing zeros removed
			bigDecimal.UnscaledValue = strings.Trim(decimalPrefixNoTrailingZeros, "0") // remove leading and trailing zeros from the prefix with trailing zeros removed to get the unscaled value
		} else {
			return nil, err // Do I need this?
		}
	}

	bigDecimal.Precision = len(bigDecimal.UnscaledValue) // set the precision to the length of the unscaled value

	return bigDecimal, nil
}
