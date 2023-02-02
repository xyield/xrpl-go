package transactions

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type NFTokenCreateOffer struct {
	BaseTx
	Owner       Address `json:",omitempty"`
	NFTokenID   NFTokenID
	Amount      CurrencyAmount
	Expiration  uint    `json:",omitempty"`
	Destination Address `json:",omitempty"`
}

func (*NFTokenCreateOffer) TxType() TxType {
	return NFTokenCreateOfferTx
}

func UnmarshalNFTokenCreateOfferTx(data json.RawMessage) (Tx, error) {
	var ret NFTokenCreateOffer
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
