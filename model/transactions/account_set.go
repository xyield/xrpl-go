package transactions

import (
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type AccountSet struct {
	BaseTx
	ClearFlag     uint32        `json:",omitempty"`
	Domain        string        `json:",omitempty"`
	EmailHash     types.Hash128 `json:",omitempty"`
	MessageKey    string        `json:",omitempty"`
	NFTokenMinter string        `json:",omitempty"`
	SetFlag       uint32        `json:",omitempty"`
	TransferRate  uint          `json:",omitempty"`
	TickSize      uint8         `json:",omitempty"`
	WalletLocator types.Hash256 `json:",omitempty"`
	WalletSize    uint32        `json:",omitempty"`
}

func (*AccountSet) TxType() TxType {
	return AccountSetTx
}
