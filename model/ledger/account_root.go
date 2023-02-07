package ledger

import (
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountRoot struct {
	Account           Address `json:"Account"`
	AccountTxnID      Hash256 `json:"AccountTxnID"`
	Balance           string  `json:"Balance"`
	BurnedNFTokens    uint32  `json:"BurnedNFTokens"`
	Domain            string  `json:"Domain"`
	EmailHash         Hash128 `json:"EmailHash"`
	Flags             uint64  `json:"Flags"`
	LedgerEntryType   string  `json:"LedgerEntryType"`
	MessageKey        string  `json:"MessageKey"`
	MintedNFTokens    uint32  `json:"MintedNFTokens"`
	NFTokenMinter     Address `json:"NFTokenMinter"`
	OwnerCount        uint64  `json:"OwnerCount"`
	PreviousTxnID     Hash256 `json:"PreviousTxnID"`
	PreviousTxnLgrSeq uint64  `json:"PreviousTxnLgrSeq"`
	RegularKey        Address `json:"RegularKey"`
	Sequence          uint64  `json:"Sequence"`
	TicketCount       uint32  `json:"TicketCount"`
	TickSize          uint8   `json:"TickSize"`
	TransferRate      uint32  `json:"TransferRate"`
	Index             Hash256 `json:"index"`
}

func (*AccountRoot) EntryType() LedgerEntryType {
	return AccountRootEntry
}
