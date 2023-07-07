package transactions

import (
	"encoding/json"

	"github.com/xyield/xrpl-go/model/transactions"
)

type SubmitMultisignedResponse struct {
	EngineResult        string          `json:"engine_result"`
	EngineResultCode    int             `json:"engine_result_code"`
	EngineResultMessage string          `json:"engine_result_message"`
	TxBlob              string          `json:"tx_blob"`
	Tx                  transactions.Tx `json:"tx_json"`
}

func (r *SubmitMultisignedResponse) UnmarshalJSON(data []byte) error {
	type sHelper struct {
		EngineResult        string          `json:"engine_result"`
		EngineResultCode    int             `json:"engine_result_code"`
		EngineResultMessage string          `json:"engine_result_message"`
		TxBlob              string          `json:"tx_blob"`
		Tx                  json.RawMessage `json:"tx_json"`
	}
	var h sHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = SubmitMultisignedResponse{
		EngineResult:        h.EngineResult,
		EngineResultCode:    h.EngineResultCode,
		EngineResultMessage: h.EngineResultMessage,
		TxBlob:              h.TxBlob,
	}
	tx, err := transactions.UnmarshalTx(h.Tx)
	if err != nil {
		return err
	}
	r.Tx = tx

	return nil
}
