package account

import (
	. "github.com/xyield/xrpl-go/model/transactions"
)

const (
	Burnable     NFTokenFlag = 0x0001
	OnlyXRP                  = 0x0002
	Transferable             = 0x0008
	ReservedFlag             = 0x8000
)

type NFTokenFlag uint

type Nft struct {
	Flags        NFTokenFlag
	Issuer       Address
	NFTokenID    string
	NFTokenTaxon uint
	URI          string
	NftSerial    uint `json:"nft_serial"`
}
