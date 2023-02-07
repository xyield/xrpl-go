package ledger

import . "github.com/xyield/xrpl-go/model/transactions/types"

//TODO Unmarshal CurrencyAmounts

type AMM struct {
	Asset          AMMAsset
	Asset2         AMMAsset
	AMMAccount     Address
	AuctionSlot    AMMAuctionSlot
	LPTokenBalance CurrencyAmount
	TradingFee     uint16
	VoteSlots      []AMMVoteEntry
}

type AMMAsset struct {
	Currency string
	Issuer   Address
}

type AMMAuctionSlot struct {
	Account       Address
	AuthAccounts  []AMMAuthAccount
	DiscountedFee int
	Price         CurrencyAmount
	Expiration    uint
}

type AMMAuthAccount struct {
	Account Address
}

type AMMVoteEntry struct {
	Account     Address
	TradingFee  uint
	VoteWeither uint
}
