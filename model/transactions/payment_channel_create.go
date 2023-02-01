package transactions

type PaymentChannelCreate struct {
	BaseTx
	Amount         XrpCurrencyAmount
	Destination    Address
	SettleDelay    uint
	PublicKey      []byte
	CancelAfter    uint `json:",omitempty"`
	DestinationTag uint `json:",omitempty"`
}

func (*PaymentChannelCreate) TxType() TxType {
	return PaymentChannelCreateTx
}
