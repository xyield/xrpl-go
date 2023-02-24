package account

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/client/common"
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountChannelsRequest struct {
	Account            Address         `json:"account"`
	DestinationAccount Address         `json:"destination_account,omitempty"`
	LedgerIndex        LedgerSpecifier `json:"ledger_index,omitempty"`
	LedgerHash         LedgerHash      `json:"ledger_hash,omitempty"`
	Limit              int             `json:"limit,omitempty"`
	Marker             any             `json:"marker,omitempty"`
}

func (*AccountChannelsRequest) Method() string {
	return "account_channels"
}

func (r *AccountChannelsRequest) UnmarshalJSON(data []byte) error {
	type acrHelper struct {
		Account            Address         `json:"account"`
		DestinationAccount Address         `json:"destination_account"`
		LedgerIndex        json.RawMessage `json:"ledger_index,omitempty"`
		LedgerHash         LedgerHash      `json:"ledger_hash,omitempty"`
		Limit              int             `json:"limit,omitempty"`
		Marker             any             `json:"marker,omitempty"`
	}
	var h acrHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = AccountChannelsRequest{
		Account:            h.Account,
		DestinationAccount: h.DestinationAccount,
		LedgerHash:         h.LedgerHash,
		Limit:              h.Limit,
		Marker:             h.Marker,
	}

	i, err := UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = i
	return nil
}
