package transactions

import (
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type PaymentChannelFund struct {
	BaseTx
	Channel    Hash256
	Amount     XRPCurrencyAmount
	Expiration uint `json:",omitempty"`
}

func (*PaymentChannelFund) TxType() TxType {
	return PaymentChannelFundTx
}
