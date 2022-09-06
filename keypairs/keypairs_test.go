package keypairs

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	addresscodec "github.com/xyield/xrpl-go/address-codec"
)

func TestGenerateEncodeSeed(t *testing.T) {
	fb := bytes.NewBuffer([]byte("fakeRandomString"))
	tr := randomizer{
		fb,
	}
	r = tr
	tt := []struct {
		description string
		entropy     string
		algorithm   addresscodec.CryptoAlgorithm
		expected    string
		expectedErr error
	}{
		{
			description: "Empty entropy should generate random seed (ED25519)",
			entropy:     "",
			algorithm:   addresscodec.ED25519,
			expected:    "sEdTjrdnJaPE2NNjmavQqXQdrf71NiH",
			expectedErr: nil,
		},
		{
			description: "Entropy defined and above family seed length (ED25519)",
			entropy:     "setPasswordOverLen16",
			algorithm:   addresscodec.ED25519,
			expected:    "sEdTuXdrgQobjDidph2oMDN36jGZX2U",
			expectedErr: nil,
		},
		{
			description: "Empty entropy should generate random seed (SECP256K1)",
			entropy:     "",
			algorithm:   addresscodec.SECP256K1,
			expected:    "sh3pdwcaoo7vt5rtrEZJ7a75LnDo3",
			expectedErr: nil,
		},
		{
			description: "Entropy defined and above family seed length (SECP256K1)",
			entropy:     "setPasswordOverLen16",
			algorithm:   addresscodec.SECP256K1,
			expected:    "shJYdazRN9dvWbGqCehzHcBKWBaFR",
			expectedErr: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			if tc.entropy == "" {
				fb := bytes.NewBuffer([]byte("fakeRandomString"))
				tr := randomizer{
					fb,
				}
				r = tr
			}
			a, err := GenerateSeed(tc.entropy, tc.algorithm)

			if tc.expectedErr != nil {
				assert.Zero(t, a)
				assert.Error(t, err, tc.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, a)
			}
		})
	}
}
