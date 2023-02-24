package account

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/client/common"
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountOffersRequest struct {
	Account     Address         `json:"account"`
	LedgerHash  LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex LedgerSpecifier `json:"ledger_index,omitempty"`
	Limit       int             `json:"limit,omitempty"`
	Marker      any             `json:"marker,omitempty"`
	Strict      bool            `json:"strict,omitempty"`
}

func (*AccountOffersRequest) Method() string {
	return "account_offers"
}

func (r *AccountOffersRequest) UnmarshalJSON(data []byte) error {
	type aorHelper struct {
		Account     Address         `json:"account"`
		LedgerHash  LedgerHash      `json:"ledger_hash,omitempty"`
		LedgerIndex json.RawMessage `json:"ledger_index,omitempty"`
		Limit       int             `json:"limit,omitempty"`
		Marker      any             `json:"marker,omitempty"`
		Strict      bool            `json:"strict,omitempty"`
	}
	var h aorHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = AccountOffersRequest{
		Account:    h.Account,
		LedgerHash: h.LedgerHash,
		Limit:      h.Limit,
		Marker:     h.Marker,
		Strict:     h.Strict,
	}

	i, err := UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = i
	return nil
}
