package common

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type LedgerSpecifier interface {
	Ledger() string
}

func UnmarshalLedgerSpecifier(data []byte) (LedgerSpecifier, error) {
	if len(data) == 0 {
		return nil, nil
	}
	switch data[0] {
	case '"':
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return nil, err
		}
		switch s {
		case CURRENT.Ledger():
			return CURRENT, nil
		case VALIDATED.Ledger():
			return VALIDATED, nil
		case CLOSED.Ledger():
			return CLOSED, nil
		}
		return nil, fmt.Errorf("decoding LedgerTitle: invalid string %s", s)
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
	VALIDATED LedgerTitle = "validated"
	CLOSED    LedgerTitle = "closed"
)

func (l LedgerTitle) Ledger() string {
	return string(l)
}

type LedgerHash string
