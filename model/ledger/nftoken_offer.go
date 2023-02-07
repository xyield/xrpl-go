package ledger

import . "github.com/xyield/xrpl-go/model/transactions/types"

type NFTokenOffer struct {
	Amount            CurrencyAmount
	Destination       Address
	Expiration        uint
	Flags             uint
	LedgerEntryType   LedgerEntryType
	NFTokenID         Hash256
	NFTokenOwnerNode  string
	Owner             Address
	OwnerNode         string
	PreviousTxnID     Hash256
	PreviousTxnLgrSeq uint
}

func (*NFTokenOffer) EntryType() LedgerEntryType {
	return NFTokenOfferEntry

}
