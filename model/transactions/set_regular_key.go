package transactions

type SetRegularKey struct {
	BaseTx
	RegularKey Address `json:",omitempty"`
}

func (*SetRegularKey) TxType() TxType {
	return SetRegularKeyTx
}
