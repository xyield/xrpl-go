package clio

import (
	"github.com/CreatureDev/xrpl-go/model/client/account"
	"github.com/CreatureDev/xrpl-go/model/client/common"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type NFTHistoryResponse struct {
	NFTokenID      types.NFTokenID              `json:"nft_id"`
	LedgerIndexMin common.LedgerIndex           `json:"ledger_index_min"`
	LedgerIndexMax common.LedgerIndex           `json:"ledger_index_max"`
	Limit          uint                         `json:"limit"`
	Marker         any                          `json:"marker"`
	Transactions   []account.AccountTransaction `json:"transactions"`
	Validated      bool                         `json:"validated"`
}
