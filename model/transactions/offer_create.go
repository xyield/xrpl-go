package transactions

type OfferCreate struct {
	BaseTx
	Expiration    uint `json:",omitempty"`
	OfferSequence uint `json:",omitempty"`
	TakerGets     CurrencyAmount
	TakerPays     CurrencyAmount
}

func (*OfferCreate) TxType() TxType {
	return OfferCreateTx
}
