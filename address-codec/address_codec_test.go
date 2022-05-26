package addresscodec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeAddressFromPublicKeyHex(t *testing.T) {
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

			got, err := EncodeAddressFromPublicKeyHex(tc.input, tc.prefix)

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
