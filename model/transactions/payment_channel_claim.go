package transactions

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type PaymentChannelClaim struct {
	BaseTx
	Channel   Hash256
	Balance   XrpCurrencyAmount `json:",omitempty"`
	Amount    XrpCurrencyAmount `json:",omitempty"`
	Signature []byte            `json:",omitempty"`
	PublicKey []byte            `json:",omitempty"`
}

func (*PaymentChannelClaim) TxType() TxType {
	return PaymentChannelClaimTx
}

func UnmarshalPaymentChannelClaimTx(data json.RawMessage) (Tx, error) {
	var ret PaymentChannelClaim
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
