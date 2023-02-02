package transactions

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type NFTokenAcceptOffer struct {
	BaseTx
	NFTokenSellOffer Hash256        `json:",omitempty"`
	NFTokenBuyOffer  Hash256        `json:",omitempty"`
	NFTokenBrokerFee CurrencyAmount `json:",omitempty"`
}

func (*NFTokenAcceptOffer) TxType() TxType {
	return NFTokenAcceptOfferTx
}

func UnmarshalNFTokenAcceptOfferTx(data json.RawMessage) (Tx, error) {
	var ret NFTokenAcceptOffer
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
