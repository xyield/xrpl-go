package transactions

import (
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type EscrowCancel struct {
	BaseTx
	Owner         Address
	OfferSequence uint
}

func (*EscrowCancel) TxType() TxType {
	return EscrowCancelTx
}
