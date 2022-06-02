package addresscodec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeClassicAddressFromPublicKeyHex(t *testing.T) {
	tt := []struct {
		description    string
		input          string
		inputPrefix    []byte
		expectedOutput string
		expectedErr    error
	}{
		{
			description:    "Successfully generate classic address from public key hex string",
			input:          "ED9434799226374926EDA3B54B1B461B4ABF7237962EAE18528FEA67595397FA32",
			inputPrefix:    []byte{AccountAddressPrefix},
			expectedOutput: "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
			expectedErr:    nil,
		},
		{
			description:    "Successfully generate classic address from randomly generated hex string",
			input:          "42e04f55c9f92d0d2ece7a028f88fd37b740ea2086f136d9ed1ac842e5a0226125",
			inputPrefix:    []byte{AccountAddressPrefix},
			expectedOutput: "rJKhsipKHooQbtS3v5Jro6N5Q7TMNPkoAs",
			expectedErr:    nil,
		},
		{
			description:    "Invalid Public Key",
			input:          "yurt",
			inputPrefix:    []byte{AccountAddressPrefix},
			expectedOutput: "",
			expectedErr:    &EncodeLengthError{Instance: "PublicKey", Input: 0, Expected: 33},
		},
		{
			description:    "Valid Public Key, invalid Type Prefix",
			input:          "ED9434799226374926EDA3B54B1B461B4ABF7237962EAE18528FEA67595397FA32",
			inputPrefix:    []byte{0x00, 0x00},
			expectedOutput: "",
			expectedErr:    &EncodeLengthError{Instance: "TypePrefix", Input: 2, Expected: 1},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {

			got, err := EncodeClassicAddressFromPublicKeyHex(tc.input, tc.inputPrefix)

			if tc.expectedErr != nil {
				assert.EqualError(t, tc.expectedErr, err.Error())
			} else {
				if assert.NoError(t, tc.expectedErr, err) {
					assert.Equal(t, tc.expectedOutput, got)
				}
			}
		})
	}
}

func TestEncodeNodePublicKey(t *testing.T) {

}

func TestDecodeAddressToAccountID(t *testing.T) {
	tt := []struct {
		description       string
		input             string
		expectedPrefix    []byte
		expectedAccountID []byte
		expectedErr       error
	}{
		{
			description:       "Successful decode - 1",
			input:             "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
			expectedPrefix:    []byte{AccountAddressPrefix},
			expectedAccountID: []byte{0x88, 0xa5, 0xa5, 0x7c, 0x82, 0x9f, 0x40, 0xf2, 0x5e, 0xa8, 0x33, 0x85, 0xbb, 0xde, 0x6c, 0x3d, 0x8b, 0x4c, 0xa0, 0x82},
			expectedErr:       nil,
		},
		{
			description:       "Successful decode - 2",
			input:             "rJKhsipKHooQbtS3v5Jro6N5Q7TMNPkoAs",
			expectedPrefix:    []byte{AccountAddressPrefix},
			expectedAccountID: []byte{0xbd, 0xe4, 0x2b, 0xbd, 0x77, 0x5b, 0x46, 0x7e, 0x34, 0xfe, 0x48, 0x52, 0xe7, 0xce, 0x3d, 0xd2, 0x61, 0x3, 0xf7, 0x6c},
			expectedErr:       nil,
		},
		{
			description:       "Unsuccessful decode - 1",
			input:             "yurt",
			expectedPrefix:    nil,
			expectedAccountID: nil,
			expectedErr:       &InvalidClassicAddressError{Input: "yurt"},
		},
		{
			description:       "Unsuccessful decode - 2",
			input:             "davidschwartz",
			expectedPrefix:    nil,
			expectedAccountID: nil,
			expectedErr:       &InvalidClassicAddressError{Input: "davidschwartz"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {

			typePrefix, accountID, err := DecodeClassicAddressToAccountID(tc.input)

			if tc.expectedErr != nil {
				assert.EqualError(t, tc.expectedErr, err.Error())
				assert.Nil(t, tc.expectedPrefix, typePrefix)
				assert.Nil(t, tc.expectedAccountID, accountID)
			} else {
				assert.Equal(t, tc.expectedPrefix, typePrefix)
				assert.Equal(t, tc.expectedAccountID, accountID)
			}
		})
	}
}

func TestIsValidClassicAddress(t *testing.T) {
	tt := []struct {
		description string
		input       string
		expected    bool
	}{
		{
			description: "Valid classic address",
			input:       "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
			expected:    true,
		},
		{
			description: "Invalid classic address",
			input:       "yurt",
			expected:    false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			if tc.expected != true {
				assert.False(t, tc.expected, IsValidClassicAddress(tc.input))
			} else {
				assert.True(t, tc.expected, IsValidClassicAddress(tc.input))
			}
		})
	}
}
