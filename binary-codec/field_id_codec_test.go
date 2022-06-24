package binarycodec

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyield/xrpl-go/binary-codec/definitions"
)

func TestEncode(t *testing.T) {
	tt := []struct {
		description string
		input       string
		expected    []byte
		expectedErr error
	}{
		{
			description: "Type Code and Field Code < 16",
			input:       "Sequence",
			expected:    []byte{36},
			expectedErr: nil,
		},
		{
			description: "Additional Type Code and Field Code < 16",
			input:       "DestinationTag",
			expected:    []byte{46},
			expectedErr: nil,
		},
		{
			description: "Type Code >= 16 and Field Code < 16",
			input:       "Paths",
			expected:    []byte{1, 18},
			expectedErr: nil,
		},
		{
			description: "Additional Type Code >= 16 and Field Code < 16",
			input:       "CloseResolution",
			expected:    []byte{1, 16},
			expectedErr: nil,
		},
		{
			description: "Type Code < 16 and Field Code >= 16",
			input:       "SetFlag",
			expected:    []byte{32, 33},
			expectedErr: nil,
		},
		{
			description: "Additional Type Code < 16 and Field Code >= 16",
			input:       "Nickname",
			expected:    []byte{80, 18},
			expectedErr: nil,
		},
		{
			description: "Type Code and Field Code >= 16",
			input:       "TickSize",
			expected:    []byte{0, 16, 16},
			expectedErr: nil,
		},
		{
			description: "Additional Type Code and Field Code >= 16",
			input:       "UNLModifyDisabling",
			expected:    []byte{0, 16, 17},
			expectedErr: nil,
		},
		{
			description: "Non existent field name",
			input:       "yurt",
			expected:    nil,
			expectedErr: &definitions.NotFoundError{Instance: "FieldName", Input: "yurt"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			got, err := EncodeFieldID(tc.input)

			if tc.expectedErr != nil {
				assert.EqualError(t, err, tc.expectedErr.Error())
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, got)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tt := []struct {
		description string
		input       []byte
		expected    string
		expectedErr error
	}{
		{
			description: "Decode Sequence fieldId (Type Code and Field Code < 16)",
			input:       []byte{36},
			expected:    "Sequence",
			expectedErr: nil,
		},
		{
			description: "Decode DestinationTag fieldId (Type Code and Field Code < 16)",
			input:       []byte{46},
			expected:    "DestinationTag",
			expectedErr: nil,
		},
		{
			description: "Decode Paths fieldId (Type Code >= 16 and Field Code < 16)",
			input:       []byte{1, 18},
			expected:    "Paths",
			expectedErr: nil,
		},
		{
			description: "Decode CloseResolution fieldId (Type Code >= 16 and Field Code < 16)",
			input:       []byte{1, 16},
			expected:    "CloseResolution",
			expectedErr: nil,
		},
		{
			description: "Decode SetFlag fieldId (Type Code < 16 and Field Code >= 16)",
			input:       []byte{32, 33},
			expected:    "SetFlag",
			expectedErr: nil,
		},
		{
			description: "Decode Nickname fieldId (Type Code < 16 and Field Code >= 16)",
			input:       []byte{80, 18},
			expected:    "Nickname",
			expectedErr: nil,
		},
		{
			description: "Decode TickSize fieldId (Type Code and Field Code >= 16)",
			input:       []byte{0, 16, 16},
			expected:    "TickSize",
			expectedErr: nil,
		},
		{
			description: "Decode UNLModifyDisabling fieldId (Type Code and Field Code >= 16)",
			input:       []byte{0, 16, 17},
			expected:    "UNLModifyDisabling",
			expectedErr: nil,
		},
		{
			description: "Non existent field name",
			input:       []byte{255},
			expected:    "",
			expectedErr: &definitions.NotFoundErrorFieldHeader{Instance: "FieldHeader"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			hex := hex.EncodeToString(tc.input)
			// fmt.Println("hex string:", hex)
			actual, err := DecodeFieldID(hex)
			// fmt.Println(actual)

			if tc.expectedErr != nil {
				assert.Error(t, err, tc.expectedErr.Error())
				assert.Zero(t, actual)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, actual)
			}
		})
	}
}
