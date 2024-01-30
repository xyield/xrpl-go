package transactions

import (
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type SetRegularKey struct {
	BaseTx
	RegularKey types.Address `json:",omitempty"`
}

func (*SetRegularKey) TxType() TxType {
	return SetRegularKeyTx
}
