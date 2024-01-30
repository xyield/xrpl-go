package transactions

import (
	"encoding/json"

	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type NFTokenAcceptOffer struct {
	BaseTx
	NFTokenSellOffer types.Hash256        `json:",omitempty"`
	NFTokenBuyOffer  types.Hash256        `json:",omitempty"`
	NFTokenBrokerFee types.CurrencyAmount `json:",omitempty"`
}

func (*NFTokenAcceptOffer) TxType() TxType {
	return NFTokenAcceptOfferTx
}

func (n *NFTokenAcceptOffer) UnmarshalJSON(data []byte) error {
	type naoHelper struct {
		BaseTx
		NFTokenSellOffer types.Hash256   `json:",omitempty"`
		NFTokenBuyOffer  types.Hash256   `json:",omitempty"`
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

	fee, err := types.UnmarshalCurrencyAmount(h.NFTokenBrokerFee)
	if err != nil {
		return err
	}
	n.NFTokenBrokerFee = fee
	return nil
}
