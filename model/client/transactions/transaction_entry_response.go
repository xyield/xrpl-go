package transactions

import (
	"encoding/json"

	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/transactions"
)

type TransactionEntryResponse struct {
	LedgerIndex common.LedgerIndex     `json:"ledger_index"`
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	Metadata    transactions.TxObjMeta `json:"metadata"`
	Tx          transactions.Tx        `json:"tx_json"`
}

func (r *TransactionEntryResponse) UnmarshalJSON(data []byte) error {
	type terHelper struct {
		LedgerIndex common.LedgerIndex     `json:"ledger_index"`
		LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
		Metadata    transactions.TxObjMeta `json:"metadata"`
		Tx          json.RawMessage        `json:"tx_json"`
	}
	var h terHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = TransactionEntryResponse{
		LedgerIndex: h.LedgerIndex,
		LedgerHash:  h.LedgerHash,
		Metadata:    h.Metadata,
	}
	tx, err := transactions.UnmarshalTx(h.Tx)
	if err != nil {
		return err
	}
	r.Tx = tx

	return nil
}
