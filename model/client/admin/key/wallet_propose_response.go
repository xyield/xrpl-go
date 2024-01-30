package key

import "github.com/CreatureDev/xrpl-go/model/transactions/types"

type WalletProposeResponse struct {
	KeyType       string        `json:"key_type"`
	MasterSeed    string        `json:"master_seed"`
	MasterSeedHex string        `json:"master_seed_hex"`
	AccountId     types.Address `json:"account_id"`
	PublicKey     string        `json:"public_key"`
	PublicKeyHex  string        `json:"public_key_hex"`
	Warning       string        `json:"warning,omitempty"`
}
