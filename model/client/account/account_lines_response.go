package account

import (
	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountLinesReponse struct {
	Account            types.Address      `json:"account"`
	Lines              []TrustLine        `json:"lines"`
	LedgerCurrentIndex common.LedgerIndex `json:"ledger_current_index,omitempty"`
	LedgerIndex        common.LedgerIndex `json:"ledger_index,omitempty"`
	LedgerHash         common.LedgerHash  `json:"ledger_hash,omitempty"`
	Marker             any                `json:"marker,omitempty"`
}
