package transactions

import (
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type SetRegularKey struct {
	BaseTx
	RegularKey Address `json:",omitempty"`
}

func (*SetRegularKey) TxType() TxType {
	return SetRegularKeyTx
}
