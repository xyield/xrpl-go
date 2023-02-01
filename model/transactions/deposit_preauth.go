package transactions

type DepositPreauth struct {
	BaseTx
	Authorize   Address `json:",omitempty"`
	Unauthorize Address `json:",omitempty"`
}

func (*DepositPreauth) TxType() TxType {
	return DepositPreauthTx
}
