package subscribe

import "github.com/xyield/xrpl-go/model/transactions/types"

type SubscribeRequest struct {
	Streams          []string             `json:"streams,omitempty"`
	Accounts         []types.Address      `json:"accounts,omitempty"`
	AccountsProposed []types.Address      `json:"accounts_proposed,omitempty"`
	Books            []SubscribeOrderBook `json:"books,omitempty"`
	Url              string               `json:"url,omitempty"`
	UrlUsername      string               `json:"url_username,omitempty"`
	UrlPassword      string               `json:"url_password,omitempty"`
}

func (*SubscribeRequest) Method() string {
	return "subscribe"
}

type SubscribeOrderBook struct {
	TakerGets types.IssuedCurrencyAmount `json:"taker_gets"`
	TakerPays types.IssuedCurrencyAmount `json:"taker_pays"`
	Taker     types.Address              `json:"taker"`
	Snapshot  bool                       `json:"snapshot,omitempty"`
	Both      bool                       `json:"both,omitempty"`
}
