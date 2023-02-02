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
	// TODO AMM endpoint support
	type txTypeParser struct {
		TransactionType TxType
	}
	var txType txTypeParser
	json.Unmarshal(data, &txType)
	switch txType.TransactionType {
	case AccountSetTx:
		return UnmarshalAccountSetTx(data)
	case AccountDeleteTx:
		return UnmarshalAccountDeleteTx(data)
	case CheckCancelTx:
		return UnmarshalCheckCancelTx(data)
	case CheckCashTx:
		return UnmarshalCheckCashTx(data)
	case CheckCreateTx:
		return UnmarshalCheckCreateTx(data)
	case DepositPreauthTx:
		return UnmarshalDepositPreauthTx(data)
	case EscrowCancelTx:
		return UnmarshalEscrowCancelTx(data)
	case EscrowCreateTx:
		return UnmarshalEscrowCreateTx(data)
	case EscrowFinishTx:
		return UnmarshalEscrowFinishTx(data)
	case NFTokenAcceptOfferTx:
		return UnmarshalNFTokenAcceptOfferTx(data)
	case NFTokenBurnTx:
		return UnmarshalNFTokenBurnTx(data)
	case NFTokenCancelOfferTx:
		return UnmarshalNFTokenCancelOfferTx(data)
	case NFTokenCreateOfferTx:
		return UnmarshalNFTokenCreateOfferTx(data)
	case NFTokenMintTx:
		return UnmarshalNFTokenMintTx(data)
	case OfferCreateTx:
		return UnmarshalOfferCreateTx(data)
	case OfferCancelTx:
		return UnmarshalOfferCancelTx(data)
	case PaymentTx:
		return UnmarshalPaymentTx(data)
	case PaymentChannelClaimTx:
		return UnmarshalPaymentChannelClaimTx(data)
	case PaymentChannelCreateTx:
		return UnmarshalPaymentChannelCreateTx(data)
	case PaymentChannelFundTx:
		return UnmarshalPaymentChannelFundTx(data)
	case SetRegularKeyTx:
		return UnmarshalSetRegularKeyTx(data)
	case SignerListSetTx:
		return UnmarshalSignerListSetTx(data)
	case TrustSetTx:
		return UnmarshalTrustSetTx(data)
	case TicketCreateTx:
		return UnmarshalTicketCreateTx(data)
	}

	return nil, fmt.Errorf("Unsupported transaction type %s", txType.TransactionType)
}
