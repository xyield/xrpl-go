package clio

import (
	"encoding/json"
	"fmt"

	"github.com/CreatureDev/xrpl-go/model/client/common"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type NFTHistoryRequest struct {
	NFTokenID      types.NFTokenID        `json:"nft_id"`
	LedgerIndexMin common.LedgerIndex     `json:"ledger_index_min,omitempty"`
	LedgerIndexMax common.LedgerIndex     `json:"ledger_index_max,omitempty"`
	LedgerHash     common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex    common.LedgerSpecifier `json:"ledger_index,omitempty"`
	Binary         bool                   `json:"binary,omitempty"`
	Forward        bool                   `json:"forward,omitempty"`
	Limit          uint32                 `json:"limit,omitempty"`
	Marker         any                    `json:"marker,omitempty"`
}

func (*NFTHistoryRequest) Method() string {
	return "nft_history"
}

func (r *NFTHistoryRequest) Validate() error {
	if err := r.NFTokenID.Validate(); err != nil {
		return fmt.Errorf("nft history request: %w", err)
	}
	return nil
}

func (r *NFTHistoryRequest) UnmarshalJSON(data []byte) error {
	type nhrHelper struct {
		NFTokenID      types.NFTokenID    `json:"nft_id"`
		LedgerIndexMin common.LedgerIndex `json:"ledger_index_min,omitempty"`
		LedgerIndexMax common.LedgerIndex `json:"ledger_index_max,omitempty"`
		LedgerHash     common.LedgerHash  `json:"ledger_hash,omitempty"`
		LedgerIndex    json.RawMessage    `json:"ledger_index,omitempty"`
		Binary         bool               `json:"binary,omitempty"`
		Forward        bool               `json:"forward,omitempty"`
		Limit          uint32             `json:"limit,omitempty"`
		Marker         any                `json:"marker,omitempty"`
	}
	var h nhrHelper
	err := json.Unmarshal(data, &h)
	if err != nil {
		return err
	}
	*r = NFTHistoryRequest{
		NFTokenID:      h.NFTokenID,
		LedgerIndexMin: h.LedgerIndexMin,
		LedgerIndexMax: h.LedgerIndexMax,
		LedgerHash:     h.LedgerHash,
		Binary:         h.Binary,
		Forward:        h.Forward,
		Limit:          h.Limit,
		Marker:         h.Marker,
	}
	r.LedgerIndex, err = common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	return err
}
