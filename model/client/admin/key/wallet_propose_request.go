package key

type WalletProposeRequest struct {
	KeyType    string `json:"key_type,omitempty"`
	Passphrase string `json:"passphrase,omitempty"`
	Seed       string `json:"seed,omitempty"`
	SeedHex    string `json:"seed_hex,omitempty"`
}

func (*WalletProposeRequest) Method() string {
	return "wallet_propose"
}
