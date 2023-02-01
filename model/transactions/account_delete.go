package transactions

type AccountDelete struct {
	BaseTx
	Destination    Address
	DestinationTag uint `json:",omitempty"`
}

func (*AccountDelete) TxType() TxType {
	return AccountDeleteTx
}
