package transactions

import (
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type NFTokenAcceptOffer struct {
	BaseTx
	NFTokenSellOffer Hash256        `json:",omitempty"`
	NFTokenBuyOffer  Hash256        `json:",omitempty"`
	NFTokenBrokerFee CurrencyAmount `json:",omitempty"`
}

func (*NFTokenAcceptOffer) TxType() TxType {
	return NFTokenAcceptOfferTx
}
