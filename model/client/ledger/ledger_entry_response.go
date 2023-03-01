package ledger

import (
	"encoding/json"

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
	type lerHelper struct {
		Index       string             `json:"index"`
		LedgerIndex common.LedgerIndex `json:"ledger_index"`
		Node        json.RawMessage    `json:"node,omitempty"`
		NodeBinary  string             `json:"node_binary,omitempty"`
	}
	var h lerHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = LedgerEntryResponse{
		Index:       h.Index,
		LedgerIndex: h.LedgerIndex,
		NodeBinary:  h.NodeBinary,
	}
	obj, err := ledger.UnmarshalLedgerObject(h.Node)
	if err != nil {
		return err
	}
	r.Node = obj

	return nil
}
