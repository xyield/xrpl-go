package transactions

import (
	"github.com/xyield/xrpl-go/model/ledger"
)

type SignerListSet struct {
	BaseTx
	SignerQuorum  uint
	SignerEntries []ledger.SignerWrapper
}

func (*SignerListSet) TxType() TxType {
	return SignerListSetTx
}
