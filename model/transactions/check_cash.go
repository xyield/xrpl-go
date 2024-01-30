package transactions

import (
	"encoding/json"

	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type CheckCash struct {
	BaseTx
	CheckID    types.Hash256
	Amount     types.CurrencyAmount `json:",omitempty"`
	DeliverMin types.CurrencyAmount `json:",omitempty"`
}

func (*CheckCash) TxType() TxType {
	return CheckCashTx
}

func (tx *CheckCash) UnmarshalJSON(data []byte) error {
	type ccHelper struct {
		BaseTx
		CheckID    types.Hash256
		Amount     json.RawMessage `json:",omitempty"`
		DeliverMin json.RawMessage `json:",omitempty"`
	}
	var h ccHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*tx = CheckCash{
		BaseTx:  h.BaseTx,
		CheckID: h.CheckID,
	}

	var amount, min types.CurrencyAmount
	var err error
	amount, err = types.UnmarshalCurrencyAmount(h.Amount)
	if err != nil {
		return err
	}
	min, err = types.UnmarshalCurrencyAmount(h.DeliverMin)
	if err != nil {
		return err
	}
	tx.Amount = amount
	tx.DeliverMin = min
	return nil

}
