package transactions

import (
	"encoding/json"

	"github.com/CreatureDev/xrpl-go/model/client/common"
	"github.com/CreatureDev/xrpl-go/model/transactions"
)

type SubmitResponse struct {
	EngineResult             string             `json:"engine_result"`
	EngineResultCode         int                `json:"engine_result_code"`
	EngineResultMessage      string             `json:"engine_result_message"`
	TxBlob                   string             `json:"tx_blob"`
	Tx                       transactions.Tx    `json:"tx_json"`
	Accepted                 bool               `json:"accepted"`
	AccountSequenceAvailable uint32             `json:"account_sequence_available"`
	AccountSequenceNext      uint32             `json:"account_sequence_next"`
	Applied                  bool               `json:"applied"`
	Broadcast                bool               `json:"broadcast"`
	Kept                     bool               `json:"kept"`
	Queued                   bool               `json:"queued"`
	OpenLedgerCost           string             `json:"open_ledger_cost"`
	ValidatedLedgerIndex     common.LedgerIndex `json:"validated_ledger_index"`
}

func (r *SubmitResponse) UnmarshalJSON(data []byte) error {
	type sHelper struct {
		EngineResult             string             `json:"engine_result"`
		EngineResultCode         int                `json:"engine_result_code"`
		EngineResultMessage      string             `json:"engine_result_message"`
		TxBlob                   string             `json:"tx_blob"`
		Tx                       json.RawMessage    `json:"tx_json"`
		Accepted                 bool               `json:"accepted"`
		AccountSequenceAvailable uint32             `json:"account_sequence_available"`
		AccountSequenceNext      uint32             `json:"account_sequence_next"`
		Applied                  bool               `json:"applied"`
		Broadcast                bool               `json:"broadcast"`
		Kept                     bool               `json:"kept"`
		Queued                   bool               `json:"queued"`
		OpenLedgerCost           string             `json:"open_ledger_cost"`
		ValidatedLedgerIndex     common.LedgerIndex `json:"validated_ledger_index"`
	}
	var h sHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = SubmitResponse{
		EngineResult:             h.EngineResult,
		EngineResultCode:         h.EngineResultCode,
		EngineResultMessage:      h.EngineResultMessage,
		TxBlob:                   h.TxBlob,
		Accepted:                 h.Accepted,
		AccountSequenceAvailable: h.AccountSequenceAvailable,
		AccountSequenceNext:      h.AccountSequenceNext,
		Applied:                  h.Applied,
		Broadcast:                h.Broadcast,
		Kept:                     h.Kept,
		Queued:                   h.Queued,
		OpenLedgerCost:           h.OpenLedgerCost,
		ValidatedLedgerIndex:     h.ValidatedLedgerIndex,
	}
	tx, err := transactions.UnmarshalTx(h.Tx)
	if err != nil {
		return err
	}
	r.Tx = tx

	return nil
}
