package ledger

import (
	"encoding/json"

	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/ledger"
)

type LedgerRequest struct {
	LedgerHash   common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex  common.LedgerSpecifier `json:"ledger_index,omitempty"`
	Full         bool                   `json:"full,omitempty"`
	Accounts     bool                   `json:"accounts,omitempty"`
	Transactions bool                   `json:"transactions,omitempty"`
	OwnerFunds   bool                   `json:"owner_funds,omitempty"`
	Binary       bool                   `json:"binary,omitempty"`
	Queue        bool                   `json:"queue,omitempty"`
	Type         ledger.LedgerEntryType `json:"type,omitempty"`
}

func (*LedgerRequest) Method() string {
	return "ledger"
}

func (*LedgerRequest) Validate() error {
	return nil
}

func (r *LedgerRequest) UnmarshalJSON(data []byte) error {
	type lrHelper struct {
		LedgerHash   common.LedgerHash      `json:"ledger_hash,omitempty"`
		LedgerIndex  json.RawMessage        `json:"ledger_index,omitempty"`
		Full         bool                   `json:"full,omitempty"`
		Accounts     bool                   `json:"accounts,omitempty"`
		Transactions bool                   `json:"transactions,omitempty"`
		OwnerFunds   bool                   `json:"owner_funds,omitempty"`
		Binary       bool                   `json:"binary,omitempty"`
		Queue        bool                   `json:"queue,omitempty"`
		Type         ledger.LedgerEntryType `json:"type,omitempty"`
	}
	var h lrHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = LedgerRequest{
		LedgerHash:   h.LedgerHash,
		Full:         h.Full,
		Accounts:     h.Accounts,
		Transactions: h.Transactions,
		OwnerFunds:   h.OwnerFunds,
		Binary:       h.Binary,
		Queue:        h.Queue,
		Type:         h.Type,
	}

	i, err := common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = i
	return nil
}
