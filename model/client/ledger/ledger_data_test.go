package ledger

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/model/client/common"
	"github.com/CreatureDev/xrpl-go/model/ledger"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/CreatureDev/xrpl-go/test"
)

func TestLedgerDataRequest(t *testing.T) {
	s := LedgerDataRequest{
		LedgerIndex: common.CLOSED,
		Binary:      true,
		Limit:       5,
	}
	j := `{
	"ledger_index": "closed",
	"binary": true,
	"limit": 5
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestLedgerDataResponse(t *testing.T) {
	s := LedgerDataResponse{
		LedgerIndex: "6885842",
		LedgerHash:  "842B57C1CC0613299A686D3E9F310EC0422C84D3911E5056389AA7E5808A93C8",
		State: []LedgerState{
			{
				LedgerEntryType: ledger.AccountRootEntry,
				LedgerObject: &ledger.AccountRoot{
					Account:           "rKKzk9ghA2iuy3imqMXUHJqdRPMtNDGf4c",
					Balance:           types.XRPCurrencyAmount(893730848),
					LedgerEntryType:   ledger.AccountRootEntry,
					Flags:             types.SetFlag(0),
					PreviousTxnID:     "C204A65CF2542946289A3358C67D991B5E135FABFA89F271DBA7A150C08CA046",
					PreviousTxnLgrSeq: 6487716,
					Sequence:          1,
				},
				Index: "00001A2969BE1FC85F1D7A55282FA2E6D95C71D2E4B9C0FDD3D9994F3C00FF8F",
			},
			{
				LedgerEntryType: ledger.OfferEntry,
				LedgerObject: &ledger.Offer{
					Account:           "rGryPmNWFognBgMtr9k4quqPbbEcCrhNmD",
					BookDirectory:     "71633D7DE1B6AEB32F87F1A73258B13FC8CC32942D53A66D4F038D7EA4C68000",
					BookNode:          "0000000000000000",
					LedgerEntryType:   ledger.OfferEntry,
					OwnerNode:         "0000000000000000",
					PreviousTxnID:     "555B93628BF3EC318892BB7C7CDCB6732FF53D12B6EEC4FAF60DD1AEE1C6101F",
					PreviousTxnLgrSeq: 3504261,
					Sequence:          3,
					TakerGets:         types.XRPCurrencyAmount(1000000),
					TakerPays: types.IssuedCurrencyAmount{
						Currency: "BTC",
						Issuer:   "rnuF96W4SZoCJmbHYBFoJZpR8eCaxNvekK",
						Value:    "1",
					},
				},
				Index: "000037C6659BB98F8D09F2F4CFEB27DE8EFEAFE54DD9E1C13AECDF5794B0C0F5",
			},
		},
		Marker: "0002A590029B53BE7857EFF9985F770EC792CE483720EB5E963C4D6A607D43DF",
	}
	j := `{
	"ledger_index": "6885842",
	"ledger_hash": "842B57C1CC0613299A686D3E9F310EC0422C84D3911E5056389AA7E5808A93C8",
	"state": [
		{
			"Account": "rKKzk9ghA2iuy3imqMXUHJqdRPMtNDGf4c",
			"Balance": "893730848",
			"Flags": 0,
			"LedgerEntryType": "AccountRoot",
			"OwnerCount": 0,
			"PreviousTxnID": "C204A65CF2542946289A3358C67D991B5E135FABFA89F271DBA7A150C08CA046",
			"PreviousTxnLgrSeq": 6487716,
			"Sequence": 1,
			"index": "00001A2969BE1FC85F1D7A55282FA2E6D95C71D2E4B9C0FDD3D9994F3C00FF8F"
		},
		{
			"Account": "rGryPmNWFognBgMtr9k4quqPbbEcCrhNmD",
			"BookDirectory": "71633D7DE1B6AEB32F87F1A73258B13FC8CC32942D53A66D4F038D7EA4C68000",
			"BookNode": "0000000000000000",
			"Flags": 0,
			"LedgerEntryType": "Offer",
			"OwnerNode": "0000000000000000",
			"PreviousTxnID": "555B93628BF3EC318892BB7C7CDCB6732FF53D12B6EEC4FAF60DD1AEE1C6101F",
			"PreviousTxnLgrSeq": 3504261,
			"Sequence": 3,
			"TakerGets": "1000000",
			"TakerPays": {
				"currency": "BTC",
				"issuer": "rnuF96W4SZoCJmbHYBFoJZpR8eCaxNvekK",
				"value": "1"
			},
			"index": "000037C6659BB98F8D09F2F4CFEB27DE8EFEAFE54DD9E1C13AECDF5794B0C0F5"
		}
	],
	"marker": "0002A590029B53BE7857EFF9985F770EC792CE483720EB5E963C4D6A607D43DF"
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

/*

,
		{
			"Balance": {
				"currency": "BTC",
				"issuer": "rrrrrrrrrrrrrrrrrrrrBZbvji",
				"value": "0"
			},
			"Flags": 131072,
			"HighLimit": {
				"currency": "BTC",
				"issuer": "rKUK9omZqVEnraCipKNFb5q4tuNTeqEDZS",
				"value": "10"
			},
			"HighNode": "0000000000000000",
			"LedgerEntryType": "RippleState",
			"LowLimit": {
				"currency": "BTC",
				"issuer": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
				"value": "0"
			},
			"LowNode": "0000000000000000",
			"PreviousTxnID": "87591A63051645F37B85D1FBA55EE69B1C96BFF16904F5C99F03FB93D42D0375",
			"PreviousTxnLgrSeq": 746872,
			"index": "000103996A3BAD918657F86E12A67D693E8FC8A814DA4B958A244B5F14D93E58"
		}

*/
