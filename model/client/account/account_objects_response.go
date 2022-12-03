package account

import (
	. "github.com/xyield/xrpl-go/model/client/common"
	. "github.com/xyield/xrpl-go/model/ledger"
	. "github.com/xyield/xrpl-go/model/transactions"
)

type AccountObjectsResponse struct {
	Account            Address        `json:"account"`
	AccountObjects     []LedgerObject `json:"account_objects"`
	LedgerHash         LedgerHash     `json:"ledger_hash"`
	LedgerIndex        LedgerIndex    `json:"ledger_index"`
	LedgerCurrentIndex LedgerIndex    `json:"ledger_current_index"`
	Limit              int            `json:"limit"`
	Marker             interface{}    `json:"marker"`
	Validated          bool           `json:"validated"`
}
