package account

import (
	"github.com/xyield/xrpl-go/model/transactions/types"
)

const (
	Burnable     NFTokenFlag = 0x0001
	OnlyXRP      NFTokenFlag = 0x0002
	Transferable NFTokenFlag = 0x0008
	ReservedFlag NFTokenFlag = 0x8000
)

type NFTokenFlag uint

type NFT struct {
	Flags        NFTokenFlag
	Issuer       types.Address
	NFTokenID    types.NFTokenID
	NFTokenTaxon uint
	URI          types.NFTokenURI
	NFTSerial    uint `json:"nft_serial"`
}
