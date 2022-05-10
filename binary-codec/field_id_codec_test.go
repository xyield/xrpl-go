package binarycodec

import (
	"encoding/hex"
	"fmt"
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
			description: "Type and field code < 16",
			input:       "Sequence",
			expected:    []byte{36},
			expectedErr: nil,
		},
		{
			description: "Additional type and field code < 16",
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
			description: "Additional Type Code >=16 and Field code < 16",
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
			description: "Additional Type Code < 16 and Field Code > = 16",
			input:       "Nickname",
			expected:    []byte{80, 18},
			expectedErr: nil,
		},
		{
			description: "Type Code >= 16 and Field Code >=16",
			input:       "TickSize",
			expected:    []byte{0, 16, 16},
			expectedErr: nil,
		},
		{
			description: "Additional Type code and field code >= 16",
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
			got, err := Encode(tc.input)

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
		// {
		// 	description: "Decode Sequence fieldId",
		// 	input:       []byte{36},
		// 	expected:    "Sequence",
		// 	expectedErr: nil,
		// },
		{
			description: "Decode Paths fieldId",
			input:       []byte{1, 18},
			expected:    "Paths",
			expectedErr: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			hex := hex.EncodeToString(tc.input)
			fmt.Println("hex string:", hex)
			actual, err := Decode(hex)

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
