package ledger

import "github.com/xyield/xrpl-go/model/client/common"

type LedgerClosedResponse struct {
	LedgerHash  string             `json:"ledger_hash"`
	LedgerIndex common.LedgerIndex `json:"ledger_index"`
}
