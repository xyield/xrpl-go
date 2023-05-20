package account

import (
	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountNFTsResponse struct {
	Account            types.Address      `json:"account"`
	AccountNFTs        []NFT              `json:"account_nfts"`
	LedgerIndex        common.LedgerIndex `json:"ledger_index,omitempty"`
	LedgerHash         common.LedgerHash  `json:"ledger_hash,omitempty"`
	LedgerCurrentIndex common.LedgerIndex `json:"ledger_current_index,omitempty"`
	Validated          bool               `json:"validated"`
}
