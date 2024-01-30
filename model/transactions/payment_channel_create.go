package transactions

import (
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type PaymentChannelCreate struct {
	BaseTx
	Amount         types.XRPCurrencyAmount
	Destination    types.Address
	SettleDelay    uint
	PublicKey      string
	CancelAfter    uint `json:",omitempty"`
	DestinationTag uint `json:",omitempty"`
}

func (*PaymentChannelCreate) TxType() TxType {
	return PaymentChannelCreateTx
}
