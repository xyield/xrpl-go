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

func (n *NFTokenAcceptOffer) UnmarshalJSON(data []byte) error {
	type naoHelper struct {
		BaseTx
		NFTokenSellOffer Hash256         `json:",omitempty"`
		NFTokenBuyOffer  Hash256         `json:",omitempty"`
		NFTokenBrokerFee json.RawMessage `json:",omitempty"`
	}
	var h naoHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*n = NFTokenAcceptOffer{
		BaseTx:           h.BaseTx,
		NFTokenSellOffer: h.NFTokenSellOffer,
		NFTokenBuyOffer:  h.NFTokenBuyOffer,
	}

	fee, err := UnmarshalCurrencyAmount(h.NFTokenBrokerFee)
	if err != nil {
		return err
	}
	n.NFTokenBrokerFee = fee
	return nil
}
