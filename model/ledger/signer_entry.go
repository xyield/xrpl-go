package ledger

import (
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type SignerEntry struct {
	Account       types.Address
	SignerWeight  uint64
	WalletLocator types.Hash256 `json:",omitempty"`
}
