package types

import (
	"encoding/json"
	"strconv"
)

type CurrencyKind int

const (
	XRP CurrencyKind = iota
	ISSUED
)

type CurrencyAmount interface {
	Kind() CurrencyKind
}

func UnmarshalCurrencyAmount(data []byte) (CurrencyAmount, error) {
	if len(data) == 0 {
		return nil, nil
	}
	switch data[0] {
	case '{':
		var i IssuedCurrencyAmount
		if err := json.Unmarshal(data, &i); err != nil {
			return nil, err
		}
		return i, nil
	default:
		var x XRPCurrencyAmount
		if err := json.Unmarshal(data, &x); err != nil {
			return nil, err
		}
		return x, nil
	}
}

type IssuedCurrencyAmount struct {
	Issuer   Address `json:"issuer,omitempty"`
	Currency string  `json:"currency,omitempty"`
	Value    string  `json:"value,omitempty"`
}

func (IssuedCurrencyAmount) Kind() CurrencyKind {
	return ISSUED
}

type XRPCurrencyAmount uint64

func (XRPCurrencyAmount) Kind() CurrencyKind {
	return XRP
}

func (a XRPCurrencyAmount) MarshalJSON() ([]byte, error) {
	s := strconv.FormatUint(uint64(a), 10)
	return json.Marshal(s)
}

func (a *XRPCurrencyAmount) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return err
	}
	*a = XRPCurrencyAmount(v)
	return nil
}

func (a *XRPCurrencyAmount) UnmarshalText(data []byte) error {

	v, err := strconv.ParseUint(string(data), 10, 64)
	if err != nil {
		return err
	}
	*a = XRPCurrencyAmount(v)
	return nil
}
