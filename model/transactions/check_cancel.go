package transactions

type CheckCancel struct {
	BaseTx
	CheckID Hash256
}

func (*CheckCancel) TxType() TxType {
	return CheckCancelTx
}
