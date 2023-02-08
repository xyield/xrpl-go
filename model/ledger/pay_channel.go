package ledger

import . "github.com/xyield/xrpl-go/model/transactions/types"

type PayChannel struct {
	Account           Address
	Amount            XRPCurrencyAmount
	Balance           XRPCurrencyAmount
	CancelAfter       uint `json:",omitempty"`
	Destination       Address
	DestinationTag    uint   `json:",omitempty"`
	DestinationNode   string `json:",omitempty"`
	Expiration        uint   `json:",omitempty"`
	Flags             uint
	LedgerEntryType   LedgerEntryType
	OwnerNode         string
	PreviousTxnID     Hash256
	PreviousTxnLgrSeq uint
	PublicKey         string
	SettleDelay       uint
	SourceTag         uint `json:",omitempty"`
}

func (*PayChannel) EntryType() LedgerEntryType {
	return PayChannelEntry
}
