package account

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/client/common"
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountLinesRequest struct {
	Account     Address         `json:"account"`
	LedgerHash  LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex LedgerSpecifier `json:"ledger_index,omitempty"`
	Peer        Address         `json:"peer,omitempty"`
	Limit       int             `json:"limit,omitempty"`
	Marker      any             `json:"marker,omitempty"`
}

func (*AccountLinesRequest) Method() string {
	return "account_lines"
}

func (r *AccountLinesRequest) UnmarshalJSON(data []byte) error {
	type alrHelper struct {
		Account     Address         `json:"account"`
		LedgerHash  LedgerHash      `json:"ledger_hash,omitempty"`
		LedgerIndex json.RawMessage `json:"ledger_index,omitempty"`
		Peer        Address         `json:"peer,omitempty"`
		Limit       int             `json:"limit,omitempty"`
		Marker      any             `json:"marker,omitempty"`
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

	i, err := UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = i
	return nil
}
