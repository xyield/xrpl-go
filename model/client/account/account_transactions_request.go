package account

import (
	. "github.com/xyield/xrpl-go/model/client/common"
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountTransactionsRequest struct {
	Account        Address     `json:"account"`
	LedgerIndexMin LedgerIndex `json:"ledger_index_min,omitempty"`
	LedgerIndexMax LedgerIndex `json:"ledger_index_max,omitempty"`
	LedgerHash     LedgerHash  `json:"ledger_hash,omitempty"`
	LedgerIndex    LedgerIndex `json:"ledger_index,omitempty"`
	Binary         bool        `json:"binary,omitempty"`
	Forward        bool        `json:"forward,omitempty"`
	Limit          int         `json:"limit,omitempty"`
	Marker         interface{} `json:"marker,omitempty"`
}
