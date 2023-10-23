package clio

import (
	"github.com/CreatureDev/xrpl-go/model/client/common"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type NFTInfoResponse struct {
	NFTokenID       types.NFTokenID    `json:"nft_id"`
	LedgerIndex     common.LedgerIndex `json:"ledger_index"`
	Owner           types.Address      `json:"owner"`
	IsBurned        bool               `json:"is_burned"`
	Flags           uint               `json:"flags"`
	TransferFee     uint               `json:"transfer_fee"`
	Issuer          types.Address      `json:"issuer"`
	NFTokenTaxon    uint               `json:"nft_taxon"`
	NFTokenSequence uint32             `json:"nft_sequence"`
	URI             types.NFTokenURI   `json:"uri,omitempty"`
}
