package account

import (
	"encoding/json"
	"fmt"

	. "github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/ledger"
	. "github.com/xyield/xrpl-go/model/ledger"
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

const (
	ErrAccountObjectUnmarshal string = "Unmarshal JSON AccountObjects"
)

type AccountObjectsResponse struct {
	Account            Address        `json:"account"`
	AccountObjects     []LedgerObject `json:"account_objects"`
	LedgerHash         LedgerHash     `json:"ledger_hash,omitempty"`
	LedgerIndex        LedgerIndex    `json:"ledger_index,omitempty"`
	LedgerCurrentIndex LedgerIndex    `json:"ledger_current_index,omitempty"`
	Limit              int            `json:"limit,omitempty"`
	Marker             any            `json:"marker,omitempty"`
	Validated          bool           `json:"validated,omitempty"`
}

func (r *AccountObjectsResponse) UnmarshalJSON(data []byte) error {
	type accountObjectDecodeHelper struct {
		Account            Address           `json:"account"`
		AccountObjects     []json.RawMessage `json:"account_objects"`
		LedgerHash         LedgerHash        `json:"ledger_hash"`
		LedgerIndex        LedgerIndex       `json:"ledger_index"`
		LedgerCurrentIndex LedgerIndex       `json:"ledger_current_index"`
		Limit              int               `json:"limit"`
		Marker             any               `json:"marker"`
		Validated          bool              `json:"validated"`
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
	r.AccountObjects = make([]LedgerObject, len(values.AccountObjects))
	for i, v := range values.AccountObjects {
		var err error
		r.AccountObjects[i], err = ledger.UnmarshalLedgerObject(v)
		if err != nil {
			return fmt.Errorf("%s: %w", ErrAccountObjectUnmarshal, err)
		}
	}

	return nil
}
