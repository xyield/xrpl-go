package ledger

import (
	"encoding/json"

	"github.com/xyield/xrpl-go/model/transactions/types"
)

type AMM struct {
	Asset          AMMAsset
	Asset2         AMMAsset
	AMMAccount     types.Address
	AuctionSlot    AMMAuctionSlot `json:",omitempty"`
	LPTokenBalance types.CurrencyAmount
	TradingFee     uint16
	VoteSlots      []AMMVoteEntry `json:",omitempty"`
}

func (a *AMM) UnmarshalJSON(data []byte) error {
	type ammHelper struct {
		Asset          AMMAsset
		Asset2         AMMAsset
		AMMAccount     types.Address
		AuctionSlot    AMMAuctionSlot
		LPTokenBalance json.RawMessage
		TradingFee     uint16
		VoteSlots      []AMMVoteEntry
	}
	var h ammHelper
	var err error
	if err = json.Unmarshal(data, &h); err != nil {
		return err
	}
	*a = AMM{
		Asset:       h.Asset,
		Asset2:      h.Asset2,
		AMMAccount:  h.AMMAccount,
		AuctionSlot: h.AuctionSlot,
		TradingFee:  h.TradingFee,
		VoteSlots:   h.VoteSlots,
	}

	a.LPTokenBalance, err = types.UnmarshalCurrencyAmount(h.LPTokenBalance)
	if err != nil {
		return err
	}

	return nil
}

type AMMAsset struct {
	Currency string
	Issuer   types.Address
}

type AMMAuctionSlot struct {
	Account       types.Address
	AuthAccounts  []AMMAuthAccount `json:",omitempty"`
	DiscountedFee int
	Price         types.CurrencyAmount
	Expiration    uint
}

func (s *AMMAuctionSlot) UnmarshalJSON(data []byte) error {
	type aasHelper struct {
		Account       types.Address
		AuthAccounts  []AMMAuthAccount
		DiscountedFee int
		Price         json.RawMessage
		Expiration    uint
	}
	var h aasHelper
	var err error
	if err = json.Unmarshal(data, &h); err != nil {
		return err
	}
	*s = AMMAuctionSlot{
		Account:       h.Account,
		AuthAccounts:  h.AuthAccounts,
		DiscountedFee: h.DiscountedFee,
		Expiration:    h.Expiration,
	}

	s.Price, err = types.UnmarshalCurrencyAmount(h.Price)
	if err != nil {
		return err
	}
	return nil
}

type AMMAuthAccount struct {
	Account types.Address
}

type AMMVoteEntry struct {
	Account     types.Address
	TradingFee  uint
	VoteWeither uint
}
