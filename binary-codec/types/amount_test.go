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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDecimal(tt.input); got != tt.want {
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
			name:   "valid xrp value",
			input:  "1.0",
			ExpErr: nil,
		},
		{
			name:   "invalid xrp value - out of range",
			input:  "0.000000007",
			ExpErr: errors.New("XRP value is an invalid XRP amount"),
		},
		{
			name:   "invalid xrp value - no decimal",
			input:  "1",
			ExpErr: errors.New("XRP value must contain a decimal"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := VerifyXrpValue(tt.input)

			if tt.ExpErr != nil {
				assert.Equal(t, tt.ExpErr, err)
			} else {
				assert.Nil(t, err)
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
			name:   "valid iou value",
			input:  "3.0",
			ExpErr: nil,
		},
		{
			name:   "invalid iou value - out of range precision",
			input:  "0.000000000000000000007",
			ExpErr: errors.New("IOU value is an invalid IOU amount - precision is too large > 16"),
		},
		{
			name:   "invalid iou value - out of range exponent too large",
			input:  "998000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			ExpErr: errors.New("IOU value is an invalid IOU amount - precision is too large > 16"),
		},
		{
			name:   "invalid iou value - out of range exponent too small",
			input:  "0.998456890560948998456890560948998456890560948998456890560948998456890560948998456890560948998456890560948998456890560948998456890560948998456890560948998456890560948998456890560948998456890560948",
			ExpErr: errors.New("IOU value is an invalid IOU amount - precision is too large > 16"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := VerifyIOUValue(tt.input)

			if tt.ExpErr != nil {
				assert.Equal(t, tt.ExpErr, err)
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
			input:    "3.0567",
			expected: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := SerializeIssuedCurrencyValue(tt.input)

			assert.Equal(t, tt.expected, got)

		})
	}
}
