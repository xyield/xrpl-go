package transactions

import (
	"encoding/json"
	"fmt"

	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type Tx interface {
	TxType() TxType
}

type BaseTx struct {
	Account            Address
	TransactionType    TxType
	Fee                XRPCurrencyAmount
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
	// TODO AMM endpoint support
	type txTypeParser struct {
		TransactionType TxType
	}
	var txType txTypeParser
	json.Unmarshal(data, &txType)
	var tx Tx
	switch txType.TransactionType {
	case AccountSetTx:
		tx = &AccountSet{}
	case AccountDeleteTx:
		tx = &AccountDelete{}
	case CheckCancelTx:
		tx = &CheckCancel{}
	case CheckCashTx:
		tx = &CheckCash{}
	case CheckCreateTx:
		tx = &CheckCreate{}
	case DepositPreauthTx:
		tx = &DepositPreauth{}
	case EscrowCancelTx:
		tx = &EscrowCancel{}
	case EscrowCreateTx:
		tx = &EscrowCreate{}
	case EscrowFinishTx:
		tx = &EscrowFinish{}
	case NFTokenAcceptOfferTx:
		tx = &NFTokenAcceptOffer{}
	case NFTokenBurnTx:
		tx = &NFTokenBurn{}
	case NFTokenCancelOfferTx:
		tx = &NFTokenCancelOffer{}
	case NFTokenCreateOfferTx:
		tx = &NFTokenCreateOffer{}
	case NFTokenMintTx:
		tx = &NFTokenMint{}
	case OfferCreateTx:
		tx = &OfferCreate{}
	case OfferCancelTx:
		tx = &OfferCancel{}
	case PaymentTx:
		tx = &Payment{}
	case PaymentChannelClaimTx:
		tx = &PaymentChannelClaim{}
	case PaymentChannelCreateTx:
		tx = &PaymentChannelCreate{}
	case PaymentChannelFundTx:
		tx = &PaymentChannelFund{}
	case SetRegularKeyTx:
		tx = &SetRegularKey{}
	case SignerListSetTx:
		tx = &SignerListSet{}
	case TrustSetTx:
		tx = &TrustSet{}
	case TicketCreateTx:
		tx = &TicketCreate{}
	default:
		return nil, fmt.Errorf("Unsupported transaction type %s", txType.TransactionType)
	}
	if err := json.Unmarshal(data, tx); err != nil {
		return nil, err
	}
	return tx, nil
}
