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
		name  string
		input string
		err   error
	}{
		{
			name:  "valid xrp value",
			input: "1.0",
			err:   nil,
		},
		{
			name:  "invalid xrp value - out of range",
			input: "0.000000007",
			err:   errors.New("XRP value is an invalid XRP amount"),
		},
		{
			name:  "invalid xrp value - no decimal",
			input: "1",
			err:   errors.New("XRP value must contain a decimal"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := VerifyXrpValue(tt.input)

			if err != nil {
				assert.Equal(t, tt.err.Error(), err.Error())
			} else {
				assert.Nil(t, tt.err)
			}

		})
	}
}

func TestVerifyIOUValue(t *testing.T) {

	tests := []struct {
		name  string
		input string
		err   error
	}{
		{
			name:  "valid iou value",
			input: "1.0",
			err:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := VerifyIOUValue(tt.input)

			if err != nil {
				assert.Equal(t, tt.err.Error(), err.Error())
			} else {
				assert.Nil(t, tt.err)
			}

		})
	}
}
