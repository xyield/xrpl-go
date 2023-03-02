package transactions

import (
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type EscrowFinish struct {
	BaseTx
	Owner         types.Address
	OfferSequence uint
	Condition     []byte `json:",omitempty"`
	Fulfillment   []byte `json:",omitempty"`
}

func (*EscrowFinish) TxType() TxType {
	return EscrowFinishTx
}
