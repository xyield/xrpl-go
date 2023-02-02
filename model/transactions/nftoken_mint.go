package transactions

import (
	"encoding/json"

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

func UnmarshalNFTokenMintTx(data json.RawMessage) (Tx, error) {
	var ret NFTokenMint
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
