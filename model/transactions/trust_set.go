package transactions

import (
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type TrustSet struct {
	BaseTx
	LimitAmount CurrencyAmount
	QualityIn   uint `json:",omitempty"`
	QualityOut  uint `json:",omitempty"`
}

func (*TrustSet) TxType() TxType {
	return TrustSetTx
}
