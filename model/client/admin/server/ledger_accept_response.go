package server

import "github.com/xyield/xrpl-go/model/client/common"

type LedgerAcceptResponse struct {
	LedgerCurrentIndex common.LedgerIndex `json:"ledger_current_index"`
}
