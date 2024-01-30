package account

import (
	"github.com/CreatureDev/xrpl-go/model/client/common"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type AccountChannelsResponse struct {
	Account     types.Address      `json:"account"`
	Channels    []ChannelResult    `json:"channels"`
	LedgerIndex common.LedgerIndex `json:"ledger_index,omitempty"`
	LedgerHash  common.LedgerHash  `json:"ledger_hash,omitempty"`
	Validated   bool               `json:"validated,omitempty"`
	Limit       int                `json:"limit,omitempty"`
	Marker      any                `json:"marker,omitempty"`
}
