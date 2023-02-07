package transactions

type OfferCancel struct {
	BaseTx
	OfferSequence uint
}

func (*OfferCancel) TxType() TxType {
	return OfferCancelTx
}
