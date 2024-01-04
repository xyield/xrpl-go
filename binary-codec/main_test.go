package binarycodec

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/model/transactions"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/stretchr/testify/require"
)

func TestEncode(t *testing.T) {
	tt := []struct {
		description string
		input       transactions.Tx
		output      string
		expectedErr error
	}{
		{
			description: "successfully serialized signed transaction 1",
			input: &transactions.OfferCreate{
				BaseTx: transactions.BaseTx{
					Account:         "rMBzp8CgpE441cp5PVyA9rpVV7oT8hP3ys",
					TransactionType: transactions.OfferCreateTx,
					Fee:             10,
					Flags:           types.SetFlag(524288),
					Sequence:        1752792,
					SigningPubKey:   "03EE83BB432547885C219634A1BC407A9DB0474145D69737D09CCDC63E1DEE7FE3",
					TxnSignature:    "30440220143759437C04F7B61F012563AFE90D8DAFC46E86035E1D965A9CED282C97D4CE02204CFD241E86F17E011298FC1A39B63386C74306A5DE047E213B0F29EFA4571C2C",
				},
				Expiration:    595640108,
				OfferSequence: 1752791,
				TakerGets:     types.XRPCurrencyAmount(15000000000),
				TakerPays: types.IssuedCurrencyAmount{
					Currency: "USD",
					Issuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
					Value:    "7072.8",
				},
			},
			output:      "120007220008000024001ABED82A2380BF2C2019001ABED764D55920AC9391400000000000000000000000000055534400000000000A20B3C85F482532A9578DBB3950B85CA06594D165400000037E11D60068400000000000000A732103EE83BB432547885C219634A1BC407A9DB0474145D69737D09CCDC63E1DEE7FE3744630440220143759437C04F7B61F012563AFE90D8DAFC46E86035E1D965A9CED282C97D4CE02204CFD241E86F17E011298FC1A39B63386C74306A5DE047E213B0F29EFA4571C2C8114DD76483FACDEE26E60D8A586BB58D09F27045C46",
			expectedErr: nil,
		},
		{
			description: "successfully serialized signed transaction 2",
			input: &transactions.EscrowFinish{
				BaseTx: transactions.BaseTx{
					Account:         "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA",
					TransactionType: transactions.EscrowFinishTx,
					Fee:             10101,
					Flags:           types.SetFlag(2147483648),
					Sequence:        1,
					SigningPubKey:   "0268D79CD579D077750740FA18A2370B7C2018B2714ECE70BA65C38D223E79BC9C",
					TxnSignature:    "3045022100F06FB54049D6D50142E5CF2E2AC21946AF305A13E2A2D4BA881B36484DD01A540220311557EC8BEF536D729605A4CB4D4DC51B1E37C06C93434DD5B7651E1E2E28BF",
					Memos: []transactions.MemoWrapper{
						{
							Memo: transactions.Memo{
								MemoData: "04C4D46544659A2D58525043686174",
							},
						},
					},
				},
				OfferSequence: 11,
				Owner:         "r9NpyVfLfUG8hatuCCHKzosyDtKnBdsEN3",
			},
			output:      "1200022280000000240000000120190000000B68400000000000277573210268D79CD579D077750740FA18A2370B7C2018B2714ECE70BA65C38D223E79BC9C74473045022100F06FB54049D6D50142E5CF2E2AC21946AF305A13E2A2D4BA881B36484DD01A540220311557EC8BEF536D729605A4CB4D4DC51B1E37C06C93434DD5B7651E1E2E28BF811452C7F01AD13B3CA9C1D133FA8F3482D2EF08FA7D82145A380FBD236B6A1CD14B939AD21101E5B6B6FFA2F9EA7D0F04C4D46544659A2D58525043686174E1F1",
			expectedErr: nil,
		},
		{
			description: "successfully serialized signed transaction 3",
			input: &transactions.Payment{
				BaseTx: transactions.BaseTx{
					Account:         "rweYz56rfmQ98cAdRaeTxQS9wVMGnrdsFp",
					TransactionType: transactions.PaymentTx,
					Fee:             12,
					Flags:           types.SetFlag(0),
					Sequence:        842,
					Memos: []transactions.MemoWrapper{
						{
							Memo: transactions.Memo{
								MemoData: "7274312E312E31",
								MemoType: "636C69656E74",
							},
						},
					},
					LastLedgerSequence: 9902014,
					SigningPubKey:      "0379F17CFA0FFD7518181594BE69FE9A10471D6DE1F4055C6D2746AFD6CF89889E",
					TxnSignature:       "3045022100D55ED1953F860ADC1BC5CD993ABB927F48156ACA31C64737865F4F4FF6D015A80220630704D2BD09C8E99F26090C25F11B28F5D96A1350454402C2CED92B39FFDBAF",
				},
				Amount:      types.XRPCurrencyAmount(10000000),
				Destination: "rweYz56rfmQ98cAdRaeTxQS9wVMGnrdsFp",
				SendMax: types.IssuedCurrencyAmount{
					Currency: "USD",
					Issuer:   "rweYz56rfmQ98cAdRaeTxQS9wVMGnrdsFp",
					Value:    "0.6275558355",
				},
				Paths: [][]transactions.PathStep{
					{
						{
							Account: "rPDXxSZcuVL3ZWoyU82bcde3zwvmShkRyF",
						},
						{
							Currency: "XRP",
						},
					},
					{
						{
							Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
						},
						{
							Account: "rMwjYedjc7qqtKYVLiAccJSmCwih4LnE2q",
						},
						{
							Currency: "XRP",
						},
					},
				},
			},
			expectedErr: nil,
			output:      "1200002200000000240000034A201B009717BE61400000000098968068400000000000000C69D4564B964A845AC0000000000000000000000000555344000000000069D33B18D53385F8A3185516C2EDA5DEDB8AC5C673210379F17CFA0FFD7518181594BE69FE9A10471D6DE1F4055C6D2746AFD6CF89889E74473045022100D55ED1953F860ADC1BC5CD993ABB927F48156ACA31C64737865F4F4FF6D015A80220630704D2BD09C8E99F26090C25F11B28F5D96A1350454402C2CED92B39FFDBAF811469D33B18D53385F8A3185516C2EDA5DEDB8AC5C6831469D33B18D53385F8A3185516C2EDA5DEDB8AC5C6F9EA7C06636C69656E747D077274312E312E31E1F1011201F3B1997562FD742B54D4EBDEA1D6AEA3D4906B8F100000000000000000000000000000000000000000FF014B4E9C06F24296074F7BC48F92A97916C6DC5EA901DD39C650A96EDA48334E70CC4A85B8B2E8502CD310000000000000000000000000000000000000000000",
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			got, err := Encode(tc.input)

			if tc.expectedErr != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
				require.Empty(t, got)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.output, got)
			}
		})
	}

}

func TestDecode(t *testing.T) {
	tt := []struct {
		description string
		input       string
		output      map[string]any
		expectedErr error
	}{
		{
			description: "zero issued currency amount",
			output: map[string]any{
				"LowLimit": map[string]any{
					"currency": "LUC",
					"issuer":   "rsygE5ynt2iSasscfCCeqaGBGiFKMCAUu7",
					"value":    "0",
				},
			},
			input:       "6680000000000000000000000000000000000000004C5543000000000020A85019EA62B48F79EB67273B797EB916438FA4",
			expectedErr: nil,
		},
		{
			description: "decode tx1",
			input:       "120007220008000024001ABED82A2380BF2C2019001ABED764D55920AC9391400000000000000000000000000055534400000000000A20B3C85F482532A9578DBB3950B85CA06594D165400000037E11D60068400000000000000A732103EE83BB432547885C219634A1BC407A9DB0474145D69737D09CCDC63E1DEE7FE3744630440220143759437C04F7B61F012563AFE90D8DAFC46E86035E1D965A9CED282C97D4CE02204CFD241E86F17E011298FC1A39B63386C74306A5DE047E213B0F29EFA4571C2C8114DD76483FACDEE26E60D8A586BB58D09F27045C46",
			output: map[string]any{
				"Account":       "rMBzp8CgpE441cp5PVyA9rpVV7oT8hP3ys",
				"Expiration":    595640108,
				"Fee":           "10",
				"Flags":         524288,
				"OfferSequence": 1752791,
				"Sequence":      1752792,
				"SigningPubKey": "03EE83BB432547885C219634A1BC407A9DB0474145D69737D09CCDC63E1DEE7FE3",
				"TakerGets":     "15000000000",
				"TakerPays": map[string]any{
					"currency": "USD",
					"issuer":   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
					"value":    "7072.8",
				},
				"TransactionType": "OfferCreate",
				"TxnSignature":    "30440220143759437C04F7B61F012563AFE90D8DAFC46E86035E1D965A9CED282C97D4CE02204CFD241E86F17E011298FC1A39B63386C74306A5DE047E213B0F29EFA4571C2C",
			},
			expectedErr: nil,
		},
		{
			description: "deserialize Uint64 correctly",
			input:       "34000000044B82FA09",
			output:      map[string]any{"OwnerNode": "000000044B82FA09"},
			expectedErr: nil,
		},
		{
			description: "deserialize Uint16 LedgerEntryType",
			input:       "110072",
			output:      map[string]any{"LedgerEntryType": "RippleState"},
			expectedErr: nil,
		},
		{
			description: "deserialize Uint16 TransferFee",
			input:       "14789A",
			output:      map[string]any{"TransferFee": 30874},
			expectedErr: nil,
		},
		{
			description: "deserialize Uint8 int correctly",
			input:       "011019",
			output:      map[string]any{"CloseResolution": 25},
			expectedErr: nil,
		},
		{
			description: "deserialize Vector256 successfully,",
			input:       "03134073734B611DDA23D3F5F62E20A173B78AB8406AC5015094DA53F53D39B9EDB06C73734B611DDA23D3F5F62E20A173B78AB8406AC5015094DA53F53D39B9EDB06C",
			output:      map[string]any{"Amendments": []string{"73734B611DDA23D3F5F62E20A173B78AB8406AC5015094DA53F53D39B9EDB06C", "73734B611DDA23D3F5F62E20A173B78AB8406AC5015094DA53F53D39B9EDB06C"}},
			expectedErr: nil,
		},
		{
			description: "deserialize hash 128",
			input:       "4173734B611DDA23D3F5F62E20A173B78A",
			output:      map[string]any{"EmailHash": "73734B611DDA23D3F5F62E20A173B78A"},
			expectedErr: nil,
		},
		{
			description: "deserialize hash 160",
			input:       "011173734B611DDA23D3F5F62E20A173B78AB8406AC5",
			output:      map[string]any{"TakerPaysCurrency": "73734B611DDA23D3F5F62E20A173B78AB8406AC5"},
			expectedErr: nil,
		},
		{
			description: "deserialize hash 256",
			input:       "501573734B611DDA23D3F5F62E20A173B78AB8406AC5015094DA53F53D39B9EDB06C",
			output:      map[string]any{"Digest": "73734B611DDA23D3F5F62E20A173B78AB8406AC5015094DA53F53D39B9EDB06C"},
			expectedErr: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			act, err := Decode(tc.input)
			if tc.expectedErr != nil {
				require.Error(t, err, tc.expectedErr.Error())
				require.Nil(t, act)
			} else {
				require.NoError(t, err)
				require.EqualValues(t, tc.output, act)
			}
		})
	}

}

func TestEncodeForMultisigning(t *testing.T) {
	tt := []struct {
		description string
		input       transactions.Tx
		accountID   string
		output      string
		expectedErr error
	}{
		{
			description: "serialize tx1 for signing correctly",
			input: &transactions.OfferCreate{
				BaseTx: transactions.BaseTx{
					Account:         "rMBzp8CgpE441cp5PVyA9rpVV7oT8hP3ys",
					TransactionType: transactions.OfferCreateTx,
					Fee:             10,
					Flags:           types.SetFlag(524288),
					Sequence:        1752792,
					SigningPubKey:   "03EE83BB432547885C219634A1BC407A9DB0474145D69737D09CCDC63E1DEE7FE3",
					TxnSignature:    "30440220143759437C04F7B61F012563AFE90D8DAFC46E86035E1D965A9CED282C97D4CE02204CFD241E86F17E011298FC1A39B63386C74306A5DE047E213B0F29EFA4571C2C",
				},
				Expiration:    595640108,
				OfferSequence: 1752791,
				TakerGets:     types.XRPCurrencyAmount(15000000000),
				TakerPays: types.IssuedCurrencyAmount{
					Currency: "USD",
					Issuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
					Value:    "7072.8",
				},
			},
			accountID:   "rMBzp8CgpE441cp5PVyA9rpVV7oT8hP3ys",
			output:      "534D5400120007220008000024001ABED82A2380BF2C2019001ABED764D55920AC9391400000000000000000000000000055534400000000000A20B3C85F482532A9578DBB3950B85CA06594D165400000037E11D60068400000000000000A73008114DD76483FACDEE26E60D8A586BB58D09F27045C46DD76483FACDEE26E60D8A586BB58D09F27045C46",
			expectedErr: nil,
		},
		{
			description: "SigningPubKey is not present",
			input: &transactions.OfferCreate{
				BaseTx: transactions.BaseTx{
					Account:         "rMBzp8CgpE441cp5PVyA9rpVV7oT8hP3ys",
					TransactionType: transactions.OfferCreateTx,
					Fee:             10,
					Flags:           types.SetFlag(524288),
					Sequence:        1752792,
					TxnSignature:    "30440220143759437C04F7B61F012563AFE90D8DAFC46E86035E1D965A9CED282C97D4CE02204CFD241E86F17E011298FC1A39B63386C74306A5DE047E213B0F29EFA4571C2C",
				},
				Expiration:    595640108,
				OfferSequence: 1752791,
				TakerGets:     types.XRPCurrencyAmount(15000000000),
				TakerPays: types.IssuedCurrencyAmount{
					Currency: "USD",
					Issuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
					Value:    "7072.8",
				},
			},
			accountID:   "rMBzp8CgpE441cp5PVyA9rpVV7oT8hP3ys",
			output:      "534D5400120007220008000024001ABED82A2380BF2C2019001ABED764D55920AC9391400000000000000000000000000055534400000000000A20B3C85F482532A9578DBB3950B85CA06594D165400000037E11D60068400000000000000A73008114DD76483FACDEE26E60D8A586BB58D09F27045C46DD76483FACDEE26E60D8A586BB58D09F27045C46",
			expectedErr: nil,
		},
		{
			description: "SigningPubKey empty string",
			input: &transactions.OfferCreate{
				BaseTx: transactions.BaseTx{
					Account:         "rMBzp8CgpE441cp5PVyA9rpVV7oT8hP3ys",
					TransactionType: transactions.OfferCreateTx,
					Fee:             10,
					Flags:           types.SetFlag(524288),
					Sequence:        1752792,
					SigningPubKey:   "",
					TxnSignature:    "30440220143759437C04F7B61F012563AFE90D8DAFC46E86035E1D965A9CED282C97D4CE02204CFD241E86F17E011298FC1A39B63386C74306A5DE047E213B0F29EFA4571C2C",
				},
				Expiration:    595640108,
				OfferSequence: 1752791,
				TakerGets:     types.XRPCurrencyAmount(15000000000),
				TakerPays: types.IssuedCurrencyAmount{
					Currency: "USD",
					Issuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
					Value:    "7072.8",
				},
			},
			accountID:   "rMBzp8CgpE441cp5PVyA9rpVV7oT8hP3ys",
			output:      "534D5400120007220008000024001ABED82A2380BF2C2019001ABED764D55920AC9391400000000000000000000000000055534400000000000A20B3C85F482532A9578DBB3950B85CA06594D165400000037E11D60068400000000000000A73008114DD76483FACDEE26E60D8A586BB58D09F27045C46DD76483FACDEE26E60D8A586BB58D09F27045C46",
			expectedErr: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			got, err := EncodeForMultisigning(tc.input, tc.accountID)

			if tc.expectedErr != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
				require.Empty(t, got)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.output, got)
			}
		})
	}
}

func TestEncodeForSigningClaim(t *testing.T) {

	tt := []struct {
		description string
		input       transactions.PaymentChannelClaim
		output      string
		expectedErr error
	}{
		{
			description: "successfully encode claim",
			input: transactions.PaymentChannelClaim{
				Channel: "43904CBFCDCEC530B4037871F86EE90BF799DF8D2E0EA564BC8A3F332E4F5FB1",
				Amount:  1000,
			},
			output:      "434C4D0043904CBFCDCEC530B4037871F86EE90BF799DF8D2E0EA564BC8A3F332E4F5FB100000000000003E8",
			expectedErr: nil,
		},
		{
			description: "fail to encode claim - no channel",
			input: transactions.PaymentChannelClaim{
				Amount: 1000,
			},
			output:      "",
			expectedErr: ErrSigningClaimFieldNotFound,
		},
		{
			description: "fail to encode claim - no amount",
			input: transactions.PaymentChannelClaim{
				Channel: "43904CBFCDCEC530B4037871F86EE90BF799DF8D2E0EA564BC8A3F332E4F5FB1",
			},
			output:      "",
			expectedErr: ErrSigningClaimFieldNotFound,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			got, err := EncodeForSigningClaim(tc.input)

			if tc.expectedErr != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
				require.Empty(t, got)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.output, got)
			}
		})
	}
}

func TestEncodeForSigning(t *testing.T) {
	tt := []struct {
		description string
		input       transactions.Tx
		output      string
		expectedErr error
	}{
		{
			description: "serialize tx1 for signing correctly",
			input: &transactions.OfferCreate{
				BaseTx: transactions.BaseTx{
					Account:         "rMBzp8CgpE441cp5PVyA9rpVV7oT8hP3ys",
					TransactionType: transactions.OfferCreateTx,
					Fee:             10,
					Flags:           types.SetFlag(524288),
					Sequence:        1752792,
					SigningPubKey:   "03EE83BB432547885C219634A1BC407A9DB0474145D69737D09CCDC63E1DEE7FE3",
					TxnSignature:    "30440220143759437C04F7B61F012563AFE90D8DAFC46E86035E1D965A9CED282C97D4CE02204CFD241E86F17E011298FC1A39B63386C74306A5DE047E213B0F29EFA4571C2C",
				},
				Expiration:    595640108,
				OfferSequence: 1752791,
				TakerGets:     types.XRPCurrencyAmount(15000000000),
				TakerPays: types.IssuedCurrencyAmount{
					Currency: "USD",
					Issuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
					Value:    "7072.8",
				},
			},
			output:      "53545800120007220008000024001ABED82A2380BF2C2019001ABED764D55920AC9391400000000000000000000000000055534400000000000A20B3C85F482532A9578DBB3950B85CA06594D165400000037E11D60068400000000000000A732103EE83BB432547885C219634A1BC407A9DB0474145D69737D09CCDC63E1DEE7FE38114DD76483FACDEE26E60D8A586BB58D09F27045C46",
			expectedErr: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			got, err := EncodeForSigning(tc.input)

			if tc.expectedErr != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
				require.Empty(t, got)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.output, got)
			}
		})
	}
}
