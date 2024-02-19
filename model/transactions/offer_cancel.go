package transactions

type OfferCancel struct {
	BaseTx
	OfferSequence uint32
}

func (*OfferCancel) TxType() TxType {
	return OfferCancelTx
}
