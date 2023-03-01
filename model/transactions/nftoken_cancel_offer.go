package transactions

import (
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type NFTokenCancelOffer struct {
	BaseTx
	NFTokenOffers []types.Hash256
}

func (*NFTokenCancelOffer) TxType() TxType {
	return NFTokenCancelOfferTx
}
