package account

import (
	"github.com/xyield/xrpl-go/model/client/common"
	. "github.com/xyield/xrpl-go/model/ledger"
)

type AccountInfoResponse struct {
	AccountData        AccountRoot        `json:"account_data"`
	SignerLists        []SignerList       `json:"signer_lists,omitempty"`
	LedgerCurrentIndex common.LedgerIndex `json:"ledger_current_index,omitempty"`
	LedgerIndex        common.LedgerIndex `json:"ledger_index,omitempty"`
	QueueData          QueueData          `json:"queue_data,omitempty"`
	Validated          bool               `json:"validated"`
}
