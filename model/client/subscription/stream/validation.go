package stream

import (
	"github.com/CreatureDev/xrpl-go/model/client/common"
)

type ValidationStream struct {
	Type                StreamType         `json:"type"`
	Amendments          []string           `json:"amendments,omitempty"`
	BaseFee             int                `json:"base_fee,omitempty"`
	Cookie              string             `json:"cookie,omitempty"`
	Flags               uint               `json:"flags"`
	Full                bool               `json:"full"`
	LedgerHash          common.LedgerHash  `json:"ledger_hash"`
	LedgerIndex         common.LedgerIndex `json:"ledger_index"`
	LoadFee             int                `json:"load_fee,omitempty"`
	MasterKey           string             `json:"master_key,omitempty"`
	ReserveBase         int                `json:"reserve_base,omitempty"`
	ReserveInc          int                `json:"reserve_inc,omitempty"`
	ServerVersion       string             `json:"server_version,omitempty"`
	Signature           string             `json:"signature"`
	SigningTime         uint64             `json:"singing_time"`
	ValidatedHash       string             `json:"validated_hash"`
	ValidationPublicKey string             `json:"validation_public_key"`
	// TODO validation public key as base58 string
}
