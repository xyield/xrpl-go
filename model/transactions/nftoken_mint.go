package transactions

import (
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type NFTokenMint struct {
	BaseTx
	NFTokenTaxon uint
	Issuer       Address    `json:",omitempty"`
	TransferFee  uint16     `json:",omitempty"`
	URI          NFTokenURI `json:",omitempty"`
}

func (*NFTokenMint) TxType() TxType {
	return NFTokenMintTx
}
