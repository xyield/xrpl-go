package account

import (
	. "github.com/xyield/xrpl-go/model/client/common"
	. "github.com/xyield/xrpl-go/model/transactions"
)

type AccountLinesReponse struct {
	Account            Address     `json:"account"`
	Lines              []TrustLine `json:"lines"`
	LedgerCurrentIndex LedgerIndex `json:"ledger_current_index"`
	LedgerIndex        LedgerIndex `json:"ledger_index"`
	LedgerHash         LedgerHash  `json:"ledger_hash"`
	Marker             interface{} `json:"marker"`
}
