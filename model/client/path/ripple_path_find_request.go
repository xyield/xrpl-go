package path

import (
	"encoding/json"
	"fmt"

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

func (r *RipplePathFindRequest) Validate() error {
	if err := r.SourceAccount.Validate(); err != nil {
		return fmt.Errorf("ripple path find source: %w", err)
	}
	if err := r.DestinationAccount.Validate(); err != nil {
		return fmt.Errorf("ripple path find destination: %w", err)
	}
	if err := r.DestinationAmount.Validate(); err != nil {
		return fmt.Errorf("ripple path find destination amount: %w", err)
	}
	if r.SendMax != nil && len(r.SourceCurrencies) != 0 {
		return fmt.Errorf("ripple path find cannot have send max and source currencies set simultaneously")
	}
	if r.SendMax != nil {
		if err := r.SendMax.Validate(); err != nil {
			return fmt.Errorf("ripple path find send max: %w", err)
		}
	}
	if len(r.SourceCurrencies) != 0 {
		for _, c := range r.SourceCurrencies {
			if err := c.Validate(); err != nil {
				return fmt.Errorf("ripple path find source currencies: %w", err)
			}
		}
	}
	return nil
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
