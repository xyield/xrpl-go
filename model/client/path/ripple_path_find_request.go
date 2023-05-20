package path

import (
	"encoding/json"

	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type RipplePathFindRequest struct {
	SourceAccount      types.Address                `json:"source_account"`
	DestinationAccount types.Address                `json:"destination_account"`
	DestinationAmount  types.CurrencyAmount         `json:"destination_amount"`
	SendMax            types.CurrencyAmount         `json:"send_max,omitempty"`
	SourceCurrencies   []types.IssuedCurrencyAmount `json:"source_currencies,omitempty"`
	LedgerHash         common.LedgerHash            `json:"ledger_hash,omitempty"`
	LedgerIndex        common.LedgerSpecifier       `json:"ledger_index,omitempty"`
}

func (*RipplePathFindRequest) Method() string {
	return "ripple_path_find"
}

func (r *RipplePathFindRequest) UnmarshalJSON(data []byte) error {
	type rpfHelper struct {
		SourceAccount      types.Address                `json:"source_account"`
		DestinationAccount types.Address                `json:"destination_account"`
		DestinationAmount  json.RawMessage              `json:"destination_amount"`
		SendMax            json.RawMessage              `json:"send_max,omitempty"`
		SourceCurrencies   []types.IssuedCurrencyAmount `json:"source_currencies,omitempty"`
		LedgerHash         common.LedgerHash            `json:"ledger_hash,omitempty"`
		LedgerIndex        json.RawMessage              `json:"ledger_index,omitempty"`
	}
	var h rpfHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = RipplePathFindRequest{
		SourceAccount:      h.SourceAccount,
		DestinationAccount: h.DestinationAccount,
		SourceCurrencies:   h.SourceCurrencies,
		LedgerHash:         h.LedgerHash,
	}
	var dst, max types.CurrencyAmount
	var err error

	dst, err = types.UnmarshalCurrencyAmount(h.DestinationAmount)
	if err != nil {
		return err
	}
	r.DestinationAmount = dst

	max, err = types.UnmarshalCurrencyAmount(h.SendMax)
	if err != nil {
		return err
	}
	r.SendMax = max

	var i common.LedgerSpecifier
	i, err = common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = i
	return nil

}
