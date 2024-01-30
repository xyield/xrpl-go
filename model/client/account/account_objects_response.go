package account

import (
	"encoding/json"
	"fmt"

	"github.com/CreatureDev/xrpl-go/model/client/common"
	"github.com/CreatureDev/xrpl-go/model/ledger"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

const (
	ErrAccountObjectUnmarshal string = "Unmarshal JSON AccountObjects"
)

type AccountObjectsResponse struct {
	Account            types.Address         `json:"account"`
	AccountObjects     []ledger.LedgerObject `json:"account_objects"`
	LedgerHash         common.LedgerHash     `json:"ledger_hash,omitempty"`
	LedgerIndex        common.LedgerIndex    `json:"ledger_index,omitempty"`
	LedgerCurrentIndex common.LedgerIndex    `json:"ledger_current_index,omitempty"`
	Limit              int                   `json:"limit,omitempty"`
	Marker             any                   `json:"marker,omitempty"`
	Validated          bool                  `json:"validated,omitempty"`
}

func (r *AccountObjectsResponse) UnmarshalJSON(data []byte) error {
	type accountObjectDecodeHelper struct {
		Account            types.Address      `json:"account"`
		AccountObjects     []json.RawMessage  `json:"account_objects"`
		LedgerHash         common.LedgerHash  `json:"ledger_hash"`
		LedgerIndex        common.LedgerIndex `json:"ledger_index"`
		LedgerCurrentIndex common.LedgerIndex `json:"ledger_current_index"`
		Limit              int                `json:"limit"`
		Marker             any                `json:"marker"`
		Validated          bool               `json:"validated"`
	}
	var values accountObjectDecodeHelper
	if err := json.Unmarshal(data, &values); err != nil {
		return fmt.Errorf("%s: %w", ErrAccountObjectUnmarshal, err)
	}
	*r = AccountObjectsResponse{
		Account:            values.Account,
		LedgerHash:         values.LedgerHash,
		LedgerIndex:        values.LedgerIndex,
		LedgerCurrentIndex: values.LedgerCurrentIndex,
		Limit:              values.Limit,
		Marker:             values.Marker,
		Validated:          values.Validated,
	}
	r.AccountObjects = make([]ledger.LedgerObject, len(values.AccountObjects))
	for i, v := range values.AccountObjects {
		var err error
		r.AccountObjects[i], err = ledger.UnmarshalLedgerObject(v)
		if err != nil {
			return fmt.Errorf("%s: %w", ErrAccountObjectUnmarshal, err)
		}
	}

	return nil
}
