package account

import (
	"encoding/json"

	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountChannelsResponse struct {
	Account     types.Address      `json:"account"`
	Channels    []ChannelResult    `json:"channels"`
	LedgerIndex common.LedgerIndex `json:"ledger_index,omitempty"`
	LedgerHash  common.LedgerHash  `json:"ledger_hash,omitempty"`
	Validated   bool               `json:"validated,omitempty"`
	Limit       int                `json:"limit,omitempty"`
	Marker      any                `json:"marker,omitempty"`
}

func (r *AccountChannelsResponse) UnmarshallJSON(data []byte) error {
	type Alias AccountChannelsResponse
	var aux Alias
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*r = AccountChannelsResponse(aux)
	return nil
}
