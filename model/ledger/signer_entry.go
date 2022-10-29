package ledger

import (
	. "github.com/xyield/xrpl-go/model/transactions"
)

type SignerEntry struct {
	Account       Address `json:"Account"`
	SignerWeight  uint64  `json:"SignerWeight"`
	WalletLocator Hash256 `json:"WalletLocator"`
}
