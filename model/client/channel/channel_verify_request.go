package channel

import (
	"fmt"

	"github.com/xyield/xrpl-go/model/transactions/types"
)

type ChannelVerifyRequest struct {
	Amount    types.XRPCurrencyAmount `json:"amount"`
	ChannelID types.Hash256           `json:"channel_id"`
	PublicKey string                  `json:"public_key"`
	Signature string                  `json:"signature"`
}

func (*ChannelVerifyRequest) Method() string {
	return "channel_verify"
}

func (r *ChannelVerifyRequest) Validate() error {
	if err := r.ChannelID.Validate(); err != nil {
		return fmt.Errorf("channel verify request: channel id: %w", err)
	}
	if r.PublicKey == "" {
		return fmt.Errorf("channel verify request: public key not set")
	}
	if r.Signature == "" {
		return fmt.Errorf("channel verify request: signature not set")
	}
	return nil
}
