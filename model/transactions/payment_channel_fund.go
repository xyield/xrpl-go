package transactions

import (
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type PaymentChannelFund struct {
	BaseTx
	Channel    types.Hash256
	Amount     types.XRPCurrencyAmount
	Expiration uint `json:",omitempty"`
}

func (*PaymentChannelFund) TxType() TxType {
	return PaymentChannelFundTx
}
