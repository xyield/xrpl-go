package ledger

import (
	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/ledger"
)

type LedgerEntryResponse struct {
	Index       string              `json:"index"`
	LedgerIndex common.LedgerIndex  `json:"ledger_index"`
	Node        ledger.LedgerObject `json:"node,omitempty"`
	NodeBinary  string              `json:"node_binary,omitempty"`
}

func (r *LedgerEntryResponse) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
}
