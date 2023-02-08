package account

import (
	. "github.com/xyield/xrpl-go/model/client/common"
)

type AccountCurrenciesResponse struct {
	LedgerHash        LedgerHash  `json:"ledger_hash,omitempty"`
	LedgerIndex       LedgerIndex `json:"ledger_index"`
	ReceiveCurrencies []string    `json:"receive_currencies"`
	SendCurrencies    []string    `json:"send_currencies"`
	Validated         bool        `json:"validated"`
}
