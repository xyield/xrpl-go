package transactions

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type DepositPreauth struct {
	BaseTx
	Authorize   Address `json:",omitempty"`
	Unauthorize Address `json:",omitempty"`
}

func (*DepositPreauth) TxType() TxType {
	return DepositPreauthTx
}

func UnmarshalDepositPreauthTx(data json.RawMessage) (Tx, error) {
	var ret DepositPreauth
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
