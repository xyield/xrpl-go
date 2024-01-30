package ledger

import (
	"encoding/json"

	"github.com/CreatureDev/xrpl-go/model/client/common"
	"github.com/CreatureDev/xrpl-go/model/ledger"
)

type LedgerDataResponse struct {
	LedgerIndex string            `json:"ledger_index"`
	LedgerHash  common.LedgerHash `json:"ledger_hash"`
	State       []LedgerState     `json:"state"`
	Marker      any               `json:"marker"`
}

type LedgerState struct {
	Data            string                 `json:"data,omitempty"`
	LedgerEntryType ledger.LedgerEntryType `json:",omitempty"`
	LedgerObject    ledger.LedgerObject    `json:"-"`
	Index           string                 `json:"index"`
}

func (l LedgerState) MarshalJSON() ([]byte, error) {
	type lsHelper LedgerState
	c := lsHelper(l)
	c.LedgerObject = nil
	first, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	second, err := json.Marshal(l.LedgerObject)
	if err != nil {
		return nil, err
	}
	data := make(map[string]interface{})
	err = json.Unmarshal(first, &data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(second, &data)
	if err != nil {
		return nil, err
	}
	return json.Marshal(data)
}

func (l *LedgerState) UnmarshalJSON(data []byte) error {
	type lsHelper LedgerState
	h := lsHelper(*l)
	err := json.Unmarshal(data, &h)
	if err != nil {
		return err
	}
	*l = LedgerState(h)
	l.LedgerObject, err = ledger.UnmarshalLedgerObject(data)
	return err
}
