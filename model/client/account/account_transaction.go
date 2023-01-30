package account

import (
	"encoding/json"
	"fmt"

	. "github.com/xyield/xrpl-go/model/transactions"
)

const (
	AccountTxUnmarshalErr string = "Unmarshal JSON AccountTransaction"
)

type AccountTransaction struct {
	// WARNING: modifications to AccountTransaction
	// must be reflected in UnmarshalJSON
	LedgerIndex uint64              `json:"ledger_index"`
	Meta        TransactionMetadata `json:"meta"`
	Tx          Tx                  `json:"tx"`
	TxBlob      string              `json:"tx_blob"`
	Validated   bool                `json:"validated"`
}

func (at *AccountTransaction) UnmarshalJSON(data []byte) error {
	type accountTxDecodeHelper struct {
		LedgerIndex uint64              `json:"ledger_index"`
		Meta        TransactionMetadata `json:"meta"`
		Tx          json.RawMessage     `json:"tx"`
		TxBlob      string              `json:"tx_blob"`
		Validated   bool                `json:"validated"`
	}
	var values accountTxDecodeHelper
	if err := json.Unmarshal(data, &values); err != nil {
		return fmt.Errorf("%s: %w", AccountTxUnmarshalErr, err)
	}
	at.LedgerIndex = values.LedgerIndex
	at.Meta = values.Meta
	at.TxBlob = values.TxBlob
	at.Validated = values.Validated

	var err error
	at.Tx, err = UnmarshalTx(values.Tx)
	if err != nil {
		return fmt.Errorf("%s: %w", AccountTxUnmarshalErr, err)
	}

	return nil
}
