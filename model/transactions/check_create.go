package transactions

type CheckCreate struct {
	BaseTx
	Destination    Address
	SendMax        CurrencyAmount
	DestinationTag uint    `json:",omitempty"`
	Expiration     uint    `json:",omitempty"`
	InvoiceID      Hash256 `json:",omitempty"`
}

func (*CheckCreate) TxType() TxType {
	return CheckCreateTx
}
