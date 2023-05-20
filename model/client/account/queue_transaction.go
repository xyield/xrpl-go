package account

import "github.com/xyield/xrpl-go/model/transactions/types"

type QueueTransaction struct {
	AuthChange    bool                    `json:"auth_change"`
	Fee           types.XRPCurrencyAmount `json:"fee,omitempty"`
	FeeLevel      types.XRPCurrencyAmount `json:"fee_level,omitempty"`
	MaxSpendDrops types.XRPCurrencyAmount `json:"max_spend_drops,omitempty"`
	Seq           int                     `json:"seq,omitempty"`
}
