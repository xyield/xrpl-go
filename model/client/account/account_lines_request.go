package account

import (
	. "github.com/xyield/xrpl-go/model/client/common"
	. "github.com/xyield/xrpl-go/model/transactions"
)

type AccountLinesRequest struct {
	Account     Address     `json:"account"`
	LedgerHash  LedgerHash  `json:"ledger_hash,omitempty"`
	LedgerIndex LedgerIndex `json:"ledger_index,omitempty"`
	Peer        Address     `json:"peer,omitempty"`
	Limit       int         `json:"limit,omitempty"`
	Marker      interface{} `json:"marker,omitempty"`
}
