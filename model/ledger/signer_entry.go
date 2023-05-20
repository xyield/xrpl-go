package ledger

import (
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type SignerWrapper struct {
	SignerEntry SignerEntry
}

type SignerEntry struct {
	Account       types.Address
	SignerWeight  uint64
	WalletLocator types.Hash256 `json:",omitempty"`
}
