package transactions

type PaymentChannelFund struct {
	BaseTx
	Channel    Hash256
	Amount     XrpCurrencyAmount
	Expiration uint `json:",omitempty"`
}

func (*PaymentChannelFund) TxType() TxType {
	return PaymentChannelFundTx
}
