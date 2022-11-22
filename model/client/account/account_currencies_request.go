package account

import (
	. "github.com/xyield/xrpl-go/model/client/common"
	. "github.com/xyield/xrpl-go/model/transactions"
)

type AccountCurrenciesRequest struct {
	Account     Address     `json:"account"`
	LedgerHash  LedgerHash  `json:"ledger_hash,omitempty"`
	LedgerIndex LedgerIndex `json:"ledger_index,omitempty"`
	Strict      bool        `json:"strict,omitempty"`
}
