package transactions

type AccountSet struct {
	BaseTx
	ClearFlag     uint    `json:",omitempty"`
	Domain        []byte  `json:",omitempty"`
	EmailHash     Hash128 `json:",omitempty"`
	MessageKey    []byte  `json:",omitempty"`
	NFTokenMinter []byte  `json:",omitempty"`
	SetFlag       uint    `json:",omitempty"`
	TransferRate  uint    `json:",omitempty"`
	TickSize      uint8   `json:",omitempty"`
	WalletLocator Hash256 `json:",omitempty"`
	WalletSize    uint    `json:",omitempty"`
}

func (*AccountSet) TxType() TxType {
	return AccountSetTx
}
