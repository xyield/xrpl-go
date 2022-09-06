package keypairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
	addresscodec "github.com/xyield/xrpl-go/address-codec"
)

func TestED25519DeriveKeypair(t *testing.T) {
	e := &ed25519Alg{}
	tt := []struct {
		description string
		seed        string
		validator   bool
		expPubKey   string
		expPrivKey  string
		expErr      error
	}{
		{
			description: "Successfully derive keypair",
			seed:        "sEdTjrdnJaPE2NNjmavQqXQdrf71NiH",
			validator:   false,
			expPubKey:   "ED4924A9045FE5ED8B22BAA7B6229A72A287CCF3EA287AADD3A032A24C0F008FA6",
			expPrivKey:  "EDBB3ECA8985E1484FA6A28C4B30FB0042A2CC5DF3EC8DC37B5F3D126DDFD3CA14",
			expErr:      nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			ds, _, _ := addresscodec.DecodeSeed(tc.seed)
			priv, pub, err := e.deriveKeypair(ds, tc.validator)
			if tc.expErr != nil {
				assert.Nil(t, pub)
				assert.Nil(t, priv)
				assert.Error(t, err, tc.expErr.Error())
			} else {
				assert.Equal(t, tc.expPrivKey, priv)
				assert.Equal(t, tc.expPubKey, pub)
			}
		})
	}
}
