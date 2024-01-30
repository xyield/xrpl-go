package key

import "fmt"

type WalletProposeRequest struct {
	KeyType    string `json:"key_type,omitempty"`
	Passphrase string `json:"passphrase,omitempty"`
	Seed       string `json:"seed,omitempty"`
	SeedHex    string `json:"seed_hex,omitempty"`
}

func (*WalletProposeRequest) Method() string {
	return "wallet_propose"
}

func (p *WalletProposeRequest) Validate() error {
	cnt := 0
	if p.Passphrase != "" {
		cnt++
	}
	if p.Seed != "" {
		cnt++
	}
	if p.SeedHex != "" {
		cnt++
	}
	if cnt > 1 {
		return fmt.Errorf("wallet propose request: only one of (passphrase, seed, seedhex) may be set")
	}
	return nil
}
