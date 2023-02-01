package account

import (
	. "github.com/xyield/xrpl-go/model/client/common"
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountOffersRequest struct {
	Account     Address     `json:"address"`
	LedgerHash  LedgerHash  `json:"ledger_hash,omitempty"`
	LedgerIndex LedgerIndex `json:"ledger_index,omitempty"`
	Limit       int         `json:"limit,omitempty"`
	Marker      interface{} `json:"marker,omitempty"`
	Strict      bool        `json:"strict,omitempty"`
}
