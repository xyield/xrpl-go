package transactions

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type SetRegularKey struct {
	BaseTx
	RegularKey Address `json:",omitempty"`
}

func (*SetRegularKey) TxType() TxType {
	return SetRegularKeyTx
}

func UnmarshalSetRegularKeyTx(data json.RawMessage) (Tx, error) {
	var ret SetRegularKey
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
