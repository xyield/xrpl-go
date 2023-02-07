package transactions

type TicketCreate struct {
	BaseTx
	TicketCount uint
}

func (*TicketCreate) TxType() TxType {
	return TicketCreateTx
}
