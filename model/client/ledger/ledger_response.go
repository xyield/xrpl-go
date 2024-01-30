package ledger

import (
	"encoding/json"

	"github.com/CreatureDev/xrpl-go/model/client/common"
	"github.com/CreatureDev/xrpl-go/model/ledger"
	"github.com/CreatureDev/xrpl-go/model/transactions"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type LedgerResponse struct {
	Ledger      LedgerHeader       `json:"ledger"`
	LedgerHash  string             `json:"ledger_hash"`
	LedgerIndex common.LedgerIndex `json:"ledger_index"`
	Validated   bool               `json:"validated,omitempty"`
	QueueData   []LedgerQueueData  `json:"queue_data,omitempty"`
}

type LedgerHeader struct {
	AccountHash         string                  `json:"account_hash"`
	AccountState        []ledger.LedgerObject   `json:"accountState,omitempty"`
	CloseFlags          int                     `json:"close_flags"`
	CloseTime           int                     `json:"close_time"`
	CloseTimeHuman      string                  `json:"close_time_human"`
	CloseTimeResolution int                     `json:"close_time_resolution"`
	Closed              bool                    `json:"closed"`
	LedgerHash          string                  `json:"ledger_hash"`
	LedgerIndex         string                  `json:"ledger_index"`
	ParentCloseTime     int                     `json:"parent_close_time"`
	ParentHash          string                  `json:"parent_hash"`
	TotalCoins          types.XRPCurrencyAmount `json:"total_coins"`
	TransactionHash     string                  `json:"transaction_hash"`
	Transactions        []transactions.Tx       `json:"transactions,omitempty"`
}

func (r *LedgerHeader) UnmarshalJSON(data []byte) error {
	type lhHelper struct {
		AccountHash         string                  `json:"account_hash"`
		AccountState        []json.RawMessage       `json:"accountState,omitempty"`
		CloseFlags          int                     `json:"close_flags"`
		CloseTime           int                     `json:"close_time"`
		CloseTimeHuman      string                  `json:"close_time_human"`
		CloseTimeResolution int                     `json:"close_time_resolution"`
		Closed              bool                    `json:"closed"`
		LedgerHash          string                  `json:"ledger_hash"`
		LedgerIndex         string                  `json:"ledger_index"`
		ParentCloseTime     int                     `json:"parent_close_time"`
		ParentHash          string                  `json:"parent_hash"`
		TotalCoins          types.XRPCurrencyAmount `json:"total_coins"`
		TransactionHash     string                  `json:"transaction_hash"`
		Transactions        []json.RawMessage       `json:"transactions,omitempty"`
	}
	var h lhHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = LedgerHeader{
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

	if h.AccountState != nil && len(h.AccountState) > 0 {
		var obj []ledger.LedgerObject
		for _, s := range h.AccountState {
			o, e := ledger.UnmarshalLedgerObject(s)
			if e != nil {
				return e
			}
			obj = append(obj, o)
		}
		r.AccountState = obj
	}

	if h.Transactions != nil && len(h.Transactions) > 0 {
		var tx []transactions.Tx
		for _, s := range h.Transactions {
			t, e := transactions.UnmarshalTx(s)
			if e != nil {
				return e
			}
			tx = append(tx, t)
		}
		r.Transactions = tx
	}

	return nil
}

type LedgerQueueData struct {
	Account          types.Address           `json:"account"`
	Tx               transactions.Tx         `json:"tx"`
	RetriesRemaining int                     `json:"retries_remaining"`
	PreflightResult  string                  `json:"preflight_result"`
	LastResult       string                  `json:"last_result,omitempty"`
	AuthChange       bool                    `json:"auth_change,omitempty"`
	Fee              types.XRPCurrencyAmount `json:"fee,omitempty"`
	FeeLevel         types.XRPCurrencyAmount `json:"fee_level,omitempty"`
	MaxSpendDrops    types.XRPCurrencyAmount `json:"max_spend_drops,omitempty"`
}

func (r *LedgerQueueData) UnmarshalJSON(data []byte) error {
	type lqdHelper struct {
		Account          types.Address           `json:"account"`
		Tx               json.RawMessage         `json:"tx"`
		RetriesRemaining int                     `json:"retries_remaining"`
		PreflightResult  string                  `json:"preflight_result"`
		LastResult       string                  `json:"last_result,omitempty"`
		AuthChange       bool                    `json:"auth_change,omitempty"`
		Fee              types.XRPCurrencyAmount `json:"fee,omitempty"`
		FeeLevel         types.XRPCurrencyAmount `json:"fee_level,omitempty"`
		MaxSpendDrops    types.XRPCurrencyAmount `json:"max_spend_drops,omitempty"`
	}
	var h lqdHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = LedgerQueueData{
		Account:          h.Account,
		RetriesRemaining: h.RetriesRemaining,
		PreflightResult:  h.PreflightResult,
		LastResult:       h.LastResult,
		AuthChange:       h.AuthChange,
		Fee:              h.Fee,
		FeeLevel:         h.FeeLevel,
		MaxSpendDrops:    h.MaxSpendDrops,
	}
	tx, err := transactions.UnmarshalTx(h.Tx)
	if err != nil {
		return err
	}
	r.Tx = tx

	return nil
}
