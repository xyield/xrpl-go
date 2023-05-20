package path

import (
	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type DepositAuthorizedResponse struct {
	DepositAuthorized  bool               `json:"deposit_authorized"`
	DestinationAccount types.Address      `json:"destination_account"`
	LedgerHash         common.LedgerHash  `json:"ledger_hash,omitempty"`
	LedgerIndex        common.LedgerIndex `json:"ledger_index,omitempty"`
	LedgerCurrentIndex common.LedgerIndex `json:"ledger_current_index,omitempty"`
	SourceAccount      types.Address      `json:"source_account"`
	Validated          bool               `json:"validated,omitempty"`
}
