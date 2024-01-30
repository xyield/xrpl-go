package channel

import (
	"fmt"

	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type ChannelAuthorizeRequest struct {
	ChannelID  string                  `json:"channel_id"`
	Secret     string                  `json:"secret,omitempty"`
	Seed       string                  `json:"seed,omitempty"`
	SeedHex    string                  `json:"seed_hex,omitempty"`
	Passphrase string                  `json:"passphrase,omitempty"`
	KeyType    string                  `json:"key_type,omitempty"`
	Amount     types.XRPCurrencyAmount `json:"amount"`
}

func (*ChannelAuthorizeRequest) Method() string {
	return "channel_authorize"
}

func (r *ChannelAuthorizeRequest) Validate() error {
	if r.ChannelID == "" {
		return fmt.Errorf("channel authorize request: missing channel id")
	}
	var set []string
	if r.Secret != "" {
		set = append(set, "secret")
	}
	if r.Seed != "" {
		set = append(set, "seed")
	}
	if r.SeedHex != "" {
		set = append(set, "seed_hex")
	}
	if r.Passphrase != "" {
		set = append(set, "passphrase")
	}
	if len(set) == 0 {
		return fmt.Errorf("channel authorize request: at least one of (secret, seed, seed_hex, passphrase) must be set")
	}
	if len(set) > 1 {
		return fmt.Errorf("channel authorize request: only one signing method required, currently set %v", set)
	}
	return nil
}

// do not allow secrets to be printed
func (c *ChannelAuthorizeRequest) Format(s fmt.State, v rune) {
	type fHelper struct {
		ChannelID string                  `json:"channel_id"`
		KeyType   string                  `json:"key_type,omitempty"`
		Amount    types.XRPCurrencyAmount `json:"amount"`
	}
	h := fHelper{
		ChannelID: c.ChannelID,
		KeyType:   c.KeyType,
		Amount:    c.Amount,
	}
	fmt.Fprintf(s, "%"+string(v), h)
}
