package ledger

import "github.com/CreatureDev/xrpl-go/model/transactions/types"

type PayChannel struct {
	Account           types.Address
	Amount            types.XRPCurrencyAmount
	Balance           types.XRPCurrencyAmount
	CancelAfter       uint `json:",omitempty"`
	Destination       types.Address
	DestinationTag    uint   `json:",omitempty"`
	DestinationNode   string `json:",omitempty"`
	Expiration        uint   `json:",omitempty"`
	Flags             uint32
	LedgerEntryType   LedgerEntryType
	OwnerNode         string
	PreviousTxnID     types.Hash256
	PreviousTxnLgrSeq uint32
	PublicKey         string
	SettleDelay       uint
	SourceTag         uint `json:",omitempty"`
}

func (*PayChannel) EntryType() LedgerEntryType {
	return PayChannelEntry
}
