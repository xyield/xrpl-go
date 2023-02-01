package transactions

import (
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type CheckCash struct {
	BaseTx
	CheckID    Hash256
	Amount     CurrencyAmount `json:",omitempty"`
	DeliverMin CurrencyAmount `json:",omitempty"`
}

func (*CheckCash) TxType() TxType {
	return CheckCashTx
}
