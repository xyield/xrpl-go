package account

import (
	. "github.com/xyield/xrpl-go/model/client/common"
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountTransactionsResponse struct {
	Account        Address              `json:"account"`
	LedgerIndexMin LedgerIndex          `json:"ledger_index_min"`
	LedgerIndexMax LedgerIndex          `json:"ledger_index_max"`
	Limit          int                  `json:"limit"`
	Marker         interface{}          `json:"marker,omitempty"`
	Transactions   []AccountTransaction `json:"transactions"`
	Validated      bool                 `json:"validated"`
}
