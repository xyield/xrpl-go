package path

import (
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type NFTokenBuyOffersResponse struct {
	NFTokenID types.NFTokenID `json:"nft_id"`
	Offers    []NFTokenOffer  `json:"offers"`
	Limit     int             `json:"limit,omitempty"`
	Marker    any             `json:"marker,omitempty"`
}
