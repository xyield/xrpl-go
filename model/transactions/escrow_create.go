package transactions

import (
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type EscrowCreate struct {
	BaseTx
	Amount         XRPCurrencyAmount
	Destination    Address
	CancelAfter    uint   `json:",omitempty"`
	FinishAfter    uint   `json:",omitempty"`
	Condition      []byte `json:",omitempty"`
	DestinationTag uint   `json:",omitempty"`
}

func (*EscrowCreate) TxType() TxType {
	return EscrowCreateTx
}
