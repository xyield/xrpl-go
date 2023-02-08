package account

import (
	. "github.com/xyield/xrpl-go/model/client/common"
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountLinesReponse struct {
	Account            Address     `json:"account"`
	Lines              []TrustLine `json:"lines"`
	LedgerCurrentIndex LedgerIndex `json:"ledger_current_index,omitempty"`
	LedgerIndex        LedgerIndex `json:"ledger_index,omitempty"`
	LedgerHash         LedgerHash  `json:"ledger_hash,omitempty"`
	Marker             interface{} `json:"marker,omitempty"`
}
