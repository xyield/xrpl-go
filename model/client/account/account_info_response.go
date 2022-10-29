package account

import (
	. "github.com/xyield/xrpl-go/model/ledger"
)

type AccountInfoResponse struct {
	AccountData        AccountRoot  `json:"account_data"`
	SignerLists        []SignerList `json:"signer_lists,omitempty"`
	LedgerCurrentIndex uint64       `json:"ledger_current_index"`
	QueueData          QueueData    `json:"queue_data"`
	Validated          bool         `json:"validated"`
}
