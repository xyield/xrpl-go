package transactions

import (
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type EscrowFinish struct {
	BaseTx
	Owner         types.Address
	OfferSequence uint
	Condition     string `json:",omitempty"`
	Fulfillment   string `json:",omitempty"`
}

func (*EscrowFinish) TxType() TxType {
	return EscrowFinishTx
}
