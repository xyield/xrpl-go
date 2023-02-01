package account

import (
	. "github.com/xyield/xrpl-go/model/client/common"
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountChannelsResponse struct {
	Account     Address         `json:"account"`
	Channels    []ChannelResult `json:"channels"`
	LedgerIndex LedgerIndex     `json:"ledger_index,omitempty"`
	LedgerHash  LedgerHash      `json:"ledger_hash,omitempty"`
	Validated   bool            `json:"validated,omitempty"`
	Limit       int             `json:"limit,omitempty"`
	Marker      interface{}     `json:"marker,omitempty"`
}
