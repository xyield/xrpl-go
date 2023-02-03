package transactions

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type CheckCash struct {
	BaseTx
	CheckID    Hash256
	Amount     CurrencyAmount `json:",omitempty"`
	DeliverMin CurrencyAmount `json:",omitempty"`
}

func (*CheckCash) TxType() TxType {
	return CheckCashTx
}

func UnmarshalCheckCashTx(data json.RawMessage) (Tx, error) {
	type ccHelper struct {
		BaseTx
		CheckID    Hash256
		Amount     json.RawMessage `json:",omitempty"`
		DeliverMin json.RawMessage `json:",omitempty"`
	}
	var ret CheckCash
	var h ccHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return nil, err
	}
	ret = CheckCash{
		BaseTx:  h.BaseTx,
		CheckID: h.CheckID,
	}

	var amount, min CurrencyAmount
	var err error
	amount, err = UnmarshalCurrencyAmount(h.Amount)
	if err != nil {
		return nil, err
	}
	min, err = UnmarshalCurrencyAmount(h.DeliverMin)
	if err != nil {
		return nil, err
	}
	ret.Amount = amount
	ret.DeliverMin = min

	return &ret, nil
}
