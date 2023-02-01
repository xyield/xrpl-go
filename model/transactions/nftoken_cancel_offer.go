package transactions

type NFTokenCancelOffer struct {
	BaseTx
	NFTokenOffer []Hash256
}

func (*NFTokenCancelOffer) TxType() TxType {
	return NFTokenCancelOfferTx
}
