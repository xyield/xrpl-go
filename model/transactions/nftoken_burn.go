package transactions

type NFTokenBurn struct {
	BaseTx
	NFTokenID NFTokenID
	Owner     Address `json:",omitempty"`
}

func (*NFTokenBurn) TxType() TxType {
	return NFTokenBurnTx
}
