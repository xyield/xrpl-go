package transactions

import (
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type EscrowCancel struct {
	BaseTx
	Owner         types.Address
	OfferSequence uint
}

func (*EscrowCancel) TxType() TxType {
	return EscrowCancelTx
}
