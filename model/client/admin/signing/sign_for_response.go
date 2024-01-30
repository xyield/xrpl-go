package signing

import (
	"encoding/json"

	"github.com/CreatureDev/xrpl-go/model/transactions"
)

type SignForResponse struct {
	TxBlob string          `json:"tx_blob"`
	TxJson transactions.Tx `json:"tx_json"`
}

func (r *SignForResponse) UnmarshalJSON(data []byte) error {
	type srHelper struct {
		TxBlob string          `json:"tx_blob"`
		TxJson json.RawMessage `json:"tx_json"`
	}
	var h srHelper
	err := json.Unmarshal(data, &h)
	if err != nil {
		return nil
	}
	r.TxBlob = h.TxBlob
	r.TxJson, err = transactions.UnmarshalTx(h.TxJson)
	return err
}
