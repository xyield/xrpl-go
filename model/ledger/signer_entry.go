package ledger

import (
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type SignerEntry struct {
	Account       Address
	SignerWeight  uint64
	WalletLocator Hash256 `json:",omitempty"`
}
