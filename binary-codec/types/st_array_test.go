package types

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/CreatureDev/xrpl-go/model/ledger"
	"github.com/CreatureDev/xrpl-go/model/transactions"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/stretchr/testify/require"
)

func TestSTArrayFromJson(t *testing.T) {
	tt := []struct {
		description string
		input       any
		output      string
		expectedErr error
	}{
		{
			description: "nested stobject test",
			input: []transactions.AffectedNode{
				{
					DeletedNode: &transactions.DeletedNode{
						FinalFields: &ledger.DirectoryNode{
							Flags:             types.SetFlag(0),
							RootIndex:         "036D7E923EF22B65E19D95A6365C3373E1E96586E27015074A0745621D069432",
							TakerGetsCurrency: "0000000000000000000000000000000000000000",
							TakerGetsIssuer:   "0000000000000000000000000000000000000000",
							TakerPaysCurrency: "0000000000000000000000004254430000000000",
							TakerPaysIssuer:   "06A148131B436B2561C85967685B098E050EED4E",
						},
						LedgerEntryType: ledger.DirectoryNodeEntry,
						LedgerIndex:     "036D7E923EF22B65E19D95A6365C3373E1E96586E27015074A0745621D069432",
					},
				},
				{
					ModifiedNode: &transactions.ModifiedNode{
						FinalFields: &ledger.DirectoryNode{
							Flags:     types.SetFlag(0),
							Owner:     "r68xfQkhFxZrbwo6RRKq728JF2fJYQRE1",
							RootIndex: "5ED0913938CD6D43BD6450201737394A9991753C4581E5682D61F35048D8FBFC",
						},
						LedgerEntryType: ledger.DirectoryNodeEntry,
						LedgerIndex:     "5ED0913938CD6D43BD6450201737394A9991753C4581E5682D61F35048D8FBFC",
					},
				},
				{
					DeletedNode: &transactions.DeletedNode{
						FinalFields: &ledger.Offer{
							Account:           "r68xfQkhFxZrbwo6RRKq728JF2fJYQRE1",
							BookDirectory:     "036D7E923EF22B65E19D95A6365C3373E1E96586E27015074A0745621D069432",
							BookNode:          "0",
							Flags:             131072,
							OwnerNode:         "0",
							PreviousTxnID:     "B21E9DC8DCB75AD12C4ACF4C72E6E822244CF6CFFD7BD738ACEED1264374B687",
							PreviousTxnLgrSeq: 82010985,
							Sequence:          59434311,
							TakerGets:         types.XRPCurrencyAmount(1166610661),
							TakerPays: types.IssuedCurrencyAmount{
								Currency: "BTC",
								Issuer:   "rchGBxcD1A1C2tdxF6papQYZ8kjRKMYcL",
								Value:    "0.023876",
							},
						},
						LedgerEntryType: ledger.OfferEntry,
						LedgerIndex:     "C171AA6003AE67B218E5EFFEF3E49CB5A4FE5A06D82C1F7EEF322B036EF76CE7",
					},
				},
				{
					ModifiedNode: &transactions.ModifiedNode{
						FinalFields: &ledger.AccountRoot{
							Account:    "r68xfQkhFxZrbwo6RRKq728JF2fJYQRE1",
							Balance:    types.XRPCurrencyAmount(3310960263),
							Flags:      types.SetFlag(0),
							MessageKey: "020000000000000000000000002B8F120A4CD1A6B236CC4389A30AB676CD722144",
							OwnerCount: 8,
							Sequence:   59434319,
						},
						LedgerEntryType: ledger.AccountRootEntry,
						LedgerIndex:     "F8C1F1AE7AEF06FEFB2311232CA30A7803820A79AE65F567354629D579BB38B1",
						PreviousFields: &ledger.AccountRoot{
							Balance:    types.XRPCurrencyAmount(3310960413),
							OwnerCount: 9,
							Sequence:   59434318,
						},
						PreviousTxnID:     "DC99E25A951DBBF89D331428DFFE72880159C03D0A4F48EBC81A277CFFBCE683",
						PreviousTxnLgrSeq: 82011654,
					},
				},
			},
			output:      "E411006456036D7E923EF22B65E19D95A6365C3373E1E96586E27015074A0745621D069432E7220000000058036D7E923EF22B65E19D95A6365C3373E1E96586E27015074A0745621D06943201110000000000000000000000004254430000000000021106A148131B436B2561C85967685B098E050EED4E0311000000000000000000000000000000000000000004110000000000000000000000000000000000000000E1E1E5110064565ED0913938CD6D43BD6450201737394A9991753C4581E5682D61F35048D8FBFCE72200000000585ED0913938CD6D43BD6450201737394A9991753C4581E5682D61F35048D8FBFC821407B6FEE98BBDD29F92FDFAF6107BC741C4130E8FE1E1E411006F56C171AA6003AE67B218E5EFFEF3E49CB5A4FE5A06D82C1F7EEF322B036EF76CE7E7220002000024038AE5472504E3636933000000000000000034000000000000000055B21E9DC8DCB75AD12C4ACF4C72E6E822244CF6CFFD7BD738ACEED1264374B6875010036D7E923EF22B65E19D95A6365C3373E1E96586E27015074A0745621D06943264D4087B8271DDA000000000000000000000000000425443000000000006A148131B436B2561C85967685B098E050EED4E6540000000458910E5811407B6FEE98BBDD29F92FDFAF6107BC741C4130E8FE1E1E51100612504E3660655DC99E25A951DBBF89D331428DFFE72880159C03D0A4F48EBC81A277CFFBCE68356F8C1F1AE7AEF06FEFB2311232CA30A7803820A79AE65F567354629D579BB38B1E624038AE54E2D000000096240000000C5593F1DE1E7220000000024038AE54F2D000000086240000000C5593E877221020000000000000000000000002B8F120A4CD1A6B236CC4389A30AB676CD722144811407B6FEE98BBDD29F92FDFAF6107BC741C4130E8FE1E1F1",
			expectedErr: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			sa := &STArray{}
			got, err := sa.FromJson(tc.input)
			if tc.expectedErr != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
				require.Empty(t, got)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.output, strings.ToUpper(hex.EncodeToString(got)))
			}
		})
	}

}

// func TestSTArrayToJson(t *testing.T) {
// 	tt := []struct {
// 		description string
// 		input       string
// 		output      map[string]any
// 		expectedErr error
// 	}{
// 		{
// 			description: "large starray",
// 			input:       "FAEC5A000913881D8CE533B1355B905B5C75F85B1F9A5149B94D7BABAF4475000011D9E1E010754368747470733A2F2F697066732E696F2F697066732F516D57356232715736744D3975615659376E555162774356334D63514B566E6D5659516A6A4368784E65656B476BE1F1",
// 			output: map[string]any{
// 				"NFTokens": []any{
// 					map[string]any{
// 						"NFToken": map[string]any{
// 							"NFTokenID": "000913881D8CE533B1355B905B5C75F85B1F9A5149B94D7BABAF4475000011D9",
// 						},
// 					},
// 					map[string]any{
// 						"Signer": map[string]any{
// 							"URI": "68747470733A2F2F697066732E696F2F697066732F516D57356232715736744D3975615659376E555162774356334D63514B566E6D5659516A6A4368784E65656B476B",
// 						},
// 					},
// 				},
// 			},
// 			expectedErr: nil,
// 		},
// 		{
// 			description: "simple starray",
// 			input:       "F9EA364A0745621D0694327D0F04C4D46544659A2D58525043686174E1E9364A0745621D0694327D0F04C4D46544659A2D58525043686174E1F1",
// 			output: map[string]any{
// 				"Memos": []any{
// 					map[string]any{
// 						"Memo": map[string]any{
// 							"MemoData":     "04C4D46544659A2D58525043686174",
// 							"ExchangeRate": "4A0745621D069432",
// 						},
// 					},
// 					map[string]any{
// 						"TemplateEntry": map[string]any{
// 							"MemoData":     "04C4D46544659A2D58525043686174",
// 							"ExchangeRate": "4A0745621D069432",
// 						},
// 					},
// 				},
// 			},
// 		},
// 		{
// 			description: "smaller stobject test",
// 			input:       "E51100612504E3660655DC99E25A951DBBF89D331428DFFE72880159C03D0A4F48EBC81A277CFFBCE68356F8C1F1AE7AEF06FEFB2311232CA30A7803820A79AE65F567354629D579BB38B1E624038AE54E2D000000096240000000C5593F1DE1E7220000000024038AE54F2D000000086240000000C5593E877221020000000000000000000000002B8F120A4CD1A6B236CC4389A30AB676CD722144811407B6FEE98BBDD29F92FDFAF6107BC741C4130E8FE1E1",
// 			output: map[string]any{
// 				"ModifiedNode": map[string]any{
// 					"FinalFields": map[string]any{
// 						"Account":    "r68xfQkhFxZrbwo6RRKq728JF2fJYQRE1",
// 						"Balance":    "3310960263",
// 						"Flags":      0,
// 						"MessageKey": "020000000000000000000000002B8F120A4CD1A6B236CC4389A30AB676CD722144",
// 						"OwnerCount": 8,
// 						"Sequence":   59434319,
// 					},
// 					"LedgerEntryType": "AccountRoot",
// 					"LedgerIndex":     "F8C1F1AE7AEF06FEFB2311232CA30A7803820A79AE65F567354629D579BB38B1",
// 					"PreviousFields": map[string]any{
// 						"Balance":    "3310960413",
// 						"OwnerCount": 9,
// 						"Sequence":   59434318,
// 					},
// 					"PreviousTxnID":     "DC99E25A951DBBF89D331428DFFE72880159C03D0A4F48EBC81A277CFFBCE683",
// 					"PreviousTxnLgrSeq": 82011654,
// 				},
// 			},
// 		},
// 	}

// 	for _, tc := range tt {
// 		t.Run(tc.description, func(t *testing.T) {
// 			p := serdes.NewBinaryParser([]byte(tc.input))
// 			sa := &STArray{}
// 			act, err := sa.ToJson(p)
// 			if tc.expectedErr != nil {
// 				require.Error(t, err, tc.expectedErr.Error())
// 				require.Nil(t, act)
// 			} else {
// 				require.NoError(t, err)
// 				require.EqualValues(t, tc.output, act)
// 			}
// 		})
// 	}
// }
