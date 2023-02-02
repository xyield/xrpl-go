package transactions

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type EscrowCancel struct {
	BaseTx
	Owner         Address
	OfferSequence uint
}

func (*EscrowCancel) TxType() TxType {
	return EscrowCancelTx
}

func UnmarshalEscrowCancelTx(data json.RawMessage) (Tx, error) {
	var ret EscrowCancel
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
