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
	Fee                string
	Sequence           uint
	AccountTxnID       Hash256    `json:",omitempty"`
	Flags              uint       `json:",omitempty"`
	LastLedgerSequence uint       `json:",omitempty"`
	Memos              []Memo     `json:",omitempty"`
	Signers            []TxSigner `json:",omitempty"`
	SourceTag          uint       `json:",omitempty"`
	SigningPubKey      []byte
	TicketSequence     uint `json:",omitempty"`
	TxnSignature       []byte
}

type Memo struct {
	MemoData   []byte
	MemoFormat []byte
	MemoType   []byte
}

type TxSigner struct {
	Account       Address
	TxnSignature  []byte
	SigningPubKey []byte
}

func (tx *BaseTx) TxType() TxType {
	return tx.TransactionType
}

type AccountSet struct {
	BaseTx
}

func (*AccountSet) TxType() TxType {
	return AccountSetTx
}

type AccountDelete struct {
	BaseTx
}

func (*AccountDelete) TxType() TxType {
	return AccountDeleteTx
}

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

type CheckCancel struct {
	BaseTx
}

func (*CheckCancel) TxType() TxType {
	return CheckCancelTx
}

type CheckCash struct {
	BaseTx
}

func (*CheckCash) TxType() TxType {
	return CheckCashTx
}

type CheckCreate struct {
	BaseTx
}

func (*CheckCreate) TxType() TxType {
	return CheckCreateTx
}

type DepositPreauth struct {
	BaseTx
}

func (*DepositPreauth) TxType() TxType {
	return DepositPreauthTx
}

type EscrowCancel struct {
	BaseTx
}

func (*EscrowCancel) TxType() TxType {
	return EscrowCancelTx
}

type EscrowCreate struct {
	BaseTx
}

func (*EscrowCreate) TxType() TxType {
	return EscrowCreateTx
}

type EscrowFinish struct {
	BaseTx
}

func (*EscrowFinish) TxType() TxType {
	return EscrowFinishTx
}

type NFTokenAcceptOffer struct {
	BaseTx
}

func (*NFTokenAcceptOffer) TxType() TxType {
	return NFTokenAcceptOfferTx
}

type NFTokenBurn struct {
	BaseTx
}

func (*NFTokenBurn) TxType() TxType {
	return NFTokenBurnTx
}

type NFTokenCancelOffer struct {
	BaseTx
}

func (*NFTokenCancelOffer) TxType() TxType {
	return NFTokenCancelOfferTx
}

type NFTokenCreateOffer struct {
	BaseTx
}

func (*NFTokenCreateOffer) TxType() TxType {
	return NFTokenCreateOfferTx
}

type NFTokenMint struct {
	BaseTx
}

func (*NFTokenMint) TxType() TxType {
	return NFTokenMintTx
}

type OfferCancel struct {
	BaseTx
}

func (*OfferCancel) TxType() TxType {
	return OfferCancelTx
}

type OfferCreate struct {
	BaseTx
}

func (*OfferCreate) TxType() TxType {
	return OfferCreateTx
}

type Payment struct {
	BaseTx
}

func (*Payment) TxType() TxType {
	return PaymentTx
}

type PaymentChannelClaim struct {
	BaseTx
}

func (*PaymentChannelClaim) TxType() TxType {
	return PaymentChannelClaimTx
}

type PaymentChannelCreate struct {
	BaseTx
}

func (*PaymentChannelCreate) TxType() TxType {
	return PaymentChannelCreateTx
}

type PaymentChannelFund struct {
	BaseTx
}

func (*PaymentChannelFund) TxType() TxType {
	return PaymentChannelFundTx
}

type SetRegularKey struct {
	BaseTx
}

func (*SetRegularKey) TxType() TxType {
	return SetRegularKeyTx
}

type SignerListSet struct {
	BaseTx
}

func (*SignerListSet) TxType() TxType {
	return SignerListSetTx
}

type TicketCreate struct {
	BaseTx
}

func (*TicketCreate) TxType() TxType {
	return TicketCreateTx
}

type TrustSet struct {
	BaseTx
}

func (*TrustSet) TxType() TxType {
	return TrustSetTx
}

func UnmarshalTx(data json.RawMessage) (Tx, error) {
	// TODO
	return nil, fmt.Errorf("Unimplemented")
}
