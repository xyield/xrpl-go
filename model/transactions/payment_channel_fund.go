package transactions

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type PaymentChannelFund struct {
	BaseTx
	Channel    Hash256
	Amount     XrpCurrencyAmount
	Expiration uint `json:",omitempty"`
}

func (*PaymentChannelFund) TxType() TxType {
	return PaymentChannelFundTx
}

func UnmarshalPaymentChannelFundTx(data json.RawMessage) (Tx, error) {
	var ret PaymentChannelFund
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
