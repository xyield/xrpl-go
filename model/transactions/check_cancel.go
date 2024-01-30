package transactions

import (
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type CheckCancel struct {
	BaseTx
	CheckID types.Hash256
}

func (*CheckCancel) TxType() TxType {
	return CheckCancelTx
}
