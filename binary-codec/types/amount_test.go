package types

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDecimal(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "contains decimal",
			input: "1.0",
			want:  true,
		},
		{
			name:  "does not contain decimal",
			input: "1",
			want:  false,
		},
		{
			name:  "contains decimal - double dot",
			input: "1..0",
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := containsDecimal(tt.input); got != tt.want {
				t.Errorf("containsDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerifyXrpValue(t *testing.T) {

	tests := []struct {
		name   string
		input  string
		ExpErr error
	}{
		{
			name:   "invalid xrp value",
			input:  "1.0",
			ExpErr: errors.New("XRP value must not contain a decimal"),
		},
		{
			name:   "invalid xrp value - out of range",
			input:  "0.000000007",
			ExpErr: errors.New("XRP value must not contain a decimal"),
		},
		{
			name:   "valid xrp value - no decimal",
			input:  "125000708",
			ExpErr: nil,
		},
		{
			name:   "valid xrp value - no decimal - negative value",
			input:  "-125000708",
			ExpErr: errors.New("XRP value is an invalid XRP amount"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.ExpErr != nil {
				assert.Equal(t, tt.ExpErr, VerifyXrpValue(tt.input))
			} else {
				assert.Nil(t, VerifyXrpValue(tt.input))
			}
		})
	}
}

func TestVerifyIOUValue(t *testing.T) {

	tests := []struct {
		name   string
		input  string
		ExpErr error
	}{
		{
			name:   "valid iou value with decimal",
			input:  "3.6",
			ExpErr: nil,
		},
		{
			name:   "valid iou value - leading zero after decimal",
			input:  "345.023857",
			ExpErr: nil,
		},
		{
			name:   "valid iou value - negative value & multiple leading zeros before decimal",
			input:  "-000.2345",
			ExpErr: nil,
		},
		{
			name:   "invalid iou value - out of range precision",
			input:  "0.000000000000000000007265675687436598345739475",
			ExpErr: errors.New("IOU value is an invalid IOU amount - precision is too large > 16"),
		},
		{
			name:   "invalid iou value - out of range exponent too large",
			input:  "998000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			ExpErr: errors.New("IOU value is an invalid IOU amount - exponent is out of range"),
		},
		{
			name:   "invalid iou value - out of range exponent too small",
			input:  "0.0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000998",
			ExpErr: errors.New("IOU value is an invalid IOU amount - exponent is out of range"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := VerifyIOUValue(tt.input)
			if tt.ExpErr != nil {
				assert.EqualError(t, tt.ExpErr, err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestSerializeIssuedCurrencyValue(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []byte
	}{
		{
			name:     "valid zero value",
			input:    "0",
			expected: []byte{0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
		{
			name:     "valid value",
			input:    "334767.0567",
			expected: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, _ := serializeIssuedCurrencyValue(tt.input)

			assert.Equal(t, tt.expected, got)

		})
	}
}

func TestIsNative(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected bool
	}{
		{
			name:     "native XRP",
			input:    []byte{0, 64, 128, 32}, // 0 in binary is 00000000. If the first bit of the first byte is 0, it is deemed to be native XRP
			expected: true,
		},
		{
			name:     "not native XRP",
			input:    []byte{128, 0, 0, 1, 0, 1, 0, 0}, // 128 in binary is 10000000. If the first bit of the first byte is not 0, it is deemed to be not native XRP
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, isNative(tt.input))
		})
	}
}

func TestIsPositive(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected bool
	}{
		{
			name:     "positive",
			input:    []byte{64, 0, 0, 0}, // 64 in binary is 01000000. If the second bit of the first byte is 1, it is deemed positive
			expected: true,
		},
		{
			name:     "negative",
			input:    []byte{128, 0, 0, 0, 0, 0, 0, 0}, // 128 in binary is 10000000. If the second bit of the first byte is 0, it is deemed negative
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, isPositive(tt.input))
		})
	}
}

func TestCalculatePrecision(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		expPrec int
		expExp  int
		expDig  int
		expSign string
		expErr  error
	}{
		{
			name:    "valid value with E",
			input:   "-3.6E-6",
			expPrec: 2,
			expExp:  -6,
			expDig:  36,
			expSign: "-",
			expErr:  nil,
		},
		{
			name:    "valid value with e",
			input:   "-365e-6",
			expPrec: 3,
			expExp:  -6,
			expDig:  365,
			expSign: "-",
			expErr:  nil,
		},
		{
			name:    "correct precision",
			input:   "0.00007890000",
			expPrec: 3,
			expExp:  -7,
			expDig:  789,
			expSign: "",
			expErr:  nil,
		},
		{
			name:    "correct precision 2 - trailing zeros",
			input:   "00004326458.0001000",
			expPrec: 11,
			expExp:  -4,
			expDig:  43264580001,
			expSign: "",
			expErr:  nil,
		},
		{
			name:    "correct precision 3 - big number",
			input:   "0.0099845689056094899845689056094899845689056094899845689056094899845689056094899845689056094899845689056094899845689056094899845689056094899845689056094899845689056094899845689056094899845689056094800000000000",
			expPrec: 0,
			expExp:  0,
			expDig:  0,
			expSign: "",
			expErr:  errors.New("IOU value is an invalid IOU amount - precision is too large > 16"),
		},
		{
			name:    "correct precision 4 - leading zeros",
			input:   "0000.005466000",
			expPrec: 4,
			expExp:  -6,
			expDig:  5466,
			expSign: "",
			expErr:  nil,
		},
		{
			name:    "correct precision 5 - no decimal",
			input:   "5632890000",
			expPrec: 6,
			expExp:  4,
			expDig:  563289,
			expSign: "",
			expErr:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, exp, digits, sign, err := getSignificantDigits(tt.input)
			if tt.expErr != nil {
				assert.EqualError(t, tt.expErr, err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expPrec, got)
				assert.Equal(t, tt.expExp, exp)
				assert.Equal(t, tt.expDig, digits)
				assert.Equal(t, tt.expSign, sign)
			}
		})
	}
}
