package account

import (
	"fmt"

	. "github.com/xyield/xrpl-go/model/transactions"
)

type AccountTransaction struct {
	LedgerIndex uint64              `json:"ledger_index"`
	Meta        TransactionMetadata `json:"meta"`
	Tx          Tx                  `json:"tx"`
	TxBlob      string              `json:"tx_blob"`
	Validated   bool                `json:"validated"`
}

func (at *AccountTransaction) UnmarshalJSON(data []byte) error {
	// TODO parse Tx interface
	return fmt.Errorf("Unimplemented")
}
