package signing

import (
	"encoding/json"

	"github.com/xyield/xrpl-go/model/transactions"
	"github.com/xyield/xrpl-go/model/transactions/types"
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
	return "sign"
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
