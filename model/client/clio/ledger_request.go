package clio

import (
	"encoding/json"

	"github.com/xyield/xrpl-go/model/client/common"
)

type LedgerRequest struct {
	LedgerHash   common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex  common.LedgerSpecifier `json:"ledger_index,omitempty"`
	Full         bool                   `json:"full"`
	Accounts     bool                   `json:"accounts"`
	Transactions bool                   `json:"transactions"`
	Expand       bool                   `json:"expand"`
	OwnerFunds   bool                   `json:"owner_funds"`
	Binary       bool                   `json:"binary"`
	Queue        bool                   `json:"queue"`
	Diff         bool                   `json:"diff"`
}

func (*LedgerRequest) Method() string {
	return "ledger"
}

func (*LedgerRequest) Validate() error {
	return nil
}

func (r *LedgerRequest) UnmarshalJSON(data []byte) error {
	type lrHelper struct {
		LedgerHash   common.LedgerHash `json:"ledger_hash"`
		LedgerIndex  json.RawMessage   `json:"ledger_index"`
		Full         bool              `json:"full"`
		Accounts     bool              `json:"accounts"`
		Transactions bool              `json:"transactions"`
		Expand       bool              `json:"expand"`
		OwnerFunds   bool              `json:"owner_funds"`
		Binary       bool              `json:"binary"`
		Queue        bool              `json:"queue"`
		Diff         bool              `json:"diff"`
	}
	var h lrHelper
	err := json.Unmarshal(data, &h)
	if err != nil {
		return err
	}
	*r = LedgerRequest{
		LedgerHash:   h.LedgerHash,
		Accounts:     h.Accounts,
		Full:         h.Full,
		Transactions: h.Transactions,
		Expand:       h.Expand,
		OwnerFunds:   h.OwnerFunds,
		Binary:       h.Binary,
		Queue:        h.Queue,
		Diff:         h.Diff,
	}
	r.LedgerIndex, err = common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}

	return nil
}
