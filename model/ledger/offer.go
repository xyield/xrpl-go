package ledger

import . "github.com/xyield/xrpl-go/model/transactions/types"

type OfferFlags uint

const (
	PassiveOffer OfferFlags = 0x00010000
	SellOffer               = 0x00020000
)

// TODO Unmarshal CurrencyAmounts
type Offer struct {
	Account           Address
	BookDirectory     Hash256
	BookNode          string
	Expiration        uint
	Flags             OfferFlags
	LedgerEntryType   string
	OwnerNode         string
	PreviousTxnID     Hash256
	PreviousTxnLgrSeq uint
	Sequence          uint
	TakerPays         CurrencyAmount
	TakerGets         CurrencyAmount
}

func (*Offer) EntryType() LedgerEntryType {
	return OfferEntry
}
