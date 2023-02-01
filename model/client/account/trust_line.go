package account

import (
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type TrustLine struct {
	Account        Address `json:"account"`
	Balance        string  `json:"balance"`
	Currency       string  `json:"currency"`
	Limit          string  `json:"limit"`
	LimitPeer      string  `json:"limit_peer"`
	QualityIn      uint    `json:"quality_in"`
	QualityOut     uint    `json:"quality_out"`
	NoRipple       bool    `json:"no_ripple"`
	NoRipplePeer   bool    `json:"no_ripple_peer"`
	Authorized     bool    `json:"authorized"`
	PeerAuthorized bool    `json:"peer_authorized"`
	Freeze         bool    `json:"freeze"`
	FreezePeer     bool    `json:"freeze_peer"`
}
