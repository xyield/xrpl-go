package transactions

import (
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type NFTokenBurn struct {
	BaseTx
	NFTokenID NFTokenID
	Owner     Address `json:",omitempty"`
}

func (*NFTokenBurn) TxType() TxType {
	return NFTokenBurnTx
}
