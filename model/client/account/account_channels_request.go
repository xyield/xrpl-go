package account

import (
	"encoding/json"
	"errors"

	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountChannelsRequest struct {
	Account            types.Address          `json:"account"`
	DestinationAccount types.Address          `json:"destination_account,omitempty"`
	LedgerIndex        common.LedgerSpecifier `json:"ledger_index,omitempty"`
	LedgerHash         common.LedgerHash      `json:"ledger_hash,omitempty"`
	Limit              int                    `json:"limit,omitempty"`
	Marker             any                    `json:"marker,omitempty"`
}

func (*AccountChannelsRequest) Method() string {
	return "account_channels"
}

// Below mean struct satisfies paginated response interface
func (a *AccountChannelsRequest) SetMarker(m any) {
	a.Marker = m
}

// Validate method to be added to each request struct
func (a *AccountChannelsRequest) Validate() error {
	if a.Account == "" {
		return errors.New("no account ID specified")
	}

	return nil
}

func (r *AccountChannelsRequest) UnmarshalJSON(data []byte) error {
	type acrHelper struct {
		Account            types.Address     `json:"account"`
		DestinationAccount types.Address     `json:"destination_account"`
		LedgerIndex        json.RawMessage   `json:"ledger_index,omitempty"`
		LedgerHash         common.LedgerHash `json:"ledger_hash,omitempty"`
		Limit              int               `json:"limit,omitempty"`
		Marker             any               `json:"marker,omitempty"`
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

	i, err := common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = i
	return nil
}
