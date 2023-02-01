package transactions

import (
	"encoding/json"
	"fmt"
)

type Tx interface {
	TxType() TxType
}

type BaseTx struct {
	Account            Address
	TransactionType    TxType
	Fee                XrpCurrencyAmount
	Sequence           uint
	AccountTxnID       Hash256  `json:",omitempty"`
	Flags              uint     `json:",omitempty"`
	LastLedgerSequence uint     `json:",omitempty"`
	Memos              []Memo   `json:",omitempty"`
	Signers            []Signer `json:",omitempty"`
	SourceTag          uint     `json:",omitempty"`
	SigningPubKey      []byte
	TicketSequence     uint `json:",omitempty"`
	TxnSignature       []byte
}

func (tx *BaseTx) TxType() TxType {
	return tx.TransactionType
}

// TODO AMM support
type AMMBid struct {
	BaseTx
}

func (*AMMBid) TxType() TxType {
	return AMMBidTx
}

type AMMCreate struct {
	BaseTx
}

func (*AMMCreate) TxType() TxType {
	return AMMCreateTx
}

type AMMDeposit struct {
	BaseTx
}

func (*AMMDeposit) TxType() TxType {
	return AMMDepositTx
}

type AMMVote struct {
	BaseTx
}

func (*AMMVote) TxType() TxType {
	return AMMVoteTx
}

type AMMWithdraw struct {
	BaseTx
}

func (*AMMWithdraw) TxType() TxType {
	return AMMWithdrawTx
}

func UnmarshalTx(data json.RawMessage) (Tx, error) {
	// TODO
	return nil, fmt.Errorf("Unimplemented")
}
