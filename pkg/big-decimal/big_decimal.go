package bigdecimal

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	AllowedCharacters = "0123456789.-eE"
	BigDecRegEx       = "-?(?:[0|1-9]\\d*)(?:\\.\\d+)?(?:[eE][+\\-]?\\d+)?"
)

var (
	ErrInvalidCharacter = fmt.Errorf("value contains invalid characters. Only the following are allowed: %q", AllowedCharacters)
	ErrInvalidZeroValue = fmt.Errorf("value cannot be zero")
	ErrInvalidScale     = fmt.Errorf("scale too large")
)

type BigDecimal struct {
	Scale         int
	Precision     int
	UnscaledValue string
	Sign          int // 1 for negative, 0 for positive
}

// Creates a new custom BigDecimal object from a value string
func NewBigDecimal(value string) (*BigDecimal, error) {

	if !bigDecimalRegEx(value) { // check if the value string contains only allowed characters
		return nil, ErrInvalidCharacter
	}

	v := strings.ToLower(value)
	bd := new(BigDecimal)
	bd.Sign, v = checkAndSetSign(v)     // check if the value is negative and set the sign accordingly
	p, s, eFound := strings.Cut(v, "e") // check if the value contains the 'e' character and split the string into prefix and suffix accordingly
	trimP := strings.Trim(p, "0")

	if trimP == "" || trimP == "." { // if the prefix without trailing & leading zeros is empty or only contains a decimal character, return an error
		return nil, ErrInvalidZeroValue
	}

	decP, decS, decFound := strings.Cut(p, ".")                                                       // if the decimal character is present, split the prefix into two parts, otherwise set the decimalPrefix to the prefix
	bd.Scale, bd.UnscaledValue = getScaleAndUnscaledVal(eFound, decFound, bd.Scale, p, s, decP, decS) // calculate the scale and unscaled value

	if bd.UnscaledValue == "" {
		return nil, ErrInvalidZeroValue
	}
	bd.Precision = len(bd.UnscaledValue) // set the precision to the length of the unscaled value

	return bd, nil
}

func getScaleAndUnscaledVal(eFound, decFound bool, scale int, prefix, suffix, decimalPrefix, decimalSuffix string) (int, string) {

	if eFound { // if the 'e' character is present, calculate the scale and unscaled value according to the rules for scientific notation
		scale, err := strconv.Atoi(suffix) // convert the suffix to an integer. Will return an error if the integer is too large
		if err != nil {
			return 0, "" // if the suffix cannot be converted to an integer, return empty unscaled value
		}
		if decFound {
			return valHasDecimal(eFound, scale, decimalPrefix, decimalSuffix)
		} else {
			return valNoDecimal(eFound, scale, prefix, decimalPrefix)
		}
	} else { // if the 'e' character is not present, calculate the scale and unscaled value according to the rules for decimal notation
		if decFound {
			return valHasDecimal(eFound, scale, decimalPrefix, decimalSuffix)
		} else {
			return valNoDecimal(eFound, scale, prefix, decimalPrefix)
		}
	}
}

func valHasDecimal(eFound bool, scale int, decimalPrefix, decimalSuffix string) (int, string) {
	unscaledValue := strings.Trim((decimalPrefix + decimalSuffix), "0")
	scale = scale - len(strings.TrimRight(decimalSuffix, "0"))

	lenUnscaledValue := len(unscaledValue)
	lenDecPrefixTrimL := len(strings.TrimLeft(decimalPrefix, "0"))

	if eFound {
		if strings.TrimRight(decimalSuffix, "0") == "" {
			scale = scale + lenDecPrefixTrimL - lenUnscaledValue
		}
		return scale, unscaledValue
	} else {
		if strings.TrimRight(decimalSuffix, "0") == "" {
			scale = lenDecPrefixTrimL - lenUnscaledValue
		}
		return scale, unscaledValue
	}
}

func valNoDecimal(eFound bool, scale int, prefix, decimalPrefix string) (int, string) {
	if eFound {
		unscaledValue := strings.Trim(prefix, "0")
		scale = scale + len(strings.TrimLeft(prefix, "0")) - len(unscaledValue)
		return scale, unscaledValue
	} else {
		unscaledValue := strings.Trim(decimalPrefix, "0")
		scale = len(prefix) - len(strings.TrimRight(decimalPrefix, "0"))
		return scale, unscaledValue
	}
}

func checkAndSetSign(value string) (int, string) {
	if strings.HasPrefix(value, "-") {
		return 1, strings.TrimPrefix(value, "-")
	}
	return 0, value
}

func bigDecimalRegEx(value string) bool {
	r := regexp.MustCompile(BigDecRegEx)
	m := r.FindAllString(value, -1)
	return len(m) == 1
}
