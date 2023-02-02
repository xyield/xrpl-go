package transactions

import "encoding/json"

type OfferCancel struct {
	BaseTx
	OfferSequence uint
}

func (*OfferCancel) TxType() TxType {
	return OfferCancelTx
}

func UnmarshalOfferCancelTx(data json.RawMessage) (Tx, error) {
	var ret OfferCancel
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
