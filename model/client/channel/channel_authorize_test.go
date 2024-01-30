package channel

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/CreatureDev/xrpl-go/test"
)

func TestChannelAuthorizeRequest(t *testing.T) {
	s := ChannelAuthorizeRequest{
		ChannelID: "5DB01B7FFED6B67E6B0414DED11E051D2EE2B7619CE0EAA6286D67A3A4D5BDB3",
		Seed:      "sabcdef123456abcdef12345",
		KeyType:   "secp256k1",
		Amount:    types.XRPCurrencyAmount(1000000),
	}

	j := `{
	"channel_id": "5DB01B7FFED6B67E6B0414DED11E051D2EE2B7619CE0EAA6286D67A3A4D5BDB3",
	"seed": "sabcdef123456abcdef12345",
	"key_type": "secp256k1",
	"amount": "1000000"
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestChannelAuthorizeResponse(t *testing.T) {
	s := ChannelAuthorizeResponse{
		Signature: "304402204EF0AFB78AC23ED1C472E74F4299C0C21F1B21D07EFC0A3838A420F76D783A400220154FB11B6F54320666E4C36CA7F686C16A3A0456800BBC43746F34AF50290064",
	}

	j := `{
	"signature": "304402204EF0AFB78AC23ED1C472E74F4299C0C21F1B21D07EFC0A3838A420F76D783A400220154FB11B6F54320666E4C36CA7F686C16A3A0456800BBC43746F34AF50290064"
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestChannelAuthorizeValidate(t *testing.T) {
	s := ChannelAuthorizeRequest{}
	err := s.Validate()
	assert.ErrorContains(t, err, "missing channel id")

	s = ChannelAuthorizeRequest{
		ChannelID: "abc",
	}
	err = s.Validate()
	assert.ErrorContains(t, err, "seed")
	s = ChannelAuthorizeRequest{
		ChannelID: "abc",
		Seed:      "123",
	}
	err = s.Validate()
	assert.Nil(t, err)
	s.Secret = "def"
	err = s.Validate()
	assert.ErrorContains(t, err, "seed")
	assert.ErrorContains(t, err, "secret")
}
