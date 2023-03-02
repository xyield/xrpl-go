package transactions

import (
	"encoding/json"
	"fmt"

	"github.com/xyield/xrpl-go/model/ledger"
	"github.com/xyield/xrpl-go/model/transactions/types"
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
	AffectedNodes          []AffectedNode       `json:"AffectedNodes"`
	PartialDeliveredAmount types.CurrencyAmount `json:"DeliveredAmount"`
	TransactionIndex       uint64               `json:"TransactionIndex"`
	TransactionResult      string               `json:"TransactionResult"`
	DeliveredAmount        types.CurrencyAmount `json:"delivered_amount"`
}

func (m *TxObjMeta) UnmarshalJSON(data []byte) error {
	var h struct {
		AffectedNodes          []AffectedNode  `json:"AffectedNodes"`
		PartialDeliveredAmount json.RawMessage `json:"DeliveredAmount"`
		TransactionIndex       uint64          `json:"TransactionIndex"`
		TransactionResult      string          `json:"TransactionResult"`
		DeliveredAmount        json.RawMessage `json:"delivered_amount"`
	}
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*m = TxObjMeta{
		AffectedNodes:     h.AffectedNodes,
		TransactionIndex:  h.TransactionIndex,
		TransactionResult: h.TransactionResult,
	}
	if part, err := types.UnmarshalCurrencyAmount(h.PartialDeliveredAmount); err != nil {
		return err
	} else {
		m.PartialDeliveredAmount = part
	}
	if deliv, err := types.UnmarshalCurrencyAmount(h.DeliveredAmount); err != nil {
		return err
	} else {
		m.DeliveredAmount = deliv
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
		return err
	}
	var obj ledger.LedgerObject
	if obj, err = ledger.EmptyLedgerObject(string(h.LedgerEntryType)); err != nil {
		return err
	}
	if err = json.Unmarshal(h.NewFields, obj); err != nil {
		return err
	}
	*n = CreatedNode{
		LedgerEntryType: h.LedgerEntryType,
		LedgerIndex:     h.LedgerIndex,
		NewFields:       obj,
	}
	return nil
}

type ModifiedNode struct {
	LedgerEntryType   ledger.LedgerEntryType `json:"LedgerEntryType"`
	LedgerIndex       string                 `json:"LedgerIndex"`
	FinalFields       ledger.LedgerObject    `json:"FinalFields"`
	PreviousFields    ledger.LedgerObject    `json:"PreviousFields"`
	PreviousTxnID     string                 `json:"PreviousTxnID,omitempty"`
	PreviousTxnLgrSeq uint64                 `json:"PreviousTxnLgrSeq,omitempty"`
}

func (n *ModifiedNode) UnmarshalJSON(data []byte) error {
	var h struct {
		LedgerEntryType   ledger.LedgerEntryType
		LedgerIndex       string
		FinalFields       json.RawMessage
		PreviousFields    json.RawMessage
		PreviousTxnID     string
		PreviousTxnLgrSeq uint64
	}
	err := json.Unmarshal(data, &h)
	if err != nil {
		return err
	}
	var prev, fin ledger.LedgerObject
	if prev, err = ledger.EmptyLedgerObject(string(h.LedgerEntryType)); err != nil {
		return err
	}
	if err = json.Unmarshal(h.PreviousFields, prev); err != nil {
		return err
	}
	if err = json.Unmarshal(h.FinalFields, fin); err != nil {
		return err
	}
	*n = ModifiedNode{
		LedgerEntryType:   h.LedgerEntryType,
		LedgerIndex:       h.LedgerIndex,
		FinalFields:       fin,
		PreviousFields:    prev,
		PreviousTxnID:     h.PreviousTxnID,
		PreviousTxnLgrSeq: h.PreviousTxnLgrSeq,
	}
	return nil
}

type DeletedNode struct {
	LedgerEntryType ledger.LedgerEntryType `json:"LedgerEntryType"`
	LedgerIndex     string                 `json:"LedgerIndex"`
	FinalFields     ledger.LedgerObject    `json:"FinalFields"`
}

func (n *DeletedNode) UnmarshalJSON(data []byte) error {
	var h struct {
		LedgerEntryType ledger.LedgerEntryType
		LedgerIndex     string
		FinalFields     json.RawMessage
	}
	err := json.Unmarshal(data, &h)
	if err != nil {
		return err
	}
	var obj ledger.LedgerObject
	if obj, err = ledger.EmptyLedgerObject(string(h.LedgerEntryType)); err != nil {
		return err
	}
	if err = json.Unmarshal(h.FinalFields, obj); err != nil {
		return err
	}
	*n = DeletedNode{
		LedgerEntryType: h.LedgerEntryType,
		LedgerIndex:     h.LedgerIndex,
		FinalFields:     obj,
	}
	return nil
}
