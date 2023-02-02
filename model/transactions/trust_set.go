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
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
