package transactions

import (
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type PaymentChannelClaim struct {
	BaseTx
	Channel   Hash256
	Balance   XrpCurrencyAmount `json:",omitempty"`
	Amount    XrpCurrencyAmount `json:",omitempty"`
	Signature []byte            `json:",omitempty"`
	PublicKey []byte            `json:",omitempty"`
}

func (*PaymentChannelClaim) TxType() TxType {
	return PaymentChannelClaimTx
}
