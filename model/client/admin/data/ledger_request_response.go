package data

import (
	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/client/ledger"
)

type LedgerRequestResponse struct {
	ledger.LedgerHeader
	LedgerRequest
	Acquiring LedgerRequest `json:"acquiring"`
}

type LedgerRequest struct {
	Hash              common.LedgerHash `json:"hash,omitempty"`
	HaveHeader        bool              `json:"have_header"`
	HaveState         bool              `json:"have_state,omitempty"`
	HaveTransactions  bool              `json:"have_transactions,omitempty"`
	NeededStateHashes []string          `json:"needed_state_hashes,omitempty"`
}
