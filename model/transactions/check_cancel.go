package transactions

import (
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type CheckCancel struct {
	BaseTx
	CheckID types.Hash256
}

func (*CheckCancel) TxType() TxType {
	return CheckCancelTx
}
