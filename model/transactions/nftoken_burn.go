package transactions

import (
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type NFTokenBurn struct {
	BaseTx
	NFTokenID types.NFTokenID
	Owner     types.Address `json:",omitempty"`
}

func (*NFTokenBurn) TxType() TxType {
	return NFTokenBurnTx
}
