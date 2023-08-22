package signing

import (
	"encoding/json"
	"fmt"

	"github.com/xyield/xrpl-go/model/transactions"
)

type SignRequest struct {
	TxJson     transactions.Tx `json:"tx_json"`
	Secret     string          `json:"secret,omitempty"`
	Seed       string          `json:"seed,omitempty"`
	SeedHex    string          `json:"seed_hex,omitempty"`
	Passphrase string          `json:"passphrase,omitempty"`
	KeyType    string          `json:"key_type,omitempty"`
	Offline    bool            `json:"offline,omitempty"`
	BuildPath  bool            `json:"build_path,omitempty"`
	FeeMultMax int             `json:"fee_mult_max,omitempty"`
	FeeDivMax  int             `json:"fee_div_max,omitempty"`
}

func (*SignRequest) Method() string {
	return "sign"
}

func (r *SignRequest) Validate() error {
	if r.TxJson == nil {
		return fmt.Errorf("sign request: empty tx")
	}

	return nil
}

func (r *SignRequest) UnmarshalJSON(data []byte) error {
	type srHelper struct {
		TxJson     json.RawMessage `json:"tx_json"`
		Secret     string          `json:"secret,omitempty"`
		Seed       string          `json:"seed,omitempty"`
		SeedHex    string          `json:"seed_hex,omitempty"`
		Passphrase string          `json:"passphrase,omitempty"`
		KeyType    string          `json:"key_type,omitempty"`
		Offline    bool            `json:"offline,omitempty"`
		BuildPath  bool            `json:"build_path,omitempty"`
		FeeMultMax int             `json:"fee_mult_max,omitempty"`
		FeeDivMax  int             `json:"fee_div_max,omitempty"`
	}
	var h srHelper
	err := json.Unmarshal(data, &h)
	*r = SignRequest{
		Secret:     h.Secret,
		Seed:       h.Seed,
		SeedHex:    h.SeedHex,
		Passphrase: h.Passphrase,
		KeyType:    h.KeyType,
		Offline:    h.Offline,
		BuildPath:  h.BuildPath,
		FeeMultMax: h.FeeMultMax,
		FeeDivMax:  h.FeeDivMax,
	}
	if err != nil {
		return err
	}
	r.TxJson, err = transactions.UnmarshalTx(h.TxJson)
	return err
}
