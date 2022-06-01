package addresscodec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeClassicAddressFromPublicKeyHex(t *testing.T) {
	tt := []struct {
		description   string
		input         string
		prefix        []byte
		output        string
		expectedError error
	}{
		{
			description:   "Successfully generate address from public key hex string",
			input:         "ED9434799226374926EDA3B54B1B461B4ABF7237962EAE18528FEA67595397FA32",
			prefix:        []byte{AccountAddressPrefix},
			output:        "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
			expectedError: nil,
		},
		{
			description:   "Successfully generate address from randomly generated hex string",
			input:         "42e04f55c9f92d0d2ece7a028f88fd37b740ea2086f136d9ed1ac842e5a0226125",
			prefix:        []byte{AccountAddressPrefix},
			output:        "rJKhsipKHooQbtS3v5Jro6N5Q7TMNPkoAs",
			expectedError: nil,
		},
		{
			description:   "Invalid Public Key",
			input:         "yurt",
			prefix:        []byte{AccountAddressPrefix},
			output:        "",
			expectedError: &EncodeLengthError{Instance: "PublicKey", Input: AccountPublicKeyLength},
		},
		// More test cases
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {

			got, err := EncodeClassicAddressFromPublicKeyHex(tc.input, tc.prefix)

			if tc.expectedError != nil {
				assert.Error(t, tc.expectedError, err.Error())
			} else {
				if assert.NoError(t, tc.expectedError, err) {
					assert.Equal(t, tc.output, got)
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
		expectedError     error
	}{
		{
			description:       "Successful decode - 1",
			input:             "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
			expectedPrefix:    []byte{AccountAddressPrefix},
			expectedAccountID: []byte{0x88, 0xa5, 0xa5, 0x7c, 0x82, 0x9f, 0x40, 0xf2, 0x5e, 0xa8, 0x33, 0x85, 0xbb, 0xde, 0x6c, 0x3d, 0x8b, 0x4c, 0xa0, 0x82},
			expectedError:     nil,
		},
		{
			description:       "Successful decode - 2",
			input:             "rJKhsipKHooQbtS3v5Jro6N5Q7TMNPkoAs",
			expectedPrefix:    []byte{AccountAddressPrefix},
			expectedAccountID: []byte{0xbd, 0xe4, 0x2b, 0xbd, 0x77, 0x5b, 0x46, 0x7e, 0x34, 0xfe, 0x48, 0x52, 0xe7, 0xce, 0x3d, 0xd2, 0x61, 0x3, 0xf7, 0x6c},
			expectedError:     nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {

			typePrefix, accountID, err := DecodeClassicAddressToAccountID(tc.input)

			if tc.expectedError != nil {
				assert.Error(t, tc.expectedError, err.Error())
				assert.Nil(t, tc.expectedPrefix, typePrefix)
				assert.Nil(t, tc.expectedAccountID, accountID)
			} else {
				assert.Equal(t, tc.expectedPrefix, typePrefix)
				assert.Equal(t, tc.expectedAccountID, accountID)
			}
		})
	}
}
