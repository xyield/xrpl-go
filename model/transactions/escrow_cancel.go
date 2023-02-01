package transactions

type EscrowCancel struct {
	BaseTx
	Owner         Address
	OfferSequence uint
}

func (*EscrowCancel) TxType() TxType {
	return EscrowCancelTx
}
