package transactions

import (
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type NFTokenCreateOffer struct {
	BaseTx
	Owner       Address `json:",omitempty"`
	NFTokenID   NFTokenID
	Amount      CurrencyAmount
	Expiration  uint    `json:",omitempty"`
	Destination Address `json:",omitempty"`
}

func (*NFTokenCreateOffer) TxType() TxType {
	return NFTokenCreateOfferTx
}
