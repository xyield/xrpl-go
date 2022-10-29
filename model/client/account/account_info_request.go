package client

import (
	. "github.com/xyield/xrpl-go/model/client/common"
	. "github.com/xyield/xrpl-go/model/transactions"
)

type AccountInfoRequest struct {
	AccountID   Address     `json:"account"`
	LedgerIndex LedgerIndex `json:"ledger_index,omitempty"`
	LedgerHash  LedgerHash  `json:"ledger_hash,omitempty"`
	Queue       bool        `json:"queue,omitempty"`
	SignerList  bool        `json:"signer_list,omitempty"`
	Strict      bool        `json:"strict,omitempty"`
}
