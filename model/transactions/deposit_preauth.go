package transactions

import (
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type DepositPreauth struct {
	BaseTx
	Authorize   Address `json:",omitempty"`
	Unauthorize Address `json:",omitempty"`
}

func (*DepositPreauth) TxType() TxType {
	return DepositPreauthTx
}
