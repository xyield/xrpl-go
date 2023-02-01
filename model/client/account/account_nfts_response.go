package account

import (
	. "github.com/xyield/xrpl-go/model/client/common"
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountNftsResponse struct {
	Account            Address     `json:"account"`
	AccountNfts        []Nft       `json:"account_nfts"`
	LedgerIndex        LedgerIndex `json:"ledger_index,omitempty"`
	LedgerHash         LedgerHash  `json:"ledger_hash,omitempty"`
	LedgerCurrentIndex LedgerIndex `json:"ledger_current_index,omitempty"`
	Validated          bool        `json:"validated"`
}
