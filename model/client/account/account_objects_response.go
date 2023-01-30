package account

import (
	"encoding/json"
	"fmt"

	. "github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/ledger"
	. "github.com/xyield/xrpl-go/model/ledger"
	. "github.com/xyield/xrpl-go/model/transactions"
)

const (
	AccountObjectUnmarshalErr string = "Unmarshal JSON AccountObjects"
)

type AccountObjectsResponse struct {
	Account            Address        `json:"account"`
	AccountObjects     []LedgerObject `json:"account_objects"`
	LedgerHash         LedgerHash     `json:"ledger_hash"`
	LedgerIndex        LedgerIndex    `json:"ledger_index"`
	LedgerCurrentIndex LedgerIndex    `json:"ledger_current_index"`
	Limit              int            `json:"limit"`
	Marker             interface{}    `json:"marker"`
	Validated          bool           `json:"validated"`
}

func (r *AccountObjectsResponse) UnmarshalJSON(data []byte) error {
	// TODO Unmrashal LedgerObject interface
	type accountObjectDecodeHelper struct {
		Account            Address           `json:"account"`
		AccountObjects     []json.RawMessage `json:"account_objects"`
		LedgerHash         LedgerHash        `json:"ledger_hash"`
		LedgerIndex        LedgerIndex       `json:"ledger_index"`
		LedgerCurrentIndex LedgerIndex       `json:"ledger_current_index"`
		Limit              int               `json:"limit"`
		Marker             interface{}       `json:"marker"`
		Validated          bool              `json:"validated"`
	}
	var values accountObjectDecodeHelper
	if err := json.Unmarshal(data, &values); err != nil {
		return fmt.Errorf("%s: %w", AccountObjectUnmarshalErr, err)
	}
	r.Account = values.Account
	r.LedgerHash = values.LedgerHash
	r.LedgerIndex = values.LedgerIndex
	r.LedgerCurrentIndex = values.LedgerCurrentIndex
	r.Limit = values.Limit
	r.Marker = values.Marker
	r.Validated = values.Validated

	r.AccountObjects = make([]LedgerObject, len(values.AccountObjects))
	for i, v := range values.AccountObjects {
		var err error
		r.AccountObjects[i], err = ledger.UnmarshalLedgerObject(v)
		if err != nil {
			return fmt.Errorf("%s: %w", AccountObjectUnmarshalErr, err)
		}
	}

	return nil
}
