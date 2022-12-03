package transactions

type Tx interface {
	TxType() TxType
}
