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
func NewBigDecimal(value string) (bd *BigDecimal, err error) {

	// check if the value string contains only allowed characters
	if !bigDecimalRegEx(value) {
		return nil, ErrInvalidCharacter
	}

	v := strings.ToLower(value)
	bd = new(BigDecimal)

	// check if the value is negative and set the sign accordingly
	bd.Sign, v = checkAndSetSign(v)

	// check if the value contains the 'e' character and split the string into prefix and suffix accordingly
	p, s, eFound := strings.Cut(v, "e")
	trimP := strings.Trim(p, "0")

	// if the prefix without trailing & leading zeros is empty or only contains a decimal character, return an error
	if trimP == "" || trimP == "." {
		return nil, ErrInvalidZeroValue
	}

	// check if the value contains a decimal character and split the string into prefix and suffix accordingly
	decP, decS, decFound := strings.Cut(p, ".")

	bd.Scale, bd.UnscaledValue = getScaleAndUnscaledVal(eFound, decFound, p, s, decP, decS)

	if bd.UnscaledValue == "" {
		return nil, ErrInvalidZeroValue
	}

	bd.Precision = len(bd.UnscaledValue)
	return
}

func getScaleAndUnscaledVal(eFound, decFound bool, p, s, decP, decS string) (sc int, uv string) {

	// if the 'e' character is present, calculate the scale and unscaled value according to the rules for scientific notation
	// Otherwise, calculate the scale and unscaled value according to the rules for decimal notation

	if eFound {
		// convert the suffix to an integer, which is scale. Will error if the integer is too large
		// if error occurs, return empty unscaled value
		sc, err := strconv.Atoi(s)
		if err != nil {
			return 0, ""
		}
		if decFound {
			return valHasDecimal(sc, decP, decS)
		} else {
			return valNoDecimal(eFound, sc, p, decP)
		}
	} else {
		if decFound {
			return valHasDecimal(0, decP, decS)
		} else {
			return valNoDecimal(eFound, 0, p, decP)
		}
	}
}

func valHasDecimal(scale int, decP, decS string) (sc int, uv string) {
	uv = strings.Trim((decP + decS), "0")
	sc = scale - len(strings.TrimRight(decS, "0"))
	if strings.TrimRight(decS, "0") == "" {
		sc = scale + len(strings.TrimLeft(decP, "0")) - len(uv)
	}
	return
}

func valNoDecimal(eFound bool, scale int, prefix, decP string) (sc int, uv string) {
	if eFound {
		uv = strings.Trim(prefix, "0")
		sc = scale + len(strings.TrimLeft(prefix, "0")) - len(uv)
		return
	} else {
		uv = strings.Trim(decP, "0")
		sc = len(prefix) - len(strings.TrimRight(decP, "0"))
		return
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
