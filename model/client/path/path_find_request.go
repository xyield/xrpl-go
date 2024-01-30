package path

import (
	"encoding/json"
	"fmt"

	"github.com/CreatureDev/xrpl-go/model/transactions"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type PathSubCommand string

const (
	CREATE PathSubCommand = "create"
	CLOSE  PathSubCommand = "close"
	STATUS PathSubCommand = "status"
)

type PathFindRequest struct {
	Subcommand         PathSubCommand          `json:"subcommand"`
	SourceAccount      types.Address           `json:"source_account,omitempty"`
	DestinationAccount types.Address           `json:"destination_account,omitempty"`
	DestinationAmount  types.CurrencyAmount    `json:"destination_amount,omitempty"`
	SendMax            types.CurrencyAmount    `json:"send_max,omitempty"`
	Paths              []transactions.PathStep `json:"paths,omitempty"`
}

func (*PathFindRequest) Method() string {
	return "path_find"
}

func (r *PathFindRequest) Validate() error {
	switch r.Subcommand {
	case CREATE:
		if err := r.SourceAccount.Validate(); err != nil {
			return fmt.Errorf("path find create source: %w", err)
		}
		if err := r.DestinationAccount.Validate(); err != nil {
			return fmt.Errorf("path find create destination: %w", err)
		}
		if err := r.DestinationAmount.Validate(); err != nil {
			return fmt.Errorf("path find create destination amount: %w", err)
		}
		if r.SendMax != nil {
			if err := r.SendMax.Validate(); err != nil {
				return fmt.Errorf("path find create send max: %w", err)
			}
		}
		return nil
	case CLOSE, STATUS:
		return nil
	default:
		return fmt.Errorf("path find: invalid subcommand")
	}
}

func (r *PathFindRequest) UnmarshalJSON(data []byte) error {
	type pfrHelper struct {
		Subcommand         PathSubCommand          `json:"subcommand"`
		SourceAccount      types.Address           `json:"source_account,omitempty"`
		DestinationAccount types.Address           `json:"destination_account,omitempty"`
		DestinationAmount  json.RawMessage         `json:"destination_amount,omitempty"`
		SendMax            json.RawMessage         `json:"send_max,omitempty"`
		Paths              []transactions.PathStep `json:"paths,omitempty"`
	}
	var h pfrHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = PathFindRequest{
		Subcommand:         h.Subcommand,
		SourceAccount:      h.SourceAccount,
		DestinationAccount: h.DestinationAccount,
		Paths:              h.Paths,
	}

	var dest, max types.CurrencyAmount
	var err error
	dest, err = types.UnmarshalCurrencyAmount(h.DestinationAmount)
	if err != nil {
		return err
	}
	r.DestinationAmount = dest

	max, err = types.UnmarshalCurrencyAmount(h.SendMax)
	if err != nil {
		return err
	}
	r.SendMax = max

	return nil
}
