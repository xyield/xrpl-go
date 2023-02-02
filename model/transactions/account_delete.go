package transactions

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountDelete struct {
	BaseTx
	Destination    Address
	DestinationTag uint `json:",omitempty"`
}

func (*AccountDelete) TxType() TxType {
	return AccountDeleteTx
}

func UnmarshalAccountDeleteTx(data json.RawMessage) (Tx, error) {
	var ret AccountDelete
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
