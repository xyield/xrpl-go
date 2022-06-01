package addresscodec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeClassicAddressFromPublicKeyHex(t *testing.T) {
	tt := []struct {
		description         string
		input               string
		inputPrefix         []byte
		expectedOutput      string
		expectedErrorType   error
		expectedErrorString string
	}{
		{
			description:         "Successfully generate classic address from public key hex string",
			input:               "ED9434799226374926EDA3B54B1B461B4ABF7237962EAE18528FEA67595397FA32",
			inputPrefix:         []byte{AccountAddressPrefix},
			expectedOutput:      "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
			expectedErrorType:   nil,
			expectedErrorString: "",
		},
		{
			description:         "Successfully generate classic address from randomly generated hex string",
			input:               "42e04f55c9f92d0d2ece7a028f88fd37b740ea2086f136d9ed1ac842e5a0226125",
			inputPrefix:         []byte{AccountAddressPrefix},
			expectedOutput:      "rJKhsipKHooQbtS3v5Jro6N5Q7TMNPkoAs",
			expectedErrorType:   nil,
			expectedErrorString: "",
		},
		{
			description:         "Invalid Public Key",
			input:               "yurt",
			inputPrefix:         []byte{AccountAddressPrefix},
			expectedOutput:      "",
			expectedErrorType:   &EncodeLengthError{},
			expectedErrorString: "`PublicKey` length should be 33 not 0",
		},
		{
			description:         "Valid Public Key, invalid Type Prefix",
			input:               "ED9434799226374926EDA3B54B1B461B4ABF7237962EAE18528FEA67595397FA32",
			inputPrefix:         []byte{0x00, 0x00},
			expectedOutput:      "",
			expectedErrorType:   &EncodeLengthError{},
			expectedErrorString: "`TypePrefix` length should be 1 not 2",
		},
		// More test cases
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {

			got, err := EncodeClassicAddressFromPublicKeyHex(tc.input, tc.inputPrefix)

			if tc.expectedErrorType != nil {
				assert.Error(t, tc.expectedErrorType, err.Error())
				assert.Equal(t, tc.expectedErrorString, err.Error())
			} else {
				if assert.NoError(t, tc.expectedErrorType, err) {
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
		description         string
		input               string
		expectedPrefix      []byte
		expectedAccountID   []byte
		expectedErrorType   error
		expectedErrorString string
	}{
		{
			description:         "Successful decode - 1",
			input:               "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
			expectedPrefix:      []byte{AccountAddressPrefix},
			expectedAccountID:   []byte{0x88, 0xa5, 0xa5, 0x7c, 0x82, 0x9f, 0x40, 0xf2, 0x5e, 0xa8, 0x33, 0x85, 0xbb, 0xde, 0x6c, 0x3d, 0x8b, 0x4c, 0xa0, 0x82},
			expectedErrorType:   nil,
			expectedErrorString: "",
		},
		{
			description:         "Successful decode - 2",
			input:               "rJKhsipKHooQbtS3v5Jro6N5Q7TMNPkoAs",
			expectedPrefix:      []byte{AccountAddressPrefix},
			expectedAccountID:   []byte{0xbd, 0xe4, 0x2b, 0xbd, 0x77, 0x5b, 0x46, 0x7e, 0x34, 0xfe, 0x48, 0x52, 0xe7, 0xce, 0x3d, 0xd2, 0x61, 0x3, 0xf7, 0x6c},
			expectedErrorType:   nil,
			expectedErrorString: "",
		},
		{
			description:         "Unsuccessful decode - 1",
			input:               "yurt",
			expectedPrefix:      nil,
			expectedAccountID:   nil,
			expectedErrorType:   &InvalidClassicAddressError{},
			expectedErrorString: "`yurt` is an invalid classic address",
		},
		{
			description:         "Unsuccessful decode - 2",
			input:               "davidschwartz",
			expectedPrefix:      nil,
			expectedAccountID:   nil,
			expectedErrorType:   &InvalidClassicAddressError{},
			expectedErrorString: "`davidschwartz` is an invalid classic address",
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {

			typePrefix, accountID, err := DecodeClassicAddressToAccountID(tc.input)

			if tc.expectedErrorType != nil {
				assert.Error(t, tc.expectedErrorType, err.Error())
				assert.Nil(t, tc.expectedPrefix, typePrefix)
				assert.Nil(t, tc.expectedAccountID, accountID)
				assert.Equal(t, tc.expectedErrorString, err.Error())
			} else {
				assert.Equal(t, tc.expectedPrefix, typePrefix)
				assert.Equal(t, tc.expectedAccountID, accountID)
			}
		})
	}
}

func TestIsValidClassicAddress(t *testing.T) {
	tt := []struct {
		description    string
		input          string
		expectedOutput bool
	}{
		{
			description:    "Valid classic address",
			input:          "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
			expectedOutput: true,
		},
		{
			description:    "Invalid classic address",
			input:          "yurt",
			expectedOutput: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			if tc.expectedOutput != true {
				assert.False(t, tc.expectedOutput, IsValidClassicAddress(tc.input))
			} else {
				assert.True(t, tc.expectedOutput, IsValidClassicAddress(tc.input))
			}
		})
	}
}
