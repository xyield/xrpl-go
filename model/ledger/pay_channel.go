package ledger

import . "github.com/xyield/xrpl-go/model/transactions/types"

type PayChannel struct {
	Account           Address
	Destination       Address
	DestinationTag    uint
	Amount            XRPCurrencyAmount
	Balance           XRPCurrencyAmount
	CancelAfter       uint
	DestinationNode   string
	Expiration        uint
	Flags             uint
	LedgerEntryType   LedgerEntryType
	OwnerNode         string
	PreviousTxnID     Hash256
	PreviousTxnLgrSeq uint
	PublicKey         string
	SettleDelay       uint
	SourceTag         uint
}

func (*PayChannel) EntryType() LedgerEntryType {
	return PayChannelEntry
}
