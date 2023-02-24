package account

import (
	"encoding/json"
	"fmt"

	. "github.com/xyield/xrpl-go/model/transactions"
)

const (
	ErrAccountTxUnmarshal string = "Unmarshal JSON AccountTransaction"
)

type AccountTransaction struct {
	LedgerIndex uint64 `json:"ledger_index"`
	Meta        TxMeta `json:"meta"`
	Tx          Tx     `json:"tx"`
	TxBlob      string `json:"tx_blob"`
	Validated   bool   `json:"validated"`
}

func (at *AccountTransaction) UnmarshalJSON(data []byte) error {
	type accountTxDecodeHelper struct {
		LedgerIndex uint64          `json:"ledger_index"`
		Meta        json.RawMessage `json:"meta"`
		Tx          json.RawMessage `json:"tx"`
		TxBlob      string          `json:"tx_blob"`
		Validated   bool            `json:"validated"`
	}
	var values accountTxDecodeHelper
	if err := json.Unmarshal(data, &values); err != nil {
		return fmt.Errorf("%s: %w", ErrAccountTxUnmarshal, err)
	}
	at.LedgerIndex = values.LedgerIndex
	at.TxBlob = values.TxBlob
	at.Validated = values.Validated

	var err error
	at.Tx, err = UnmarshalTx(values.Tx)
	if err != nil {
		return fmt.Errorf("%s: %w", ErrAccountTxUnmarshal, err)
	}
	at.Meta, err = UnmarshalTxMeta(values.Meta)
	if err != nil {
		return fmt.Errorf("%s: %w", ErrAccountTxUnmarshal, err)
	}

	return nil
}
