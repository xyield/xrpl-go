package account

import (
	. "github.com/xyield/xrpl-go/model/client/common"
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountChannelsRequest struct {
	Account            Address     `json:"account"`
	DestinationAccount Address     `json:"destination_account"`
	LedgerIndex        LedgerIndex `json:"ledger_index,omitempty"`
	LedgerHash         LedgerHash  `json:"ledger_hash,omitempty"`
	Limit              int         `json:"limit,omitempty"`
	Marker             interface{} `json:"marker,omitempty"`
}
