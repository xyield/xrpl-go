package transactions

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type CheckCancel struct {
	BaseTx
	CheckID Hash256
}

func (*CheckCancel) TxType() TxType {
	return CheckCancelTx
}

func UnmarshalCheckCancelTx(data json.RawMessage) (Tx, error) {
	var ret CheckCancel
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
