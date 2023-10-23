package transactions

import (
	"encoding/json"
	"fmt"

	"github.com/CreatureDev/xrpl-go/model/ledger"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type TxMeta interface {
	TxMeta()
}

func UnmarshalTxMeta(data []byte) (TxMeta, error) {
	if len(data) == 0 {
		return nil, nil
	}
	switch data[0] {
	case '"':
		var t TxBinMeta
		err := json.Unmarshal(data, &t)
		return t, err
	case '{':
		var o TxObjMeta
		err := json.Unmarshal(data, &o)
		return o, err
	default:
		return nil, fmt.Errorf("unrecognized TxMeta format")
	}
}

type TxBinMeta string

func (TxBinMeta) TxMeta() {}

type TxObjMeta struct {
	AffectedNodes          []AffectedNode       `json:"AffectedNodes,omitempty"`
	PartialDeliveredAmount types.CurrencyAmount `json:"DeliveredAmount,omitempty"`
	TransactionIndex       uint64               `json:"TransactionIndex,omitempty"`
	TransactionResult      string               `json:"TransactionResult,omitempty"`
	DeliveredAmount        types.CurrencyAmount `json:"delivered_amount,omitempty"`
}

func (m *TxObjMeta) UnmarshalJSON(data []byte) error {
	var h struct {
		AffectedNodes          []AffectedNode  `json:"AffectedNodes,omitempty"`
		PartialDeliveredAmount json.RawMessage `json:"DeliveredAmount,omitempty"`
		TransactionIndex       uint64          `json:"TransactionIndex,omitempty"`
		TransactionResult      string          `json:"TransactionResult,omitempty"`
		DeliveredAmount        json.RawMessage `json:"delivered_amount,omitempty"`
	}
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*m = TxObjMeta{
		AffectedNodes:     h.AffectedNodes,
		TransactionIndex:  h.TransactionIndex,
		TransactionResult: h.TransactionResult,
	}
	if h.PartialDeliveredAmount != nil && len(h.PartialDeliveredAmount) > 0 {
		if part, err := types.UnmarshalCurrencyAmount(h.PartialDeliveredAmount); err != nil {
			return fmt.Errorf("unmarshal TxMeta object: %w", err)
		} else {
			m.PartialDeliveredAmount = part
		}
	}
	if h.DeliveredAmount != nil && len(h.DeliveredAmount) > 0 {
		if deliv, err := types.UnmarshalCurrencyAmount(h.DeliveredAmount); err != nil {
			return fmt.Errorf("unmarshal TxMeta object: %w", err)
		} else {
			m.DeliveredAmount = deliv
		}
	}
	return nil
}

func (TxObjMeta) TxMeta() {}

type AffectedNode struct {
	CreatedNode  *CreatedNode  `json:"CreatedNode,omitempty"`
	ModifiedNode *ModifiedNode `json:"ModifiedNode,omitempty"`
	DeletedNode  *DeletedNode  `json:"DeletedNode,omitempty"`
}

type CreatedNode struct {
	LedgerEntryType ledger.LedgerEntryType `json:"LedgerEntryType,omitempty"`
	LedgerIndex     string                 `json:"LedgerIndex,omitempty"`
	NewFields       ledger.LedgerObject    `json:"NewFields,omitempty"`
}

func (n *CreatedNode) UnmarshalJSON(data []byte) error {
	var h struct {
		LedgerEntryType ledger.LedgerEntryType
		LedgerIndex     string
		NewFields       json.RawMessage
	}
	err := json.Unmarshal(data, &h)
	if err != nil {
		return fmt.Errorf("unmarshal CreatedNode: %w", err)
	}
	var obj ledger.LedgerObject
	if obj, err = ledger.EmptyLedgerObject(string(h.LedgerEntryType)); err != nil {
		return fmt.Errorf("unmarshal CreatedNode: %w", err)
	}
	et := obj.EntryType()
	if h.NewFields != nil && len(h.NewFields) > 0 {
		if err = json.Unmarshal(h.NewFields, obj); err != nil {
			return fmt.Errorf("unmarshal CreatedNode: %w", err)
		}
	} else {
		obj = nil
	}
	*n = CreatedNode{
		LedgerEntryType: et,
		LedgerIndex:     h.LedgerIndex,
		NewFields:       obj,
	}
	return nil
}

type ModifiedNode struct {
	LedgerEntryType   ledger.LedgerEntryType `json:"LedgerEntryType,omitempty"`
	LedgerIndex       string                 `json:"LedgerIndex,omitempty"`
	FinalFields       ledger.LedgerObject    `json:"FinalFields,omitempty"`
	PreviousFields    ledger.LedgerObject    `json:"PreviousFields,omitempty"`
	PreviousTxnID     string                 `json:"PreviousTxnID,omitempty"`
	PreviousTxnLgrSeq uint32                 `json:"PreviousTxnLgrSeq,omitempty"`
}

func (n *ModifiedNode) UnmarshalJSON(data []byte) error {
	var h struct {
		LedgerEntryType   ledger.LedgerEntryType
		LedgerIndex       string
		FinalFields       json.RawMessage
		PreviousFields    json.RawMessage
		PreviousTxnID     string
		PreviousTxnLgrSeq uint32
	}
	err := json.Unmarshal(data, &h)
	if err != nil {
		return fmt.Errorf("unmarshal ModifiedNode: %w", err)
	}
	var prev, fin ledger.LedgerObject

	if prev, err = ledger.EmptyLedgerObject(string(h.LedgerEntryType)); err != nil {
		return fmt.Errorf("unmarshal ModifiedNode: %w", err)
	}
	// Get entry type before possible nullification
	et := prev.EntryType()
	if fin, err = ledger.EmptyLedgerObject(string(h.LedgerEntryType)); err != nil {
		return fmt.Errorf("unmarshal ModifiedNode: %w", err)
	}
	if h.PreviousFields != nil && len(h.PreviousFields) > 0 {
		if err = json.Unmarshal(h.PreviousFields, prev); err != nil {
			return fmt.Errorf("unmarshal ModifiedNode: %w", err)
		}
	} else {
		// Do not set if no previous fields
		prev = nil
	}
	if h.FinalFields != nil && len(h.FinalFields) > 0 {
		if err = json.Unmarshal(h.FinalFields, fin); err != nil {
			return fmt.Errorf("unmarshal ModifiedNode: %w", err)
		}
	} else {
		// Do not set if no final fields
		fin = nil
	}
	*n = ModifiedNode{
		LedgerEntryType:   et,
		LedgerIndex:       h.LedgerIndex,
		FinalFields:       fin,
		PreviousFields:    prev,
		PreviousTxnID:     h.PreviousTxnID,
		PreviousTxnLgrSeq: h.PreviousTxnLgrSeq,
	}
	return nil
}

type DeletedNode struct {
	LedgerEntryType ledger.LedgerEntryType `json:"LedgerEntryType,omitempty"`
	LedgerIndex     string                 `json:"LedgerIndex,omitempty"`
	FinalFields     ledger.LedgerObject    `json:"FinalFields,omitempty"`
}

func (n *DeletedNode) UnmarshalJSON(data []byte) error {
	var h struct {
		LedgerEntryType ledger.LedgerEntryType
		LedgerIndex     string
		FinalFields     json.RawMessage
	}
	err := json.Unmarshal(data, &h)
	if err != nil {
		return fmt.Errorf("unmarshal DeletedNode: %w", err)
	}
	var obj ledger.LedgerObject
	if obj, err = ledger.EmptyLedgerObject(string(h.LedgerEntryType)); err != nil {
		return fmt.Errorf("unmarshal DeletedNode: %w", err)
	}
	et := obj.EntryType()
	if h.FinalFields != nil && len(h.FinalFields) > 0 {
		if err = json.Unmarshal(h.FinalFields, obj); err != nil {
			return fmt.Errorf("unmarshal DeletedNode: %w", err)
		}
	} else {
		obj = nil
	}
	*n = DeletedNode{
		LedgerEntryType: et,
		LedgerIndex:     h.LedgerIndex,
		FinalFields:     obj,
	}
	return nil
}
