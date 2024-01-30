package account

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/model/client/common"
	"github.com/CreatureDev/xrpl-go/test"
)

func TestAccountNFTsRequest(t *testing.T) {
	s := AccountNFTsRequest{
		Account:     "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
		LedgerIndex: common.VALIDATED,
		LedgerHash:  "123",
		Limit:       2,
	}

	j := `{
	"account": "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
	"ledger_index": "validated",
	"ledger_hash": "123",
	"limit": 2
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestAccountNFTsResponse(t *testing.T) {
	s := AccountNFTsResponse{
		Account: "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
		AccountNFTs: []NFT{
			{Flags: Burnable | OnlyXRP,
				Issuer:       "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
				NFTokenID:    "abc",
				NFTokenTaxon: 123,
				URI:          "def",
				NFTSerial:    456,
			},
		},
		LedgerIndex:        123,
		LedgerHash:         "abc",
		LedgerCurrentIndex: 1234,
		Validated:          true,
	}

	j := `{
	"account": "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
	"account_nfts": [
		{
			"Flags": 3,
			"Issuer": "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
			"NFTokenID": "abc",
			"NFTokenTaxon": 123,
			"URI": "def",
			"nft_serial": 456
		}
	],
	"ledger_index": 123,
	"ledger_hash": "abc",
	"ledger_current_index": 1234,
	"validated": true
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
