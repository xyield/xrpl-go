package ledger

import "github.com/xyield/xrpl-go/model/client/common"

type LedgerCurrentResponse struct {
	LedgerCurrentIndex common.LedgerIndex `json:"ledger_current_index"`
}
