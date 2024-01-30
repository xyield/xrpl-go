package clio

import (
	"encoding/json"

	"github.com/CreatureDev/xrpl-go/model/client/common"
	"github.com/CreatureDev/xrpl-go/model/ledger"
	"github.com/CreatureDev/xrpl-go/model/transactions"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type LedgerResponse struct {
	Ledger      ClioLedger         `json:"ledger"`
	LedgerHash  common.LedgerHash  `json:"ledger_hash"`
	LedgerIndex common.LedgerIndex `json:"ledger_index"`
	Validated   bool               `json:"validated"`
}

type ClioLedger struct {
	AccountHash         string                  `json:"account_hash"`
	AccountState        []ledger.LedgerObject   `json:"accountState,omitempty"`
	CloseFlags          int                     `json:"close_flags"`
	CloseTime           uint                    `json:"close_time"`
	CloseTimeHuman      string                  `json:"close_time_human"`
	CloseTimeResolution int                     `json:"close_time_resolution"`
	Closed              bool                    `json:"closed"`
	LedgerHash          common.LedgerHash       `json:"ledger_hash"`
	LedgerIndex         string                  `json:"ledger_index"`
	ParentCloseTime     uint                    `json:"parent_close_time"`
	ParentHash          string                  `json:"parent_hash"`
	TotalCoins          types.XRPCurrencyAmount `json:"total_coins"`
	TransactionHash     string                  `json:"transaction_hash"`
	Transactions        []transactions.Tx       `json:"transactions,omitempty"`
}

func (l *ClioLedger) UnmarshalJSON(data []byte) error {
	type clHelper struct {
		AccountHash         string                  `json:"account_hash"`
		AccountState        []json.RawMessage       `json:"accountState,omitempty"`
		CloseFlags          int                     `json:"close_flags"`
		CloseTime           uint                    `json:"close_time"`
		CloseTimeHuman      string                  `json:"close_time_human"`
		CloseTimeResolution int                     `json:"close_time_resolution"`
		Closed              bool                    `json:"closed"`
		LedgerHash          common.LedgerHash       `json:"ledger_hash"`
		LedgerIndex         string                  `json:"ledger_index"`
		ParentCloseTime     uint                    `json:"parent_close_time"`
		ParentHash          string                  `json:"parent_hash"`
		TotalCoins          types.XRPCurrencyAmount `json:"total_coins"`
		TransactionHash     string                  `json:"transaction_hash"`
		Transactions        []json.RawMessage       `json:"transactions"`
	}
	var h clHelper
	err := json.Unmarshal(data, &h)
	if err != nil {
		return err
	}
	*l = ClioLedger{
		AccountHash:         h.AccountHash,
		CloseFlags:          h.CloseFlags,
		CloseTime:           h.CloseTime,
		CloseTimeHuman:      h.CloseTimeHuman,
		CloseTimeResolution: h.CloseTimeResolution,
		Closed:              h.Closed,
		LedgerHash:          h.LedgerHash,
		LedgerIndex:         h.LedgerIndex,
		ParentCloseTime:     h.ParentCloseTime,
		ParentHash:          h.ParentHash,
		TotalCoins:          h.TotalCoins,
		TransactionHash:     h.TransactionHash,
	}

	for _, state := range h.AccountState {
		obj, err := ledger.UnmarshalLedgerObject(state)
		if err != nil {
			return err
		}
		l.AccountState = append(l.AccountState, obj)
	}

	for _, tx := range h.Transactions {
		tx, err := transactions.UnmarshalTx(tx)
		if err != nil {
			return err
		}
		l.Transactions = append(l.Transactions, tx)
	}

	return nil
}
