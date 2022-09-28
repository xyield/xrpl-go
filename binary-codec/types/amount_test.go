package types

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsInvalidCharacters(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "contains invalid character",
			input: "1.0a",
			want:  true,
		},
		{
			name:  "does not contain invalid character",
			input: "1.0",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsInvalidCharacters(tt.input); got != tt.want {
				t.Errorf("containsInvalidCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
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
		expErr error
	}{
		{
			name:   "invalid xrp value",
			input:  "1.0",
			expErr: errors.New("XRP value must not contain a decimal"),
		},
		{
			name:   "invalid xrp value - out of range",
			input:  "0.000000007",
			expErr: errors.New("XRP value must not contain a decimal"),
		},
		{
			name:   "valid xrp value - no decimal",
			input:  "125000708",
			expErr: nil,
		},
		{
			name:   "valid xrp value - no decimal - negative value",
			input:  "-125000708",
			expErr: errors.New("XRP value is an invalid XRP amount"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expErr != nil {
				assert.Equal(t, tt.expErr, VerifyXrpValue(tt.input))
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
			expErr: errors.New("IOU value is an invalid IOU amount - precision is too large > 16"),
		},
		{
			name:   "invalid iou value - out of range exponent too large",
			input:  "998000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			expErr: errors.New("IOU value is an invalid IOU amount - exponent is out of range"),
		},
		{
			name:   "invalid iou value - out of range exponent too small",
			input:  "0.0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000998",
			expErr: errors.New("IOU value is an invalid IOU amount - exponent is out of range"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := VerifyIOUValue(tt.input)
			if tt.expErr != nil {
				assert.Error(t, tt.expErr, err)
			} else {
				assert.Nil(t, err)
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
			expErr:         errors.New("XRP value is an invalid XRP amount"),
		},
		{
			name:           "invalid xrp value - decimal",
			input:          "125000708.0",
			expectedOutput: nil,
			expErr:         errors.New("XRP value must not contain a decimal"),
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
			expErr:         errors.New("XRP value is an invalid XRP amount"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := SerializeXrpAmount(tt.input)
			if tt.expErr != nil {
				assert.Error(t, tt.expErr, err)
			} else {
				assert.Nil(t, err)
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
			name:        "valid zero value",
			input:       "0",
			expected:    []byte{0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			expectedErr: nil,
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

			got, err := serializeIssuedCurrencyValue(tt.input)

			if tt.expectedErr != nil {
				assert.EqualError(t, tt.expectedErr, err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, got)
			}

		})
	}
}

func TestCreateValueObject(t *testing.T) {
	tests := []struct {
		name          string
		inputValue    string
		inputCurrency string
		inputIssuer   string
		expected      *ValueObj
		expectedErr   error
	}{
		{
			name:          "successfully created value object",
			inputValue:    "7072.8",
			inputCurrency: "USD",
			inputIssuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
			expected: &ValueObj{
				Value:    []byte{0xD5, 0x59, 0x20, 0xAC, 0x93, 0x91, 0x40, 0x00},
				Currency: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x55, 0x53, 0x44, 0x00, 0x00, 0x00, 0x00, 0x00},
				Issuer:   []byte{0xa, 0x20, 0xb3, 0xc8, 0x5f, 0x48, 0x25, 0x32, 0xa9, 0x57, 0x8d, 0xbb, 0x39, 0x50, 0xb8, 0x5c, 0xa0, 0x65, 0x94, 0xd1},
			},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := createValueObject(tt.inputValue, tt.inputCurrency, tt.inputIssuer)

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

func TestNewBigDecimal(t *testing.T) {
	tt := []struct {
		name      string
		input     string
		expBigDec *BigDecimal
		expErr    error
	}{
		{
			name:      "pos - pos 'e' - no dec - no lead 0 - no trail 0",
			input:     "123e4",
			expBigDec: &BigDecimal{Scale: 4, Precision: 3, UnscaledValue: "123", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "pos - neg 'e' - no dec",
			input:     "3E-6",
			expBigDec: &BigDecimal{Scale: -6, Precision: 1, UnscaledValue: "3", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "neg - pos 'e' - no dec",
			input:     "-123e4",
			expBigDec: &BigDecimal{Scale: 4, Precision: 3, UnscaledValue: "123", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "neg - neg 'e' - no dec",
			input:     "-3E-6",
			expBigDec: &BigDecimal{Scale: -6, Precision: 1, UnscaledValue: "3", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "pos - pos 'e' - with dec",
			input:     "123.456e4",
			expBigDec: &BigDecimal{Scale: 1, Precision: 6, UnscaledValue: "123456", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "pos - neg 'e' - with dec",
			input:     "3.456E-6",
			expBigDec: &BigDecimal{Scale: -9, Precision: 4, UnscaledValue: "3456", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "neg - pos 'e' - with dec",
			input:     "-123.456e4",
			expBigDec: &BigDecimal{Scale: 1, Precision: 6, UnscaledValue: "123456", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "neg - neg 'e' - with dec",
			input:     "-3.456E-6",
			expBigDec: &BigDecimal{Scale: -9, Precision: 4, UnscaledValue: "3456", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "pos - pos 'e' - no dec - lead 0",
			input:     "000123e4",
			expBigDec: &BigDecimal{Scale: 4, Precision: 3, UnscaledValue: "123", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "pos - neg 'e' - no dec - lead 0",
			input:     "0003E-6",
			expBigDec: &BigDecimal{Scale: -6, Precision: 1, UnscaledValue: "3", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "neg - pos 'e' - no dec - lead 0",
			input:     "-000123e4",
			expBigDec: &BigDecimal{Scale: 4, Precision: 3, UnscaledValue: "123", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "neg - neg 'e' - no dec - lead 0",
			input:     "-0003E-6",
			expBigDec: &BigDecimal{Scale: -6, Precision: 1, UnscaledValue: "3", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "pos - pos 'e' - with dec - lead 0",
			input:     "000123.456e4",
			expBigDec: &BigDecimal{Scale: 1, Precision: 6, UnscaledValue: "123456", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "pos - neg 'e' - with dec - lead 0",
			input:     "0003.456E-6",
			expBigDec: &BigDecimal{Scale: -9, Precision: 4, UnscaledValue: "3456", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "neg - pos 'e' - with dec - lead 0",
			input:     "-000123.456e4",
			expBigDec: &BigDecimal{Scale: 1, Precision: 6, UnscaledValue: "123456", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "neg - neg 'e' - with dec - lead 0",
			input:     "-0003.456E-6",
			expBigDec: &BigDecimal{Scale: -9, Precision: 4, UnscaledValue: "3456", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "pos - pos 'e' - no dec - trail 0",
			input:     "123000e4",
			expBigDec: &BigDecimal{Scale: 7, Precision: 3, UnscaledValue: "123", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "pos - neg 'e' - no dec - trail 0",
			input:     "32300E-62",
			expBigDec: &BigDecimal{Scale: -60, Precision: 3, UnscaledValue: "323", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "neg - pos 'e' - no dec - trail 0",
			input:     "-123000e4",
			expBigDec: &BigDecimal{Scale: 7, Precision: 3, UnscaledValue: "123", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "neg - neg 'e' - no dec - trail 0",
			input:     "-32300E-62",
			expBigDec: &BigDecimal{Scale: -60, Precision: 3, UnscaledValue: "323", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "pos - pos 'e' - with dec - trail 0",
			input:     "123000.456e4",
			expBigDec: &BigDecimal{Scale: 1, Precision: 9, UnscaledValue: "123000456", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "pos - neg 'e' - with dec - trail 0",
			input:     "32300.456E-62",
			expBigDec: &BigDecimal{Scale: -65, Precision: 8, UnscaledValue: "32300456", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "neg - pos 'e' - with dec - trail 0",
			input:     "-123000.456e4",
			expBigDec: &BigDecimal{Scale: 1, Precision: 9, UnscaledValue: "123000456", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "neg - neg 'e' - with dec - trail 0",
			input:     "-32300.456E-62",
			expBigDec: &BigDecimal{Scale: -65, Precision: 8, UnscaledValue: "32300456", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "pos - pos 'e' - no dec - lead & trail 0",
			input:     "000123000e4",
			expBigDec: &BigDecimal{Scale: 7, Precision: 3, UnscaledValue: "123", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "pos - neg 'e' - no dec - lead & trail 0",
			input:     "00032300E-62",
			expBigDec: &BigDecimal{Scale: -60, Precision: 3, UnscaledValue: "323", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "neg - pos 'e' - no dec - lead & trail 0",
			input:     "-000123000e4",
			expBigDec: &BigDecimal{Scale: 7, Precision: 3, UnscaledValue: "123", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "neg - neg 'e' - no dec - lead & trail 0",
			input:     "-00032300E-62",
			expBigDec: &BigDecimal{Scale: -60, Precision: 3, UnscaledValue: "323", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "pos - pos 'e' - with dec - lead & trail 0",
			input:     "000123000.456e4",
			expBigDec: &BigDecimal{Scale: 1, Precision: 9, UnscaledValue: "123000456", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "pos - neg 'e' - with dec - lead & trail 0",
			input:     "00032300.456E-62",
			expBigDec: &BigDecimal{Scale: -65, Precision: 8, UnscaledValue: "32300456", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "neg - pos 'e' - with dec - lead & trail 0",
			input:     "-000123000.456e4",
			expBigDec: &BigDecimal{Scale: 1, Precision: 9, UnscaledValue: "123000456", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "neg - neg 'e' - with dec - lead & trail 0",
			input:     "-00032300.456E-62",
			expBigDec: &BigDecimal{Scale: -65, Precision: 8, UnscaledValue: "32300456", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "pos - pos 'e' - with dec - lead 0 & trail 0 - lead 0 after dec",
			input:     "000123000.04567e4",
			expBigDec: &BigDecimal{Scale: -1, Precision: 11, UnscaledValue: "12300004567", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "pos - neg 'e' - with dec - lead 0 & trail 0 - lead 0 after dec",
			input:     "00032300.04567E-62",
			expBigDec: &BigDecimal{Scale: -67, Precision: 10, UnscaledValue: "3230004567", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "neg - pos 'e' - with dec - lead 0 & trail 0 - lead 0 after dec",
			input:     "-000123000.04567e4",
			expBigDec: &BigDecimal{Scale: -1, Precision: 11, UnscaledValue: "12300004567", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "neg - neg 'e' - with dec - lead 0 & trail 0 - lead 0 after dec",
			input:     "-00032300.04567E-62",
			expBigDec: &BigDecimal{Scale: -67, Precision: 10, UnscaledValue: "3230004567", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "pos - pos 'e' - with dec - lead 0 & trail 0 - trail 0 after dec",
			input:     "000123000.45678000e4",
			expBigDec: &BigDecimal{Scale: -1, Precision: 11, UnscaledValue: "12300045678", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "pos - neg 'e' - with dec - lead 0 & trail 0 - trail 0 after dec",
			input:     "00032300.45678000E-62",
			expBigDec: &BigDecimal{Scale: -67, Precision: 10, UnscaledValue: "3230045678", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "neg - pos 'e' - with dec - lead 0 & trail 0 - trail 0 after dec",
			input:     "-000123000.45678000e4",
			expBigDec: &BigDecimal{Scale: -1, Precision: 11, UnscaledValue: "12300045678", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "neg - neg 'e' - with dec - lead 0 & trail 0 - trail 0 after dec",
			input:     "-00032300.45678000E-62",
			expBigDec: &BigDecimal{Scale: -67, Precision: 10, UnscaledValue: "3230045678", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "pos - pos 'e' - with dec - lead 0 & trail 0 - lead & trail 0 after dec",
			input:     "000123000.045670000e4",
			expBigDec: &BigDecimal{Scale: -1, Precision: 11, UnscaledValue: "12300004567", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "pos - neg 'e' - with dec - lead 0 & trail 0 - lead & trail 0 after dec",
			input:     "00032300.045670000E-62",
			expBigDec: &BigDecimal{Scale: -67, Precision: 10, UnscaledValue: "3230004567", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "neg - pos 'e' - with dec - lead 0 & trail 0 - lead & trail 0 after dec",
			input:     "-000123000.045670000e4",
			expBigDec: &BigDecimal{Scale: -1, Precision: 11, UnscaledValue: "12300004567", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "neg - neg 'e' - with dec - lead 0 & trail 0 - lead & trail 0 after dec",
			input:     "-00032300.045670000E-62",
			expBigDec: &BigDecimal{Scale: -67, Precision: 10, UnscaledValue: "3230004567", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "pos - pos 'e' - with dec - only 0 before dec - lead & trail 0 after dec",
			input:     "000000000.045670000e4",
			expBigDec: &BigDecimal{Scale: -1, Precision: 4, UnscaledValue: "4567", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "pos - neg 'e' - with dec - only 0 before dec - lead & trail 0 after dec",
			input:     "000000000.045670000E-62",
			expBigDec: &BigDecimal{Scale: -67, Precision: 4, UnscaledValue: "4567", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "neg - pos 'e' - with dec - only 0 before dec - lead & trail 0 after dec",
			input:     "-000000000.045670000e4",
			expBigDec: &BigDecimal{Scale: -1, Precision: 4, UnscaledValue: "4567", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "neg - neg 'e' - with dec - only 0 before dec - lead & trail 0 after dec",
			input:     "-000000000.045670000E-62",
			expBigDec: &BigDecimal{Scale: -67, Precision: 4, UnscaledValue: "4567", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "pos - pos 'e' - with dec - lead 0 & trail 0 - only 0 after dec",
			input:     "000123000.000000000e4",
			expBigDec: &BigDecimal{Scale: 7, Precision: 3, UnscaledValue: "123", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "pos - neg 'e' - with dec - lead 0 & trail 0 - only 0 after dec",
			input:     "00032300.000000000E-62",
			expBigDec: &BigDecimal{Scale: -60, Precision: 3, UnscaledValue: "323", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "neg - pos 'e' - with dec - lead 0 & trail 0 - only 0 after dec",
			input:     "-000123000.000000000e4",
			expBigDec: &BigDecimal{Scale: 7, Precision: 3, UnscaledValue: "123", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "neg - neg 'e' - with dec - lead 0 & trail 0 - only 0 after dec",
			input:     "-00032300.000000000E-62",
			expBigDec: &BigDecimal{Scale: -60, Precision: 3, UnscaledValue: "323", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "pos - NO 'e' - no dec",
			input:     "123456789",
			expBigDec: &BigDecimal{Scale: 0, Precision: 9, UnscaledValue: "123456789", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "neg - NO 'e' - no dec",
			input:     "-123456789",
			expBigDec: &BigDecimal{Scale: 0, Precision: 9, UnscaledValue: "123456789", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "pos - NO 'e' - with dec",
			input:     "12345.6789",
			expBigDec: &BigDecimal{Scale: -4, Precision: 9, UnscaledValue: "123456789", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "neg - NO 'e' - with dec",
			input:     "-12345.6789",
			expBigDec: &BigDecimal{Scale: -4, Precision: 9, UnscaledValue: "123456789", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "pos - NO 'e' - no dec - lead 0",
			input:     "0000123456",
			expBigDec: &BigDecimal{Scale: 0, Precision: 6, UnscaledValue: "123456", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "neg - NO 'e' - no dec - lead 0",
			input:     "-0000123456",
			expBigDec: &BigDecimal{Scale: 0, Precision: 6, UnscaledValue: "123456", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "pos - NO 'e' - with dec - lead 0",
			input:     "000012345.6789",
			expBigDec: &BigDecimal{Scale: -4, Precision: 9, UnscaledValue: "123456789", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "neg - NO 'e' - with dec - lead 0",
			input:     "-000012345.6789",
			expBigDec: &BigDecimal{Scale: -4, Precision: 9, UnscaledValue: "123456789", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "pos - NO 'e' - no dec - trail 0",
			input:     "123456000",
			expBigDec: &BigDecimal{Scale: 3, Precision: 6, UnscaledValue: "123456", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "neg - NO 'e' - no dec - trail 0",
			input:     "-123456000",
			expBigDec: &BigDecimal{Scale: 3, Precision: 6, UnscaledValue: "123456", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "pos - NO 'e' - with dec - trail 0",
			input:     "123456000.45678",
			expBigDec: &BigDecimal{Scale: -5, Precision: 14, UnscaledValue: "12345600045678", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "neg - NO 'e' - with dec - trail 0",
			input:     "-123456000.45678",
			expBigDec: &BigDecimal{Scale: -5, Precision: 14, UnscaledValue: "12345600045678", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "pos - NO 'e' - no dec - lead 0 & trail 0",
			input:     "0000123456000",
			expBigDec: &BigDecimal{Scale: 3, Precision: 6, UnscaledValue: "123456", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "neg - NO 'e' - no dec - lead 0 & trail 0",
			input:     "-0000123456000",
			expBigDec: &BigDecimal{Scale: 3, Precision: 6, UnscaledValue: "123456", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "pos - NO 'e' - with dec - lead 0 & trail 0",
			input:     "0000123456000.45678",
			expBigDec: &BigDecimal{Scale: -5, Precision: 14, UnscaledValue: "12345600045678", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "neg - NO 'e' - with dec - lead 0 & trail 0",
			input:     "-0000123456000.45678",
			expBigDec: &BigDecimal{Scale: -5, Precision: 14, UnscaledValue: "12345600045678", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "pos - NO 'e' - with dec - lead 0 & trail 0 - lead 0 after dec",
			input:     "0000123456000.045678",
			expBigDec: &BigDecimal{Scale: -6, Precision: 15, UnscaledValue: "123456000045678", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "neg - NO 'e' - with dec - lead 0 & trail 0 - lead 0 after dec",
			input:     "-0000123456000.045678",
			expBigDec: &BigDecimal{Scale: -6, Precision: 15, UnscaledValue: "123456000045678", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "pos - NO 'e' - with dec - lead 0 & trail 0 - trail 0 after dec",
			input:     "0000123456000.45678000",
			expBigDec: &BigDecimal{Scale: -5, Precision: 14, UnscaledValue: "12345600045678", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "neg - NO 'e' - with dec - lead 0 & trail 0 - trail 0 after dec",
			input:     "-0000123456000.45678000",
			expBigDec: &BigDecimal{Scale: -5, Precision: 14, UnscaledValue: "12345600045678", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "pos - NO 'e' - with dec - lead 0 & trail 0 - lead & trail 0 after dec",
			input:     "0000123456000.045678000",
			expBigDec: &BigDecimal{Scale: -6, Precision: 15, UnscaledValue: "123456000045678", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "neg - NO 'e' - with dec - lead 0 & trail 0 - lead & trail 0 after dec",
			input:     "-0000123456000.045678000",
			expBigDec: &BigDecimal{Scale: -6, Precision: 15, UnscaledValue: "123456000045678", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "pos - NO 'e' - with dec - only 0 before dec - lead & trail 0 after dec",
			input:     "0000000000000.045678000",
			expBigDec: &BigDecimal{Scale: -6, Precision: 5, UnscaledValue: "45678", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "neg - NO 'e' - with dec - only 0 before dec - lead & trail 0 after dec",
			input:     "-0000000000000.045678000",
			expBigDec: &BigDecimal{Scale: -6, Precision: 5, UnscaledValue: "45678", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "pos - NO 'e' - with dec - lead 0 & trail 0 - only 0 after dec",
			input:     "0000123456000.000000000",
			expBigDec: &BigDecimal{Scale: 3, Precision: 6, UnscaledValue: "123456", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "neg - NO 'e' - with dec - lead 0 & trail 0 - only 0 after dec",
			input:     "-0000123456000.000000000",
			expBigDec: &BigDecimal{Scale: 3, Precision: 6, UnscaledValue: "123456", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "just 0 - no dec",
			input:     "0000",
			expBigDec: &BigDecimal{Scale: 0, Precision: 0, UnscaledValue: "", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "just 0 - no dec - with minus sign",
			input:     "-0000",
			expBigDec: &BigDecimal{Scale: 0, Precision: 0, UnscaledValue: "", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "just 0 - with dec",
			input:     "000.0000",
			expBigDec: &BigDecimal{Scale: 0, Precision: 0, UnscaledValue: "", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "just 0 - with dec - with minus sign",
			input:     "-000.0000",
			expBigDec: &BigDecimal{Scale: 0, Precision: 0, UnscaledValue: "", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "just 0 - no dec - no minus sign - with 'e' ",
			input:     "0000E72",
			expBigDec: &BigDecimal{Scale: 0, Precision: 0, UnscaledValue: "", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "just 0 - no dec - with minus sign - with 'e' ",
			input:     "-0000e72",
			expBigDec: &BigDecimal{Scale: 0, Precision: 0, UnscaledValue: "", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "just 0 - with dec - no minus sign - with 'e' ",
			input:     "000.0000E72",
			expBigDec: &BigDecimal{Scale: 0, Precision: 0, UnscaledValue: "", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "just 0 - with dec - with minus sign - with 'e' ",
			input:     "-000.0000e72",
			expBigDec: &BigDecimal{Scale: 0, Precision: 0, UnscaledValue: "", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "boundary test, exactly 16 significant digits",
			input:     "1234567891234567",
			expBigDec: &BigDecimal{Scale: 0, Precision: 16, UnscaledValue: "1234567891234567", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "boundary test, exactly 16 significant digits with decimal",
			input:     "1234567891234567.0000000000000000",
			expBigDec: &BigDecimal{Scale: 0, Precision: 16, UnscaledValue: "1234567891234567", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "boundary test, exactly 16 significant digits with decimal in middle of string",
			input:     "12345678.91234567",
			expBigDec: &BigDecimal{Scale: -8, Precision: 16, UnscaledValue: "1234567891234567", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "boundary test, exactly 16 significant digits with decimal in middle of string and trailing zeros",
			input:     "123456789123456700.0000000000000000",
			expBigDec: &BigDecimal{Scale: 2, Precision: 16, UnscaledValue: "1234567891234567", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "boundary test, exactly 16 significant digits with decimal in middle of string and leading zeros - negative",
			input:     "-123456789123456700.0000000000000000",
			expBigDec: &BigDecimal{Scale: 2, Precision: 16, UnscaledValue: "1234567891234567", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "Minimum token value",
			input:     "-9999999999999999e80",
			expBigDec: &BigDecimal{Scale: 80, Precision: 16, UnscaledValue: "9999999999999999", Sign: "-"},
			expErr:    nil,
		},
		{
			name:      "Minimum non zero absolute value",
			input:     "1000000000000000e-96",
			expBigDec: &BigDecimal{Scale: -81, Precision: 1, UnscaledValue: "1", Sign: ""},
			expErr:    nil,
		},
		{
			name:      "contains invalid chars",
			input:     "12345678r90.1234567890a",
			expBigDec: nil,
			expErr:    errors.New("value contains invalid characters: only '0-9' '.' '-' 'e' and 'E' are allowed"),
		},
		{
			name:      "contains multiple decimal points",
			input:     "12345678.90.1234567890",
			expBigDec: nil,
			expErr:    errors.New("invalid - string contains more than one decimal point"),
		},
		{
			name:      "contains multiple 'e' or 'E'",
			input:     "12345678e90E1234567890",
			expBigDec: nil,
			expErr:    errors.New("value contains multiple 'e' or 'E' characters"),
		},
		{
			name:      "contains multiple '-' signs:  excluding the exponent sign",
			input:     "-1234-567890.1234567890e-9",
			expBigDec: nil,
			expErr:    errors.New("value contains multiple '-' characters, excluding the exponent sign"),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := NewBigDecimal(tc.input)

			if tc.expErr != nil {
				assert.EqualError(t, tc.expErr, err.Error())
			} else {
				assert.Equal(t, tc.expBigDec, got)
			}
		})
	}
}
