package transactions

import (
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type EscrowFinish struct {
	BaseTx        `mapstructure:",squash"`
	Owner         types.Address
	OfferSequence uint32
	Condition     string `json:",omitempty"`
	Fulfillment   string `json:",omitempty"`
}

func (*EscrowFinish) TxType() TxType {
	return EscrowFinishTx
}
