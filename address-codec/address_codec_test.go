package addresscodec

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	tt := []struct {
		description    string
		input          []byte
		inputPrefix    []byte
		inputLength    int
		expectedOutput string
		expectedErr    error
	}{
		{
			description:    "Successful encode - 1",
			input:          []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			inputPrefix:    []byte{AccountAddressPrefix},
			inputLength:    16,
			expectedOutput: "rrrrrrrrrrrrrrrrrp9U13b",
			expectedErr:    nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {

			assert.Equal(t, tc.expectedOutput, Encode(tc.input, tc.inputPrefix, tc.inputLength))

		})
	}
}

func TestDecode(t *testing.T) {
	tt := []struct {
		description    string
		input          string
		inputPrefix    []byte
		expectedOutput []byte
		expectedErr    error
	}{
		{
			description:    "successful decode - 1",
			input:          "rrrrrrrrrrrrrrrrr",
			inputPrefix:    []byte{AccountAddressPrefix},
			expectedOutput: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			expectedErr:    nil,
		},
		{
			description:    "successful decode - 2",
			input:          "rrrrrrrrrrrrrrrrrp9U13b",
			inputPrefix:    []byte{AccountAddressPrefix},
			expectedOutput: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2c, 0xa7, 0xf0, 0x98},
			expectedErr:    nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expectedOutput, Decode(tc.input, tc.inputPrefix))
		})
	}
}

func TestEncodeClassicAddressFromPublicKeyHex(t *testing.T) {
	tt := []struct {
		description    string
		input          string
		inputPrefix    []byte
		expectedOutput string
		expectedErr    error
	}{
		{
			description:    "Successfully generate address from a 32-byte ED25519 public key hex string WITH prefix",
			input:          "ED9434799226374926EDA3B54B1B461B4ABF7237962EAE18528FEA67595397FA32",
			inputPrefix:    []byte{AccountAddressPrefix},
			expectedOutput: "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
			expectedErr:    nil,
		},
		{
			description:    "Successfully generate address from a 32-byte ED25519 public key hex string WITHOUT prefix",
			input:          "9434799226374926EDA3B54B1B461B4ABF7237962EAE18528FEA67595397FA32",
			inputPrefix:    []byte{AccountAddressPrefix},
			expectedOutput: "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
			expectedErr:    nil,
		},
		{
			description:    "Successfully generate address from randomly generated 33-byte hex string",
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
			expectedErr:    &EncodeLengthError{Instance: "PublicKey", Input: 1, Expected: 33},
		},
		{
			description:    "Valid Public Key, invalid Type Prefix length",
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
				assert.EqualError(t, err, tc.expectedErr.Error())
			} else {
				if assert.NoError(t, err) {
					assert.NoError(t, err)
					assert.Equal(t, tc.expectedOutput, got)
				}
			}
		})
	}
}

func TestEncodeNodePublicKey(t *testing.T) {

}

func TestEncodeSeed(t *testing.T) {
	tt := []struct {
		description       string
		input             []byte
		inputEncodingType CryptoAlgorithm
		expectedOutput    string
		expectedErr       error
	}{
		{
			description:       "successful encode - ED25519",
			input:             []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			inputEncodingType: ED25519,
			expectedOutput:    "E2GEWzC8MMH3E2wKHAGWdVrTbtcWC",
			expectedErr:       nil,
		},
		{
			description:       "successful encode - SECP256K1",
			input:             []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			inputEncodingType: SECP256K1,
			expectedOutput:    "zh5iEuYTaHW4JwgCadsVQRmsfzUB",
			expectedErr:       nil,
		},
		{
			description:       "unsuccessful encode - invalid entropy length",
			input:             []byte{0x00},
			inputEncodingType: ED25519,
			expectedOutput:    "",
			expectedErr:       &EncodeLengthError{Instance: "Entropy", Input: len([]byte{0x00}), Expected: FamilySeedLength},
		},
		{
			description:       "unsuccessful encode - invalid encoding type",
			input:             []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			inputEncodingType: Undefined,
			expectedOutput:    "",
			expectedErr:       errors.New("encoding type must be `ed25519` or `secp256k1`"),
		},
		{
			description:       "invalid CryptoAlgorithm Uint type returns err",
			input:             []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			inputEncodingType: CryptoAlgorithm(255),
			expectedOutput:    "",
			expectedErr:       errors.New("encoding type must be `ed25519` or `secp256k1`"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			got, err := EncodeSeed(tc.input, tc.inputEncodingType)

			if tc.expectedErr != nil {
				assert.EqualError(t, err, tc.expectedErr.Error())
			} else {
				assert.Equal(t, tc.expectedOutput, got)
			}
		})
	}
}

func TestDecodeSeed(t *testing.T) {
	tt := []struct {
		description       string
		input             string
		expectedOutput    []byte
		expectedAlgorithm CryptoAlgorithm
		expectedErr       error
	}{
		{
			description:       "successful decode",
			input:             "E2GEWzC8MMH3E2wKHAGWdVrTbtcWC",
			expectedOutput:    []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xcc, 0xf9, 0x3e, 0xfc},
			expectedAlgorithm: ED25519,
			expectedErr:       nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {

			got, algorithm, err := DecodeSeed(tc.input)

			if tc.expectedErr != nil {
				assert.EqualError(t, err, tc.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOutput, got)
				assert.Equal(t, tc.expectedAlgorithm, algorithm)
			}
		})
	}
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
				assert.EqualError(t, err, tc.expectedErr.Error())
				assert.Nil(t, tc.expectedPrefix, typePrefix)
				assert.Nil(t, tc.expectedAccountID, accountID)
			} else {
				assert.NoError(t, err)
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
				assert.False(t, IsValidClassicAddress(tc.input))
			} else {
				assert.True(t, IsValidClassicAddress(tc.input))
			}
		})
	}
}
