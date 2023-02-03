package common

import (
	"encoding/json"
	"strconv"

	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type LedgerSpecifier interface {
	Ledger() string
}

func UnmarshalLedgerSpecifier(data []byte) (LedgerSpecifier, error) {
	switch data[0] {
	case '"':
		var t LedgerTitle
		if err := json.Unmarshal(data, &t); err != nil {
			return nil, err
		}
		return t, nil
	default:
		var i LedgerIndex
		if err := json.Unmarshal(data, &i); err != nil {
			return nil, err
		}
		return i, nil
	}
}

type LedgerIndex uint

func (l LedgerIndex) Ledger() string {
	return strconv.FormatUint(uint64(l), 10)
}

type LedgerTitle string

const (
	CURRENT   LedgerTitle = "current"
	VALIDATED             = "validated"
	CLOSED                = "closed"
)

func (l LedgerTitle) Ledger() string {
	return string(l)
}

type LedgerHash Hash256

func (l LedgerHash) Ledger() string {
	return string(l)
}
