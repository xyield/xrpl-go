package transactions

import (
	"github.com/CreatureDev/xrpl-go/model/ledger"
)

type SignerListSet struct {
	BaseTx
	SignerQuorum  uint
	SignerEntries []ledger.SignerEntryWrapper
}

func (*SignerListSet) TxType() TxType {
	return SignerListSetTx
}
