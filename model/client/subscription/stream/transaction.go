package stream

import (
	"encoding/json"

	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/transactions"
)

type TransactionStream struct {
	Type                StreamType             `json:"type"`
	EngineResult        string                 `json:"engine_result"`
	EngineResultCode    int                    `json:"engine_result_code"`
	EngineResultMessage string                 `json:"engine_result_message"`
	LedgerCurrentIndex  common.LedgerIndex     `json:"ledger_current_index,omitempty"`
	LedgerHash          common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex         common.LedgerIndex     `json:"ledger_index,omitempty"`
	Meta                transactions.TxObjMeta `json:"meta,omitempty"`
	Transaction         transactions.Tx        `json:"transaction"`
	Validated           bool                   `json:"validated"`
}

func (s *TransactionStream) UnmarshalJSON(data []byte) error {
	type tsHelper struct {
		Type                StreamType             `json:"type"`
		EngineResult        string                 `json:"engine_result"`
		EngineResultCode    int                    `json:"engine_result_code"`
		EngineResultMessage string                 `json:"engine_result_message"`
		LedgerCurrentIndex  common.LedgerIndex     `json:"ledger_current_index,omitempty"`
		LedgerHash          common.LedgerHash      `json:"ledger_hash,omitempty"`
		LedgerIndex         common.LedgerIndex     `json:"ledger_index,omitempty"`
		Meta                transactions.TxObjMeta `json:"meta,omitempty"`
		Transaction         json.RawMessage        `json:"transaction"`
		Validated           bool                   `json:"validated"`
	}
	var h tsHelper
	err := json.Unmarshal(data, &h)
	if err != nil {
		return err
	}

	*s = TransactionStream{
		Type:                h.Type,
		EngineResult:        h.EngineResult,
		EngineResultCode:    h.EngineResultCode,
		EngineResultMessage: h.EngineResultMessage,
		LedgerCurrentIndex:  h.LedgerCurrentIndex,
		LedgerHash:          h.LedgerHash,
		LedgerIndex:         h.LedgerIndex,
		Meta:                h.Meta,
		Validated:           h.Validated,
	}
	var tx transactions.Tx
	tx, err = transactions.UnmarshalTx(h.Transaction)
	if err != nil {
		return err
	}
	s.Transaction = tx
	return nil
}
