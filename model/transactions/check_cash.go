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
	var ret CheckCash
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
