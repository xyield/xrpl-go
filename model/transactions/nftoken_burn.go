package transactions

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type NFTokenBurn struct {
	BaseTx
	NFTokenID NFTokenID
	Owner     Address `json:",omitempty"`
}

func (*NFTokenBurn) TxType() TxType {
	return NFTokenBurnTx
}

func UnmarshalNFTokenBurnTx(data json.RawMessage) (Tx, error) {
	var ret NFTokenBurn
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
