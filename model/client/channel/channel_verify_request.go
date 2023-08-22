package channel

import "github.com/xyield/xrpl-go/model/transactions/types"

type ChannelVerifyRequest struct {
	Amount    types.XRPCurrencyAmount `json:"amount"`
	ChannelID string                  `json:"channel_id"`
	PublicKey string                  `json:"public_key"`
	Signature string                  `json:"signature"`
}

func (*ChannelVerifyRequest) Method() string {
	return "channel_verify"
}

func (*ChannelVerifyRequest) Validate() error {
	return nil
}
