package path

import (
	"encoding/json"

	"github.com/CreatureDev/xrpl-go/model/transactions"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type PathFindResponse struct {
	Alternatives       []PathAlternative    `json:"alternatives"`
	DestinationAccount types.Address        `json:"destination_account"`
	DestinationAmount  types.CurrencyAmount `json:"destination_amount"`
	SourceAccount      types.Address        `json:"source_account"`
	FullReply          bool                 `json:"full_reply"`
	Closed             bool                 `json:"closed,omitempty"`
	Status             bool                 `json:"status,omitempty"`
}

func (r *PathFindResponse) UnmarshalJSON(data []byte) error {
	type pfrHelper struct {
		Alternatives       []PathAlternative `json:"alternatives"`
		DestinationAccount types.Address     `json:"destination_account"`
		DestinationAmount  json.RawMessage   `json:"destination_amount"`
		SourceAccount      types.Address     `json:"source_account"`
		FullReply          bool              `json:"full_reply"`
		Closed             bool              `json:"closed,omitempty"`
		Status             bool              `json:"status,omitempty"`
	}
	var h pfrHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = PathFindResponse{
		Alternatives:       h.Alternatives,
		DestinationAccount: h.DestinationAccount,
		SourceAccount:      h.SourceAccount,
		FullReply:          h.FullReply,
		Closed:             h.Closed,
		Status:             h.Status,
	}

	dst, err := types.UnmarshalCurrencyAmount(h.DestinationAmount)
	if err != nil {
		return err
	}
	r.DestinationAmount = dst

	return nil
}

type PathAlternative struct {
	PathsComputed     [][]transactions.PathStep `json:"paths_computed"`
	SourceAmount      types.CurrencyAmount      `json:"source_amount"`
	DestinationAmount types.CurrencyAmount      `json:"destination_amount,omitempty"`
}

func (p *PathAlternative) UnmarshalJSON(data []byte) error {
	type paHelper struct {
		PathsComputed     [][]transactions.PathStep `json:"paths_computed"`
		SourceAmount      json.RawMessage           `json:"source_amount"`
		DestinationAmount json.RawMessage           `json:"destination_amount,omitempty"`
	}
	var h paHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	p.PathsComputed = h.PathsComputed

	var src, dst types.CurrencyAmount
	var err error

	src, err = types.UnmarshalCurrencyAmount(h.SourceAmount)
	if err != nil {
		return err
	}
	p.SourceAmount = src

	dst, err = types.UnmarshalCurrencyAmount(h.DestinationAmount)
	if err != nil {
		return err
	}
	p.DestinationAmount = dst

	return nil
}
