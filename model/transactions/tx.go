package transactions

import (
	"encoding/json"
	"fmt"

	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type Tx interface {
	TxType() TxType
}

type TxHash string

func (*TxHash) TxType() TxType {
	return HashedTx
}

type Binary struct {
	TxBlob string `json:"tx_blob"`
}

func (tx *Binary) TxType() TxType {
	return BinaryTx
}

type BaseTx struct {
	Account            types.Address
	TransactionType    TxType
	Fee                types.XRPCurrencyAmount `json:",omitempty"`
	Sequence           uint32                  `json:",omitempty"`
	AccountTxnID       types.Hash256           `json:",omitempty"`
	Flags              *types.Flag             `json:",omitempty"`
	LastLedgerSequence uint32                  `json:",omitempty"`
	Memos              []MemoWrapper           `json:",omitempty"`
	Signers            []Signer                `json:",omitempty"`
	SourceTag          uint                    `json:",omitempty"`
	SigningPubKey      string                  `json:",omitempty"`
	TicketSequence     uint32                  `json:",omitempty"`
	TxnSignature       string                  `json:",omitempty"`
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
	if len(data) == 0 {
		return nil, nil
	}
	if data[0] == '"' {
		var ret TxHash
		if err := json.Unmarshal(data, &ret); err != nil {
			return nil, err
		}
		return &ret, nil
	} else if data[0] != '{' {
		// TODO error verbosity/record failed json
		return nil, fmt.Errorf("unexpected tx format; must be tx object or hash string")
	}
	// TODO AMM endpoint support
	type txTypeParser struct {
		TransactionType TxType
		TxBlob          string `json:"tx_blob"`
	}
	var txType txTypeParser
	if err := json.Unmarshal(data, &txType); err != nil {
		return nil, err
	}
	if len(txType.TxBlob) > 0 && len(txType.TransactionType) == 0 {
		return &Binary{
			TxBlob: txType.TxBlob,
		}, nil
	}
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
		return nil, fmt.Errorf("unsupported transaction type %s", txType.TransactionType)
	}
	if err := json.Unmarshal(data, tx); err != nil {
		return nil, err
	}
	return tx, nil
}
