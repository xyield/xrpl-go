package transactions

type EscrowCreate struct {
	BaseTx
	Amount         XrpCurrencyAmount
	Destination    Address
	CancelAfter    uint   `json:",omitempty"`
	FinishAfter    uint   `json:",omitempty"`
	Condition      []byte `json:",omitempty"`
	DestinationTag uint   `json:",omitempty"`
}

func (*EscrowCreate) TxType() TxType {
	return EscrowCreateTx
}
