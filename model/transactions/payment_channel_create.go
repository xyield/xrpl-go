package transactions

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type PaymentChannelCreate struct {
	BaseTx
	Amount         XrpCurrencyAmount
	Destination    Address
	SettleDelay    uint
	PublicKey      []byte
	CancelAfter    uint `json:",omitempty"`
	DestinationTag uint `json:",omitempty"`
}

func (*PaymentChannelCreate) TxType() TxType {
	return PaymentChannelCreateTx
}

func UnmarshalPaymentChannelCreateTx(data json.RawMessage) (Tx, error) {
	var ret PaymentChannelCreate
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
