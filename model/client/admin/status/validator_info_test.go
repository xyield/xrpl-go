package status

import (
	"testing"

	"github.com/xyield/xrpl-go/test"
)

func TestValidatorInfoResponse(t *testing.T) {
	s := ValidatorInfoResponse{
		Domain:       "mduo13.com",
		EphemeralKey: "n9KnrcCmL5psyKtk2KWP6jy14Hj4EXuZDg7XMdQJ9cSDoFSp53hu",
		Manifest:     "JAAAAAFxIe002KClGBUlRA7h5J2Y5B7Xdlxn1Z5OxY7ZC2UmqUIikHMhAkVIeB7McBf4NFsBceQQlScTVUWMdpYzwmvs115SUGDKdkcwRQIhAJnKfYWnPsBsATIIRfgkAAK+HE4zp8G8AmOPrHmLZpZAAiANiNECVQTKktoD7BEoEmS8jaFBNMgRdcG0dttPurCAGXcKbWR1bzEzLmNvbXASQPjO6wxOfhtWsJ6oMWBg8Rs5STAGvQV2ArI5MG3KbpFrNSMxbx630Ars9d9j1ORsUS5v1biZRShZfg9180JuZAo=",
		MasterKey:    "nHBk5DPexBjinXV8qHn7SEKzoxh2W92FxSbNTPgGtQYBzEF4msn9",
		Seq:          1,
	}

	j := `{
	"domain": "mduo13.com",
	"ephemeral_key": "n9KnrcCmL5psyKtk2KWP6jy14Hj4EXuZDg7XMdQJ9cSDoFSp53hu",
	"manifest": "JAAAAAFxIe002KClGBUlRA7h5J2Y5B7Xdlxn1Z5OxY7ZC2UmqUIikHMhAkVIeB7McBf4NFsBceQQlScTVUWMdpYzwmvs115SUGDKdkcwRQIhAJnKfYWnPsBsATIIRfgkAAK+HE4zp8G8AmOPrHmLZpZAAiANiNECVQTKktoD7BEoEmS8jaFBNMgRdcG0dttPurCAGXcKbWR1bzEzLmNvbXASQPjO6wxOfhtWsJ6oMWBg8Rs5STAGvQV2ArI5MG3KbpFrNSMxbx630Ars9d9j1ORsUS5v1biZRShZfg9180JuZAo=",
	"master_key": "nHBk5DPexBjinXV8qHn7SEKzoxh2W92FxSbNTPgGtQYBzEF4msn9",
	"seq": 1
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}

}
