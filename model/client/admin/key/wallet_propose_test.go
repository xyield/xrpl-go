package key

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/test"
)

func TestWalletProposeRequest(t *testing.T) {
	s := WalletProposeRequest{
		Seed:    "snoPBrXtMeMyMHUVTgbuqAfg1SUTb",
		KeyType: "secp256k1",
	}

	j := `{
	"key_type": "secp256k1",
	"seed": "snoPBrXtMeMyMHUVTgbuqAfg1SUTb"
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestWalletProposeResponse(t *testing.T) {
	s := WalletProposeResponse{
		KeyType:       "secp256k1",
		MasterSeed:    "snoPBrXtMeMyMHUVTgbuqAfg1SUTb",
		MasterSeedHex: "DEDCE9CE67B451D852FD4E846FCDE31C",
		AccountId:     "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
		PublicKey:     "aBQG8RQAzjs1eTKFEAQXr2gS4utcDiEC9wmi7pfUPTi27VCahwgw",
		PublicKeyHex:  "0330E7FC9D56BB25D6893BA3F317AE5BCF33B3291BD63DB32654A313222F7FD020",
	}

	j := `{
	"key_type": "secp256k1",
	"master_seed": "snoPBrXtMeMyMHUVTgbuqAfg1SUTb",
	"master_seed_hex": "DEDCE9CE67B451D852FD4E846FCDE31C",
	"account_id": "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
	"public_key": "aBQG8RQAzjs1eTKFEAQXr2gS4utcDiEC9wmi7pfUPTi27VCahwgw",
	"public_key_hex": "0330E7FC9D56BB25D6893BA3F317AE5BCF33B3291BD63DB32654A313222F7FD020"
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
