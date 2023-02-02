package transactions

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type NFTokenCancelOffer struct {
	BaseTx
	NFTokenOffer []Hash256
}

func (*NFTokenCancelOffer) TxType() TxType {
	return NFTokenCancelOfferTx
}

func UnmarshalNFTokenCancelOfferTx(data json.RawMessage) (Tx, error) {
	var ret NFTokenCancelOffer
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
