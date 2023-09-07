package path

import (
	"encoding/json"
	"fmt"

	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type NFTokenSellOffersRequest struct {
	NFTokenID   types.NFTokenID        `json:"nft_id"`
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex common.LedgerSpecifier `json:"ledger_index,omitempty"`
	Limit       int                    `json:"limit,omitempty"`
	Marker      any                    `json:"marker,omitempty"`
}

func (*NFTokenSellOffersRequest) Method() string {
	return "nft_sell_offers"
}

func (r *NFTokenSellOffersRequest) Validate() error {
	if r.NFTokenID == "" {
		return fmt.Errorf("nft sell offer missing token id")
	}
	return nil
}

func (r *NFTokenSellOffersRequest) UnmarshalJSON(data []byte) error {
	type borHelper struct {
		NFTokenID   types.NFTokenID   `json:"nft_id"`
		LedgerHash  common.LedgerHash `json:"ledger_hash,omitempty"`
		LedgerIndex json.RawMessage   `json:"ledger_index,omitempty"`
		Limit       int               `json:"limit,omitempty"`
		Marker      any               `json:"marker,omitempty"`
	}
	var h borHelper
	err := json.Unmarshal(data, &h)
	if err != nil {
		return err
	}
	*r = NFTokenSellOffersRequest{
		NFTokenID:  h.NFTokenID,
		LedgerHash: h.LedgerHash,
		Limit:      h.Limit,
		Marker:     h.Marker,
	}
	var i common.LedgerSpecifier
	i, err = common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = i
	return nil
}
