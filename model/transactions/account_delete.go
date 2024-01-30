package transactions

import (
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type AccountDelete struct {
	BaseTx
	Destination    types.Address
	DestinationTag uint `json:",omitempty"`
}

func (*AccountDelete) TxType() TxType {
	return AccountDeleteTx
}
