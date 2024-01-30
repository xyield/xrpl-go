package signing

import (
	"encoding/json"
	"fmt"

	"github.com/CreatureDev/xrpl-go/model/transactions"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type SignForRequest struct {
	Account    types.Address   `json:"account"`
	TxJson     transactions.Tx `json:"tx_json"`
	Secret     string          `json:"secret,omitempty"`
	Seed       string          `json:"seed,omitempty"`
	SeedHex    string          `json:"seed_hex,omitempty"`
	Passphrase string          `json:"passphrase,omitempty"`
	KeyType    string          `json:"key_type,omitempty"`
}

func (*SignForRequest) Method() string {
	return "sign_for"
}

func (r *SignForRequest) Validate() error {
	if err := r.Account.Validate(); err != nil {
		return fmt.Errorf("sign for request: %w", err)
	}

	if r.TxJson == nil {
		return fmt.Errorf("sign for request: empty tx")
	}

	cnt := 0
	if r.Secret != "" {
		cnt++
	}
	if r.Seed != "" {
		cnt++
	}
	if r.SeedHex != "" {
		cnt++
	}
	if r.Passphrase != "" {
		cnt++
	}
	if cnt != 1 {
		return fmt.Errorf("sign for request: must provide one of (secret, seed, seedhex, passphrase)")
	}

	return nil
}

func (r *SignForRequest) UnmarshalJSON(data []byte) error {
	type srHelper struct {
		Account    types.Address   `json:"account"`
		TxJson     json.RawMessage `json:"tx_json"`
		Secret     string          `json:"secret,omitempty"`
		Seed       string          `json:"seed,omitempty"`
		SeedHex    string          `json:"seed_hex,omitempty"`
		Passphrase string          `json:"passphrase,omitempty"`
		KeyType    string          `json:"key_type,omitempty"`
	}
	var h srHelper
	err := json.Unmarshal(data, &h)
	*r = SignForRequest{
		Account:    h.Account,
		Secret:     h.Secret,
		Seed:       h.Seed,
		SeedHex:    h.SeedHex,
		Passphrase: h.Passphrase,
		KeyType:    h.KeyType,
	}
	if err != nil {
		return err
	}
	r.TxJson, err = transactions.UnmarshalTx(h.TxJson)
	return err
}
