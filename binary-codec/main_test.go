package binarycodec

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xyield/xrpl-go/binary-codec/types"
)

// Binary serializations of valid transactions json
var (
	Tx1                    = "120007220008000024001ABED82A2380BF2C2019001ABED764D55920AC9391400000000000000000000000000055534400000000000A20B3C85F482532A9578DBB3950B85CA06594D165400000037E11D60068400000000000000A732103EE83BB432547885C219634A1BC407A9DB0474145D69737D09CCDC63E1DEE7FE3744630440220143759437C04F7B61F012563AFE90D8DAFC46E86035E1D965A9CED282C97D4CE02204CFD241E86F17E011298FC1A39B63386C74306A5DE047E213B0F29EFA4571C2C8114DD76483FACDEE26E60D8A586BB58D09F27045C46"
	Tx2                    = "1200022280000000240000000120190000000B68400000000000277573210268D79CD579D077750740FA18A2370B7C2018B2714ECE70BA65C38D223E79BC9C74473045022100F06FB54049D6D50142E5CF2E2AC21946AF305A13E2A2D4BA881B36484DD01A540220311557EC8BEF536D729605A4CB4D4DC51B1E37C06C93434DD5B7651E1E2E28BF811452C7F01AD13B3CA9C1D133FA8F3482D2EF08FA7D82145A380FBD236B6A1CD14B939AD21101E5B6B6FFA2F9EA7D0F04C4D46544659A2D58525043686174E1F1"
	Tx3                    = "1200002200000000240000034A201B009717BE61400000000098968068400000000000000C69D4564B964A845AC0000000000000000000000000555344000000000069D33B18D53385F8A3185516C2EDA5DEDB8AC5C673210379F17CFA0FFD7518181594BE69FE9A10471D6DE1F4055C6D2746AFD6CF89889E74473045022100D55ED1953F860ADC1BC5CD993ABB927F48156ACA31C64737865F4F4FF6D015A80220630704D2BD09C8E99F26090C25F11B28F5D96A1350454402C2CED92B39FFDBAF811469D33B18D53385F8A3185516C2EDA5DEDB8AC5C6831469D33B18D53385F8A3185516C2EDA5DEDB8AC5C6F9EA7C06636C69656E747D077274312E312E31E1F1011201F3B1997562FD742B54D4EBDEA1D6AEA3D4906B8F100000000000000000000000000000000000000000FF014B4E9C06F24296074F7BC48F92A97916C6DC5EA901DD39C650A96EDA48334E70CC4A85B8B2E8502CD310000000000000000000000000000000000000000000"
	LedgerEntryTypeExample = "110072"             // generated from the xrpl js project
	UInt64TypeExample      = "34000000044B82FA09" // generated from the xrpl js project
	UInt8IntExample        = "011019"             // generated from the xrpl js project
)

func TestEncode(t *testing.T) {
	tt := []struct {
		description string
		fromTx      string
		input       map[string]any
		output      string
		expectedErr error
	}{
		{
			description: "successfully serialized signed transaction 1",
			input: map[string]any{
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
				"hash":            "73734B611DDA23D3F5F62E20A173B78AB8406AC5015094DA53F53D39B9EDB06C",
			},
			output:      "120007220008000024001ABED82A2380BF2C2019001ABED764D55920AC9391400000000000000000000000000055534400000000000A20B3C85F482532A9578DBB3950B85CA06594D165400000037E11D60068400000000000000A732103EE83BB432547885C219634A1BC407A9DB0474145D69737D09CCDC63E1DEE7FE3744630440220143759437C04F7B61F012563AFE90D8DAFC46E86035E1D965A9CED282C97D4CE02204CFD241E86F17E011298FC1A39B63386C74306A5DE047E213B0F29EFA4571C2C8114DD76483FACDEE26E60D8A586BB58D09F27045C46",
			expectedErr: nil,
		},
		{
			description: "successfully serialized signed transaction 2",
			input: map[string]any{
				"TransactionType": "EscrowFinish",
				"Flags":           2147483648,
				"Sequence":        1,
				"OfferSequence":   11,
				"Fee":             "10101",
				"SigningPubKey":   "0268D79CD579D077750740FA18A2370B7C2018B2714ECE70BA65C38D223E79BC9C",
				"TxnSignature":    "3045022100F06FB54049D6D50142E5CF2E2AC21946AF305A13E2A2D4BA881B36484DD01A540220311557EC8BEF536D729605A4CB4D4DC51B1E37C06C93434DD5B7651E1E2E28BF",
				"Account":         "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA",
				"Owner":           "r9NpyVfLfUG8hatuCCHKzosyDtKnBdsEN3",
				"Memos": []any{
					map[string]any{
						"Memo": map[string]any{
							"MemoData": "04C4D46544659A2D58525043686174",
						},
					},
				},
			},
			output:      "1200022280000000240000000120190000000B68400000000000277573210268D79CD579D077750740FA18A2370B7C2018B2714ECE70BA65C38D223E79BC9C74473045022100F06FB54049D6D50142E5CF2E2AC21946AF305A13E2A2D4BA881B36484DD01A540220311557EC8BEF536D729605A4CB4D4DC51B1E37C06C93434DD5B7651E1E2E28BF811452C7F01AD13B3CA9C1D133FA8F3482D2EF08FA7D82145A380FBD236B6A1CD14B939AD21101E5B6B6FFA2F9EA7D0F04C4D46544659A2D58525043686174E1F1",
			expectedErr: nil,
		},
		{
			description: "successfully serialized signed transaction 3",
			input: map[string]any{
				"Account":            "rweYz56rfmQ98cAdRaeTxQS9wVMGnrdsFp",
				"Amount":             "10000000",
				"Destination":        "rweYz56rfmQ98cAdRaeTxQS9wVMGnrdsFp",
				"Fee":                "12",
				"Flags":              0,
				"LastLedgerSequence": 9902014,
				"Memos": []any{
					map[string]any{
						"Memo": map[string]any{
							"MemoData": "7274312E312E31",
							"MemoType": "636C69656E74",
						},
					},
				},
				"Paths": []any{
					[]any{
						map[string]any{
							"account":  "rPDXxSZcuVL3ZWoyU82bcde3zwvmShkRyF",
							"type":     1,
							"type_hex": "0000000000000001",
						},
						map[string]any{
							"currency": "XRP",
							"type":     16,
							"type_hex": "0000000000000010",
						},
					},
					[]any{
						map[string]any{
							"account":  "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
							"type":     1,
							"type_hex": "0000000000000001",
						},
						map[string]any{
							"account":  "rMwjYedjc7qqtKYVLiAccJSmCwih4LnE2q",
							"type":     1,
							"type_hex": "0000000000000001",
						},
						map[string]any{
							"currency": "XRP",
							"type":     16,
							"type_hex": "0000000000000010",
						},
					},
				},
				"SendMax": map[string]any{
					"currency": "USD",
					"issuer":   "rweYz56rfmQ98cAdRaeTxQS9wVMGnrdsFp",
					"value":    "0.6275558355",
				},
				"Sequence":        842,
				"SigningPubKey":   "0379F17CFA0FFD7518181594BE69FE9A10471D6DE1F4055C6D2746AFD6CF89889E",
				"TransactionType": "Payment",
				"TxnSignature":    "3045022100D55ED1953F860ADC1BC5CD993ABB927F48156ACA31C64737865F4F4FF6D015A80220630704D2BD09C8E99F26090C25F11B28F5D96A1350454402C2CED92B39FFDBAF",
				"hash":            "B521424226FC100A2A802FE20476A5F8426FD3F720176DC5CCCE0D75738CC208",
			},
			expectedErr: nil,
			output:      "1200002200000000240000034A201B009717BE61400000000098968068400000000000000C69D4564B964A845AC0000000000000000000000000555344000000000069D33B18D53385F8A3185516C2EDA5DEDB8AC5C673210379F17CFA0FFD7518181594BE69FE9A10471D6DE1F4055C6D2746AFD6CF89889E74473045022100D55ED1953F860ADC1BC5CD993ABB927F48156ACA31C64737865F4F4FF6D015A80220630704D2BD09C8E99F26090C25F11B28F5D96A1350454402C2CED92B39FFDBAF811469D33B18D53385F8A3185516C2EDA5DEDB8AC5C6831469D33B18D53385F8A3185516C2EDA5DEDB8AC5C6F9EA7C06636C69656E747D077274312E312E31E1F1011201F3B1997562FD742B54D4EBDEA1D6AEA3D4906B8F100000000000000000000000000000000000000000FF014B4E9C06F24296074F7BC48F92A97916C6DC5EA901DD39C650A96EDA48334E70CC4A85B8B2E8502CD310000000000000000000000000000000000000000000",
		},
		{ // output correct from js encode lib
			description: "serialize Destination example - AccountID",
			fromTx:      "",
			input:       map[string]any{"Destination": "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA"},
			output:      "831452C7F01AD13B3CA9C1D133FA8F3482D2EF08FA7D",
			expectedErr: nil,
		},
		{ // output correct from js encode lib
			description: "serialize Issuer example - AccountID",
			fromTx:      "",
			input:       map[string]any{"Issuer": "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA"},
			output:      "841452C7F01AD13B3CA9C1D133FA8F3482D2EF08FA7D",
			expectedErr: nil,
		},
		{ // output correct from js encode lib
			description: "serialize Authorize example - AccountID",
			fromTx:      "",
			input:       map[string]any{"Authorize": "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA"},
			output:      "851452C7F01AD13B3CA9C1D133FA8F3482D2EF08FA7D",
			expectedErr: nil,
		},
		{ // output correct from js encode lib
			description: "serialize Unauthorize example - AccountID",
			fromTx:      "",
			input:       map[string]any{"Unauthorize": "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA"},
			output:      "861452C7F01AD13B3CA9C1D133FA8F3482D2EF08FA7D",
			expectedErr: nil,
		},
		{ // output correct from js encode lib
			description: "serialize Target example - AccountID",
			fromTx:      "",
			input:       map[string]any{"Target": "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA"},
			output:      "871452C7F01AD13B3CA9C1D133FA8F3482D2EF08FA7D",
			expectedErr: nil,
		},
		{ // output correct from js encode lib
			description: "serialize NFTokenMinter example - AccountID",
			fromTx:      "",
			input:       map[string]any{"NFTokenMinter": "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA"},
			output:      "891452C7F01AD13B3CA9C1D133FA8F3482D2EF08FA7D",
			expectedErr: nil,
		},
		{
			description: "serialize OwnderNode example - UInt64",
			fromTx:      UInt64TypeExample,
			input:       map[string]any{"OwnerNode": "18446744073"},
			output:      "34000000044B82FA09",
			expectedErr: nil,
		},
		{
			description: "serialize LedgerEntryType example - UInt8",
			fromTx:      LedgerEntryTypeExample,
			input:       map[string]any{"LedgerEntryType": "RippleState"},
			output:      "110072",
			expectedErr: nil,
		},
		{
			description: "serialize int example - UInt8",
			fromTx:      UInt8IntExample,
			input:       map[string]any{"CloseResolution": 25},
			output:      "011019",
			expectedErr: nil,
		},
		{
			description: "serialize multiple fields out of sequence to check ordering for successfully signed tx 1",
			fromTx:      Tx1,
			input: map[string]any{
				"Flags":           524288,
				"OfferSequence":   1752791,
				"TransactionType": "OfferCreate",
				"Expiration":      595640108,
				"Sequence":        1752792,
			},
			output:      "120007220008000024001ABED82A2380BF2C2019001ABED7",
			expectedErr: nil,
		},
		{
			description: "serialize TransactionType from successfully signed tx 1",
			fromTx:      Tx1,
			input:       map[string]any{"TransactionType": "OfferCreate"},
			output:      "120007",
			expectedErr: nil,
		},
		{
			description: "serialize Flags from successfully signed tx 1",
			fromTx:      Tx1,
			input:       map[string]any{"Flags": 524288},
			output:      "2200080000",
			expectedErr: nil,
		},
		{
			description: "serialize Sequence from successfully signed tx 1",
			fromTx:      Tx1,
			input:       map[string]any{"Sequence": 1752792},
			output:      "24001ABED8",
			expectedErr: nil,
		},
		{
			description: "serialize Expiration from successfully signed tx 1",
			fromTx:      Tx1,
			input:       map[string]any{"Expiration": 595640108},
			output:      "2A2380BF2C",
			expectedErr: nil,
		},
		{
			description: "serialize OfferSequence from successfully signed tx 1",
			fromTx:      Tx1,
			input:       map[string]any{"OfferSequence": 1752791},
			output:      "2019001ABED7",
			expectedErr: nil,
		},
		{
			description: "serialize hash 128",
			input:       map[string]any{"EmailHash": "73734B611DDA23D3F5F62E20A173B78A"},
			output:      "4173734B611DDA23D3F5F62E20A173B78A",
			expectedErr: nil,
		},
		{
			description: "hash128 wrong length",
			input:       map[string]any{"EmailHash": "73734B611DDA23D3F5F62E20A173"},
			output:      "",
			expectedErr: &types.ErrInvalidHashLength{Expected: 16},
		},
		{
			description: "serialize hash 160",
			input:       map[string]any{"TakerPaysCurrency": "73734B611DDA23D3F5F62E20A173B78AB8406AC5"},
			output:      "011173734B611DDA23D3F5F62E20A173B78AB8406AC5",
			expectedErr: nil,
		},
		{
			description: "hash160 wrong length",
			input:       map[string]any{"TakerPaysCurrency": "73734B611DDA23D3F5F62E20A173B789"},
			output:      "",
			expectedErr: &types.ErrInvalidHashLength{Expected: 20},
		},
		{ // hash output doesn't appear in the txjson serialized binary output
			description: "serialize hash 256",
			input:       map[string]any{"Digest": "73734B611DDA23D3F5F62E20A173B78AB8406AC5015094DA53F53D39B9EDB06C"},
			output:      "501573734B611DDA23D3F5F62E20A173B78AB8406AC5015094DA53F53D39B9EDB06C",
			expectedErr: nil,
		},
		{ // hash output doesn't appear in the txjson serialized binary output
			description: "hash256 wrong length",
			input:       map[string]any{"Digest": "73734B611DDA23D3F5F62E20A173B78AB8406AC5015094DA53F537"},
			output:      "",
			expectedErr: &types.ErrInvalidHashLength{Expected: 32},
		},
		{ // output correct from js encode lib
			description: "serialize TakerPays from successfully signed tx 1",
			fromTx:      Tx1,
			input: map[string]any{"TakerPays": map[string]any{
				"currency": "USD",
				"issuer":   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
				"value":    "7072.8",
			},
			},
			output:      "64D55920AC9391400000000000000000000000000055534400000000000A20B3C85F482532A9578DBB3950B85CA06594D1",
			expectedErr: nil,
		},
		{ // output correct from js encode lib
			description: "serialize TakerGets from successfully signed tx 1",
			fromTx:      Tx1,
			input:       map[string]any{"TakerGets": "15000000000"},
			output:      "65400000037E11D600",
			expectedErr: nil,
		},
		{ // output correct from js encode lib
			description: "serialize Fee from successfully signed tx 1",
			fromTx:      Tx1,
			input:       map[string]any{"Fee": "10"},
			output:      "68400000000000000A",
			expectedErr: nil,
		},
		{ // output correct from js encode lib
			description: "serialize SigningPubKey from successfully signed tx 1",
			fromTx:      Tx1,
			input:       map[string]any{"SigningPubKey": "03EE83BB432547885C219634A1BC407A9DB0474145D69737D09CCDC63E1DEE7FE3"},
			output:      "732103EE83BB432547885C219634A1BC407A9DB0474145D69737D09CCDC63E1DEE7FE3",
			expectedErr: nil,
		},
		{ // output correct from js encode lib
			description: "serialize TxnSignature from successfully signed tx 1",
			fromTx:      Tx1,
			input:       map[string]any{"TxnSignature": "30440220143759437C04F7B61F012563AFE90D8DAFC46E86035E1D965A9CED282C97D4CE02204CFD241E86F17E011298FC1A39B63386C74306A5DE047E213B0F29EFA4571C2C"},
			output:      "744630440220143759437C04F7B61F012563AFE90D8DAFC46E86035E1D965A9CED282C97D4CE02204CFD241E86F17E011298FC1A39B63386C74306A5DE047E213B0F29EFA4571C2C",
			expectedErr: nil,
		},
		{ // output correct from js encode lib
			description: "serialize Account from successfully signed tx 1",
			fromTx:      Tx1,
			input:       map[string]any{"Account": "rMBzp8CgpE441cp5PVyA9rpVV7oT8hP3ys"},
			output:      "8114DD76483FACDEE26E60D8A586BB58D09F27045C46",
			expectedErr: nil,
		},
		{
			description: "serialize TransactionType from successfully signed tx 2",
			fromTx:      Tx2,
			input:       map[string]any{"TransactionType": "EscrowFinish"},
			output:      "120002",
			expectedErr: nil,
		},
		{
			description: "serialize Flags from successfully signed tx 2",
			fromTx:      Tx2,
			input:       map[string]any{"Flags": 2147483648},
			output:      "2280000000",
			expectedErr: nil,
		},
		{
			description: "serialize Sequence from successfully signed tx 2",
			fromTx:      Tx2,
			input:       map[string]any{"Sequence": 1},
			output:      "2400000001",
			expectedErr: nil,
		},
		{
			description: "serialize OfferSequence from successfully signed tx 2",
			fromTx:      Tx2,
			input:       map[string]any{"OfferSequence": 11},
			output:      "20190000000B",
			expectedErr: nil,
		},
		{ // output correct from js encode lib
			description: "serialize Account from successfully signed tx 2 - AccountID",
			fromTx:      Tx2,
			input:       map[string]any{"Account": "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA"},
			output:      "811452C7F01AD13B3CA9C1D133FA8F3482D2EF08FA7D",
			expectedErr: nil,
		},
		{ // output correct from js encode lib
			description: "serialize Owner from successfully signed tx 2 - AccountID",
			fromTx:      Tx2,
			input:       map[string]any{"Owner": "r9NpyVfLfUG8hatuCCHKzosyDtKnBdsEN3"},
			output:      "82145A380FBD236B6A1CD14B939AD21101E5B6B6FFA2",
			expectedErr: nil,
		},
		{
			description: "serialize TransactionType from successfully signed tx 3",
			fromTx:      Tx3,
			input:       map[string]any{"TransactionType": "Payment"},
			output:      "120000",
			expectedErr: nil,
		},
		{
			description: "serialize Flags from successfully signed tx 3",
			fromTx:      Tx3,
			input:       map[string]any{"Flags": 0},
			output:      "2200000000",
			expectedErr: nil,
		},
		{
			description: "serialize Sequence from successfully signed tx 3",
			fromTx:      Tx3,
			input:       map[string]any{"Sequence": 842},
			output:      "240000034A",
			expectedErr: nil,
		},
		{
			description: "serialize LastLedgerSequence from successfully signed tx 3",
			fromTx:      Tx3,
			input:       map[string]any{"LastLedgerSequence": 9902014},
			output:      "201B009717BE",
			expectedErr: nil,
		},
		{ // output correct from js encode lib
			description: "serialize Account from successfully signed tx 3",
			fromTx:      Tx3,
			input:       map[string]any{"Account": "rweYz56rfmQ98cAdRaeTxQS9wVMGnrdsFp"},
			output:      "811469D33B18D53385F8A3185516C2EDA5DEDB8AC5C6",
			expectedErr: nil,
		},
		{
			description: "serialize Vector256 successfully,",
			input:       map[string]any{"Amendments": []string{"73734B611DDA23D3F5F62E20A173B78AB8406AC5015094DA53F53D39B9EDB06C", "73734B611DDA23D3F5F62E20A173B78AB8406AC5015094DA53F53D39B9EDB06C"}},
			output:      "03134073734B611DDA23D3F5F62E20A173B78AB8406AC5015094DA53F53D39B9EDB06C73734B611DDA23D3F5F62E20A173B78AB8406AC5015094DA53F53D39B9EDB06C",
			expectedErr: nil,
		},
		{
			description: "invalid input for Vector256 - not a string array",
			input:       map[string]any{"Amendments": []int{1, 2, 3}},
			output:      "",
			expectedErr: &types.ErrInvalidVector256Type{Got: "[]int"},
		},
		{
			description: "invalid input for Vector256 - wrong hash length",
			input:       map[string]any{"Amendments": []string{"73734B611DDA23D3F5F62E20A173B78AB8406AC5015094DA53F53D39B9EDB06C56342689", "73734B611DDA23D3F5F62E20A173B78AB8406AC5015094DA53F53D39B9EDB06"}},
			output:      "",
			expectedErr: &types.ErrInvalidHashLength{Expected: types.HashLengthBytes},
		},
		{
			description: "serialize STObject correctly",
			input: map[string]any{
				"Memo": map[string]any{
					"MemoType": "04C4D46544659A2D58525043686174",
				},
			},
			output:      "EA7C0F04C4D46544659A2D58525043686174E1",
			expectedErr: nil,
		},
		{
			description: "serialize Paths correctly from Tx3",
			fromTx:      Tx3,
			input: map[string]any{
				"Paths": []any{
					[]any{
						map[string]any{
							"account":  "rPDXxSZcuVL3ZWoyU82bcde3zwvmShkRyF",
							"type":     1,
							"type_hex": "0000000000000001",
						},
						map[string]any{
							"currency": "XRP",
							"type":     16,
							"type_hex": "0000000000000010",
						},
					},
					[]any{
						map[string]any{
							"account":  "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
							"type":     1,
							"type_hex": "0000000000000001",
						},
						map[string]any{
							"account":  "rMwjYedjc7qqtKYVLiAccJSmCwih4LnE2q",
							"type":     1,
							"type_hex": "0000000000000001",
						},
						map[string]any{
							"currency": "XRP",
							"type":     16,
							"type_hex": "0000000000000010",
						},
					},
				},
			},
			output:      "011201F3B1997562FD742B54D4EBDEA1D6AEA3D4906B8F100000000000000000000000000000000000000000FF014B4E9C06F24296074F7BC48F92A97916C6DC5EA901DD39C650A96EDA48334E70CC4A85B8B2E8502CD310000000000000000000000000000000000000000000",
			expectedErr: nil,
		},
		{
			description: "invalid pathset",
			input: map[string]any{"Paths": []any{
				map[string]any{
					"account":  "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
					"currency": "USD",
					"issuer":   "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA",
				},
				map[string]any{
					"account":  "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
					"currency": "USD",
					"issuer":   "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA",
				},
			},
			},
			output:      "",
			expectedErr: types.ErrInvalidPathSet{},
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

			// checks if serialized elements from example transactions Json are present in full transaction binary result
			if tc.fromTx != "" {
				switch tc.fromTx {
				case Tx1:
					assert.Contains(t, Tx1, got)
				case Tx2:
					assert.Contains(t, Tx2, got)
				case Tx3:
					assert.Contains(t, Tx3, got)
				case LedgerEntryTypeExample:
					assert.Contains(t, LedgerEntryTypeExample, got)
				case UInt64TypeExample:
					assert.Contains(t, UInt64TypeExample, got)
				default:
					assert.Contains(t, UInt8IntExample, got)
				}
			}
		})
	}

}
