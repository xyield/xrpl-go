package transactions

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type EscrowFinish struct {
	BaseTx
	Owner         Address
	OfferSequence uint
	Condition     []byte `json:",omitempty"`
	Fulfillment   []byte `json:",omitempty"`
}

func (*EscrowFinish) TxType() TxType {
	return EscrowFinishTx
}

func UnmarshalEscrowFinishTx(data json.RawMessage) (Tx, error) {
	var ret EscrowFinish
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
