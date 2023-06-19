package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	bigdecimal "github.com/xyield/xrpl-go/pkg/big-decimal"
)

func TestVerifyXrpValue(t *testing.T) {

	tests := []struct {
		name   string
		input  string
		expErr error
	}{
		{
			name:   "invalid xrp value",
			input:  "1.0",
			expErr: ErrInvalidXRPValue,
		},
		{
			name:   "invalid xrp value - out of range",
			input:  "0.000000007",
			expErr: ErrInvalidXRPValue,
		},
		{
			name:   "valid xrp value - no decimal",
			input:  "125000708",
			expErr: nil,
		},
		{
			name:   "valid xrp value - no decimal - negative value",
			input:  "-125000708",
			expErr: &InvalidAmountError{Amount: "-125000708"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expErr != nil {
				assert.Equal(t, tt.expErr, verifyXrpValue(tt.input))
			} else {
				assert.NoError(t, verifyXrpValue(tt.input))
			}
		})
	}
}

func TestVerifyIOUValue(t *testing.T) {

	tests := []struct {
		name   string
		input  string
		expErr error
	}{
		{
			name:   "valid iou value with decimal",
			input:  "3.6",
			expErr: nil,
		},
		{
			name:   "valid iou value - leading zero after decimal",
			input:  "345.023857",
			expErr: nil,
		},
		{
			name:   "valid iou value - negative value & multiple leading zeros before decimal",
			input:  "-000.2345",
			expErr: nil,
		},
		{
			name:   "invalid iou value - out of range precision",
			input:  "0.000000000000000000007265675687436598345739475",
			expErr: &OutOfRangeError{Type: "Precision"},
		},
		{
			name:   "invalid iou value - out of range exponent too large",
			input:  "998000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			expErr: &OutOfRangeError{Type: "Exponent"},
		},
		{
			name:   "invalid iou value - out of range exponent too small",
			input:  "0.0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000998",
			expErr: &OutOfRangeError{Type: "Exponent"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := verifyIOUValue(tt.input)
			if tt.expErr != nil {
				assert.EqualError(t, tt.expErr, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSerializeXrpAmount(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput []byte
		expErr         error
	}{
		{
			name:           "valid xrp value - 1",
			input:          "524801",
			expectedOutput: []byte{0x40, 0x00, 0x00, 0x00, 0x00, 0x8, 0x2, 0x01},
			expErr:         nil,
		},
		{
			name:           "valid xrp value - 2",
			input:          "7696581656832",
			expectedOutput: []byte{0x40, 0x00, 0x7, 0x00, 0x00, 0x4, 0x1, 0x00},
			expErr:         nil,
		},
		{
			name:           "valid xrp value - 3",
			input:          "10000000",
			expectedOutput: []byte{0x40, 0x00, 0x00, 0x00, 0x00, 0x98, 0x96, 0x80},
			expErr:         nil,
		},
		{
			name:           "invalid xrp value - negative",
			input:          "-125000708",
			expectedOutput: nil,
			expErr:         &InvalidAmountError{Amount: "-125000708"},
		},
		{
			name:           "invalid xrp value - decimal",
			input:          "125000708.0",
			expectedOutput: nil,
			expErr:         ErrInvalidXRPValue,
		},
		{
			name:           "boundary test - 1 less than max xrp value",
			input:          "99999999999999999",
			expectedOutput: []byte{0x41, 0x63, 0x45, 0x78, 0x5d, 0x89, 0xff, 0xff},
			expErr:         nil,
		},
		{
			name:           "boundary test - max xrp value",
			input:          "10000000000000000",
			expectedOutput: []byte{0x40, 0x23, 0x86, 0xf2, 0x6f, 0xc1, 0x00, 0x00},
			expErr:         nil,
		},
		{
			name:           "boundary test - 1 greater than max xrp value",
			input:          "100000000000000001",
			expectedOutput: nil,
			expErr:         &InvalidAmountError{Amount: "100000000000000001"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := SerializeXrpAmount(tt.input)
			if tt.expErr != nil {
				assert.EqualError(t, tt.expErr, err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedOutput, got)
			}
		})
	}
}

func TestSerializeIssuedCurrencyValue(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    []byte
		expectedErr error
	}{
		{
			name:        "invalid zero value",
			input:       "0",
			expected:    nil,
			expectedErr: bigdecimal.ErrInvalidZeroValue,
		},
		{
			name:        "valid value - 2",
			input:       "1",
			expected:    []byte{0xD4, 0x83, 0x8D, 0x7E, 0xA4, 0xC6, 0x80, 0x00},
			expectedErr: nil,
		},
		{
			name:        "valid value - 3",
			input:       "2.1",
			expected:    []byte{0xD4, 0x87, 0x75, 0xF0, 0x5A, 0x07, 0x40, 0x00},
			expectedErr: nil,
		},
		{
			name:        "valid value - from Transaction 1 in main_test.go",
			input:       "7072.8",
			expected:    []byte{0xD5, 0x59, 0x20, 0xAC, 0x93, 0x91, 0x40, 0x00},
			expectedErr: nil,
		},
		{
			name:        "valid value - from Transaction 3 in main_test.go",
			input:       "0.6275558355",
			expected:    []byte{0xd4, 0x56, 0x4b, 0x96, 0x4a, 0x84, 0x5a, 0xc0},
			expectedErr: nil,
		},
		{
			name:        "valid value - negative",
			input:       "-2",
			expected:    []byte{0x94, 0x87, 0x1A, 0xFD, 0x49, 0x8D, 0x00, 0x00},
			expectedErr: nil,
		},
		{
			name:        "valid value - negative - 2",
			input:       "-7072.8",
			expected:    []byte{0x95, 0x59, 0x20, 0xAC, 0x93, 0x91, 0x40, 0x00},
			expectedErr: nil,
		},
		{
			name:        "valid value - large currency amount",
			input:       "1111111111111111.0",
			expected:    []byte{0xD8, 0x43, 0xF2, 0x8C, 0xB7, 0x15, 0x71, 0xC7},
			expectedErr: nil,
		},
		{
			name:        "boundary test - max precision - max exponent",
			input:       "9999999999999999e80",
			expected:    []byte{0xec, 0x63, 0x86, 0xf2, 0x6f, 0xc0, 0xff, 0xff},
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := SerializeIssuedCurrencyValue(tt.input)

			if tt.expectedErr != nil {
				assert.EqualError(t, tt.expectedErr, err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, got)
			}

		})
	}
}

func TestSerializeIssuedCurrencyCode(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    []byte
		expectedErr error
	}{
		{
			name:        "valid standard currency - ISO4217 - USD",
			input:       "USD",
			expected:    []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x55, 0x53, 0x44, 0x00, 0x00, 0x00, 0x00, 0x00},
			expectedErr: nil,
		},
		{
			name:        "valid standard currency - ISO4217 - USD - hex",
			input:       "0x0000000000000000000000005553440000000000",
			expected:    []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x55, 0x53, 0x44, 0x00, 0x00, 0x00, 0x00, 0x00},
			expectedErr: nil,
		},
		{
			name:        "valid standard currency - non ISO4217 - BTC",
			input:       "BTC",
			expected:    []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x42, 0x54, 0x43, 0x00, 0x00, 0x00, 0x00, 0x00},
			expectedErr: nil,
		},
		{
			name:        "valid standard currency - non ISO4217 - BTC - hex",
			input:       "0x0000000000000000000000004254430000000000",
			expected:    []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x42, 0x54, 0x43, 0x00, 0x00, 0x00, 0x00, 0x00},
			expectedErr: nil,
		},
		{
			name:        "disallowed standard currency - XRP",
			input:       "XRP",
			expected:    nil,
			expectedErr: &InvalidCodeError{"XRP uppercase"},
		},
		{
			name:        "disallowed standard currency - XRP - hex",
			input:       "0000000000000000000000005852500000000000",
			expected:    nil,
			expectedErr: &InvalidCodeError{"XRP uppercase"},
		},
		{
			name:        "invalid standard currency - 4 characters",
			input:       "ABCD",
			expected:    nil,
			expectedErr: &InvalidCodeError{"ABCD"},
		},
		{
			name:        "valid non-standard currency - 4 characters - hex",
			input:       "0x4142434400000000000000000000000000000000",
			expected:    []byte{0x41, 0x42, 0x43, 0x44, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			expectedErr: nil,
		},
		{
			name:        "special case - XRP - hex",
			input:       "0x0000000000000000000000000000000000000000",
			expected:    []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			expectedErr: nil,
		},
		{
			name:        "standard currency - valid symbols in currency code - 3 characters",
			input:       "A*B",
			expected:    []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x41, 0x2a, 0x42, 0x00, 0x00, 0x00, 0x00, 0x00},
			expectedErr: nil,
		},
		{
			name:        "standard currency - valid symbols in currency code - 3 characters - hex",
			input:       "0x000000000000000000000000412a420000000000",
			expected:    []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x41, 0x2a, 0x42, 0x00, 0x00, 0x00, 0x00, 0x00},
			expectedErr: nil,
		},
		{
			name:        "standard currency - invalid characters in currency code",
			input:       "AD/",
			expected:    nil,
			expectedErr: ErrInvalidCurrencyCode,
		},
		{
			name:        "standard currency - invalid characters in currency code - hex",
			input:       "0x00000000000000000000000041442f0000000000",
			expected:    nil,
			expectedErr: ErrInvalidCurrencyCode,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := serializeIssuedCurrencyCode(tt.input)

			if tt.expectedErr != nil {
				assert.EqualError(t, tt.expectedErr, err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, got)
			}

		})
	}
}

func TestSerializeIssuedCurrencyAmount(t *testing.T) {
	tests := []struct {
		name          string
		inputValue    string
		inputCurrency string
		inputIssuer   string
		expected      []byte
		expectedErr   error
	}{
		{
			name:          "valid serialized issued currency amount",
			inputValue:    "7072.8",
			inputCurrency: "USD",
			inputIssuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
			expected:      []byte{0xD5, 0x59, 0x20, 0xAC, 0x93, 0x91, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x55, 0x53, 0x44, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0A, 0x20, 0xB3, 0xC8, 0x5F, 0x48, 0x25, 0x32, 0xA9, 0x57, 0x8D, 0xBB, 0x39, 0x50, 0xB8, 0x5C, 0xA0, 0x65, 0x94, 0xD1},
			expectedErr:   nil,
		},
		{
			name:          "valid serialized issued currency amount - 2",
			inputValue:    "0.6275558355",
			inputCurrency: "USD",
			inputIssuer:   "rweYz56rfmQ98cAdRaeTxQS9wVMGnrdsFp",
			expected:      []byte{0xd4, 0x56, 0x4b, 0x96, 0x4a, 0x84, 0x5a, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x55, 0x53, 0x44, 0x00, 0x00, 0x00, 0x00, 0x00, 0x69, 0xd3, 0x3b, 0x18, 0xd5, 0x33, 0x85, 0xf8, 0xa3, 0x18, 0x55, 0x16, 0xc2, 0xed, 0xa5, 0xde, 0xdb, 0x8a, 0xc5, 0xc6},
			expectedErr:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := SerializeIssuedCurrencyAmount(tt.inputValue, tt.inputCurrency, tt.inputIssuer)

			if tt.expectedErr != nil {
				assert.EqualError(t, tt.expectedErr, err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, got)
			}

		})
	}
}

func TestIsNative(t *testing.T) {
	tests := []struct {
		name     string
		input    byte
		expected bool
	}{
		{
			name:     "native XRP",
			input:    64, // 64 in binary is 01000000. If the first bit of the first byte is 0, it is deemed to be native XRP
			expected: true,
		},
		{
			name:     "not native XRP",
			input:    128, // 128 in binary is 10000000. If the first bit of the first byte is not 0, it is deemed to be not native XRP
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
		input    byte
		expected bool
	}{
		{
			name:     "positive",
			input:    64, // 64 in binary is 01000000. If the second bit of the first byte is 1, it is deemed positive
			expected: true,
		},
		{
			name:     "negative",
			input:    128, // 128 in binary is 10000000. If the second bit of the first byte is 0, it is deemed negative
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, isPositive(tt.input))
		})
	}
}
