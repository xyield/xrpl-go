package ledger

import (
	"encoding/json"

	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/ledger"
)

type LedgerDataRequest struct {
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex common.LedgerSpecifier `json:"ledger_index,omitempty"`
	Binary      bool                   `json:"binary,omitempty"`
	Limit       int                    `json:"limit,omitempty"`
	Marker      any                    `json:"marker,omitempty"`
	Type        ledger.LedgerEntryType `json:"type,omitempty"`
}

func (r *LedgerDataRequest) UnmarshalJSON(data []byte) error {
	type ldrHelper struct {
		LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
		LedgerIndex json.RawMessage        `json:"ledger_index,omitempty"`
		Binary      bool                   `json:"binary,omitempty"`
		Limit       int                    `json:"limit,omitempty"`
		Marker      any                    `json:"marker,omitempty"`
		Type        ledger.LedgerEntryType `json:"type,omitempty"`
	}
	var h ldrHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = LedgerDataRequest{
		LedgerHash: h.LedgerHash,
		Binary:     h.Binary,
		Limit:      h.Limit,
		Marker:     h.Marker,
		Type:       h.Type,
	}
	i, err := common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = i

	return nil
}
