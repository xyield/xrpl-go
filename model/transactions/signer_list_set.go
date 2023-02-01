package transactions

import (
	// TODO restructure to fix import cycle
	"github.com/xyield/xrpl-go/model/ledger"
)

type SignerListSet struct {
	BaseTx
	SignerQuorum  uint
	SignerEntries []ledger.SignerEntry
}

func (*SignerListSet) TxType() TxType {
	return SignerListSetTx
}
