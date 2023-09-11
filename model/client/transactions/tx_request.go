package transactions

import "github.com/xyield/xrpl-go/model/client/common"

type TxRequest struct {
	Transaction string             `json:"transaction"`
	Binary      bool               `json:"binary,omitempty"`
	MinLedger   common.LedgerIndex `json:"min_ledger,omitempty"`
	MaxLedger   common.LedgerIndex `json:"max_ledger,omitempty"`
}

func (*TxRequest) Method() string {
	return "tx"
}

func (*TxRequest) Validate() error {
	return nil
}
