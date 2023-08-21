package account

import (
	"encoding/json"
	"fmt"

	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountLinesRequest struct {
	Account     types.Address          `json:"account"`
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex common.LedgerSpecifier `json:"ledger_index,omitempty"`
	Peer        types.Address          `json:"peer,omitempty"`
	Limit       int                    `json:"limit,omitempty"`
	Marker      any                    `json:"marker,omitempty"`
}

func (*AccountLinesRequest) Method() string {
	return "account_lines"
}

func (r *AccountLinesRequest) Validate() error {
	if err := r.Account.Validate(); err != nil {
		return err
	}

	if r.Limit != 0 && (r.Limit < 10 || r.Limit > 400) {
		return fmt.Errorf("invalid limit, must be 10 <= limit <= 400")
	}

	return nil
}
func (r *AccountLinesRequest) UnmarshalJSON(data []byte) error {
	type alrHelper struct {
		Account     types.Address     `json:"account"`
		LedgerHash  common.LedgerHash `json:"ledger_hash,omitempty"`
		LedgerIndex json.RawMessage   `json:"ledger_index,omitempty"`
		Peer        types.Address     `json:"peer,omitempty"`
		Limit       int               `json:"limit,omitempty"`
		Marker      any               `json:"marker,omitempty"`
	}
	var h alrHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = AccountLinesRequest{
		Account:    h.Account,
		LedgerHash: h.LedgerHash,
		Peer:       h.Peer,
		Limit:      h.Limit,
		Marker:     h.Marker,
	}

	i, err := common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = i
	return nil
}
