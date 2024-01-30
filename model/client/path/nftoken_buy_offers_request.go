package path

import (
	"encoding/json"
	"fmt"

	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type NFTokenBuyOffersRequest struct {
	NFTokenID   types.NFTokenID        `json:"nft_id"`
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex common.LedgerSpecifier `json:"ledger_index,omitempty"`
	Limit       int                    `json:"limit,omitempty"`
	Marker      any                    `json:"marker,omitempty"`
}

func (*NFTokenBuyOffersRequest) Method() string {
	return "nft_buy_offers"
}

func (r *NFTokenBuyOffersRequest) Validat() error {
	if r.NFTokenID == "" {
		return fmt.Errorf("nft buy offer missing token id")
	}
	return nil
}

func (r *NFTokenBuyOffersRequest) UnmarshalJSON(data []byte) error {
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
	*r = NFTokenBuyOffersRequest{
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
