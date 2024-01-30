package transactions

import (
	"fmt"

	"github.com/xyield/xrpl-go/model/client/common"
)

type TxRequest struct {
	Transaction string             `json:"transaction"`
	Binary      bool               `json:"binary,omitempty"`
	MinLedger   common.LedgerIndex `json:"min_ledger,omitempty"`
	MaxLedger   common.LedgerIndex `json:"max_ledger,omitempty"`
}

func (*TxRequest) Method() string {
	return "tx"
}

func (t *TxRequest) Validate() error {
	if t.Transaction == "" {
		return fmt.Errorf("transaction request: missing transaction")
	}
	return nil
}
