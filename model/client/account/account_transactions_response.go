package account

import (
	. "github.com/xyield/xrpl-go/model/client/common"
	. "github.com/xyield/xrpl-go/model/transactions"
)

type AccountTransactionsResponse struct {
	Account        Address              `json:"account"`
	LedgerIndexMin LedgerIndex          `json:"ledger_index_min"`
	LedgerIndexMax LedgerIndex          `json:"ledger_index_max"`
	Limit          int                  `json:"limit"`
	Marker         interface{}          `json:"marker"`
	Transactions   []AccountTransaction `json:"transactions"`
	Validated      bool                 `json:"validated"`
}

type AccountTransaction struct {
	LedgerIndex uint64              `json:"ledger_index"`
	Meta        TransactionMetadata `json:"meta"`
	// TODO parsing of interfaces via json.RawMessage intermediary
	//TxJson      json.RawMessage `json:"tx"`
	Tx        Tx     `json:"tx"`
	TxBlob    string `json:"tx_blob"`
	Validated bool   `json:"validated"`
}
