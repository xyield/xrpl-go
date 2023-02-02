package transactions

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type EscrowCreate struct {
	BaseTx
	Amount         XrpCurrencyAmount
	Destination    Address
	CancelAfter    uint   `json:",omitempty"`
	FinishAfter    uint   `json:",omitempty"`
	Condition      []byte `json:",omitempty"`
	DestinationTag uint   `json:",omitempty"`
}

func (*EscrowCreate) TxType() TxType {
	return EscrowCreateTx
}

func UnmarshalEscrowCreateTx(data json.RawMessage) (Tx, error) {
	var ret EscrowCreate
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
