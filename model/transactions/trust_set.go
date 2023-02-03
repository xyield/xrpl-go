package transactions

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type TrustSet struct {
	BaseTx
	LimitAmount CurrencyAmount
	QualityIn   uint `json:",omitempty"`
	QualityOut  uint `json:",omitempty"`
}

func (*TrustSet) TxType() TxType {
	return TrustSetTx
}

func UnmarshalTrustSetTx(data json.RawMessage) (Tx, error) {
	var ret TrustSet
	type tsHelper struct {
		BaseTx
		LimitAmount json.RawMessage
		QualityIn   uint `json:",omitempty"`
		QualityOut  uint `json:",omitempty"`
	}
	var h tsHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return nil, err
	}
	ret = TrustSet{
		BaseTx:     h.BaseTx,
		QualityIn:  h.QualityIn,
		QualityOut: h.QualityOut,
	}
	limit, err := UnmarshalCurrencyAmount(h.LimitAmount)
	if err != nil {
		return nil, err
	}
	ret.LimitAmount = limit

	return &ret, nil
}
