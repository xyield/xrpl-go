package bigdecimal

import (
	"errors"
	"fmt"
	"math/big"
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
	ErrInvalidZeroValue = errors.New("value cannot be zero")
	ErrInvalidScale     = errors.New("scale too large")
)

type BigDecimal struct {
	Scale         int
	Precision     int
	UnscaledValue string
	Sign          int // 1 for negative, 0 for positive
}

func (bd *BigDecimal) GetScaledValue() string {
	unscaled, _ := new(big.Float).SetString(bd.UnscaledValue)

	scalingFactor := new(big.Float).SetFloat64(1)
	for i := 0; i < abs(bd.Scale); i++ {
		scalingFactor.Mul(scalingFactor, big.NewFloat(10))
	}

	var scaledValue *big.Float
	if bd.Scale >= 0 {
		scaledValue = new(big.Float).Mul(unscaled, scalingFactor)
	} else {
		scaledValue = new(big.Float).Quo(unscaled, scalingFactor)
	}

	if bd.Sign == 1 {
		scaledValue.Neg(scaledValue)
	}
	return strings.TrimSuffix(strings.TrimRight(scaledValue.Text('f', bd.Scale), "0"), ".")
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
	bd.Sign, v = handleSign(v)

	// check if the value contains the 'e' character and split the string into prefix and suffix accordingly
	p, s, eFound := strings.Cut(v, "e")

	// if the prefix without trailing & leading zeros is empty or only contains a decimal character, return an error
	trimP := strings.Trim(p, "0")
	if trimP == "" || trimP == "." {
		return nil, ErrInvalidZeroValue
	}

	// if the value contains the 'e' character, call the appropriate function to get the scale and unscaled value
	if eFound {
		bd.Scale, bd.UnscaledValue = getScaleAndUnscaledValWithE(p, s)
	} else {
		bd.Scale, bd.UnscaledValue = getScaleAndUnscaledValNoE(p, s)
	}

	if bd.UnscaledValue == "" {
		return nil, ErrInvalidZeroValue
	}

	bd.Precision = len(bd.UnscaledValue)
	return
}

func getScaleAndUnscaledValNoE(p, s string) (sc int, uv string) {

	// check if the value contains a decimal character and split the string into prefix and suffix accordingly
	decP, decS, decFound := strings.Cut(p, ".")
	if decFound {
		return valHasDecimal(0, decP, decS)
	} else {
		return valNoDecimalNoE(0, p, decP)
	}
}

func getScaleAndUnscaledValWithE(p, s string) (sc int, uv string) {
	// check if the value contains a decimal character and split the string into prefix and suffix accordingly
	decP, decS, decFound := strings.Cut(p, ".")
	sc, err := strconv.Atoi(s)
	if err != nil {
		return 0, ""
	}
	if decFound {
		return valHasDecimal(sc, decP, decS)
	} else {
		return valNoDecimalHasE(sc, p, decP)
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

func valNoDecimalNoE(scale int, prefix, decP string) (sc int, uv string) {
	uv = strings.Trim(decP, "0")
	sc = len(prefix) - len(strings.TrimRight(decP, "0"))
	return
}

func valNoDecimalHasE(scale int, prefix, decP string) (sc int, uv string) {
	uv = strings.Trim(prefix, "0")
	sc = scale + len(strings.TrimLeft(prefix, "0")) - len(uv)
	return

}

func handleSign(value string) (int, string) {
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
