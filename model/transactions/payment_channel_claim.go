package transactions

import (
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type PaymentChannelClaim struct {
	BaseTx
	Channel   types.Hash256
	Balance   types.XRPCurrencyAmount `json:",omitempty"`
	Amount    types.XRPCurrencyAmount `json:",omitempty"`
	Signature []byte                  `json:",omitempty"`
	PublicKey []byte                  `json:",omitempty"`
}

func (*PaymentChannelClaim) TxType() TxType {
	return PaymentChannelClaimTx
}
