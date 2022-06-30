//go:build unit
// +build unit

package binarycodec

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyield/xrpl-go/binary-codec/definitions"
)

func TestCreateFieldInstanceMapFromJson(t *testing.T) {

	tt := []struct {
		description string
		input       map[string]interface{}
		output      map[definitions.FieldInstance]interface{}
		expectedErr error
	}{
		{
			description: "convert valid Json",
			input: map[string]interface{}{
				"Fee":           "10",
				"Flags":         524288,
				"OfferSequence": 1752791,
				"TakerGets":     "150000000000",
			},
			output: map[definitions.FieldInstance]interface{}{
				getFieldInstance(t, "Fee"):           "10",
				getFieldInstance(t, "Flags"):         524288,
				getFieldInstance(t, "OfferSequence"): 1752791,
				getFieldInstance(t, "TakerGets"):     "150000000000",
			},
			expectedErr: nil,
		},
		{
			description: "not found error",
			input: map[string]interface{}{
				"IncorrectField": 89,
				"Flags":          525288,
				"OfferSequence":  1752791,
			},
			output:      nil,
			expectedErr: errors.New("FieldName IncorrectField not found"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {

			got, err := createFieldInstanceMapFromJson(tc.input)
			if tc.expectedErr != nil {
				assert.EqualError(t, err, tc.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.output, got)
			}
		})
	}

}

func getFieldInstance(t *testing.T, fieldName string) definitions.FieldInstance {
	t.Helper()
	fi, err := definitions.Get().GetFieldInstanceByFieldName(fieldName)
	if err != nil {
		t.Fatalf("FieldInstance with FieldName %v", fieldName)
	}
	return *fi
}

func TestGetSortedKeys(t *testing.T) {
	tt := []struct {
		input  map[definitions.FieldInstance]interface{}
		output []definitions.FieldInstance
	}{
		{
			input: map[definitions.FieldInstance]interface{}{
				getFieldInstance(t, "IndexNext"):       5100000,
				getFieldInstance(t, "SourceTag"):       1232,
				getFieldInstance(t, "LedgerEntryType"): 1,
			},
			output: []definitions.FieldInstance{
				getFieldInstance(t, "LedgerEntryType"),
				getFieldInstance(t, "SourceTag"),
				getFieldInstance(t, "IndexNext"),
			},
		},
		{
			input: map[definitions.FieldInstance]interface{}{
				getFieldInstance(t, "Account"):      "rMBzp8CgpE441cp5PVyA9rpVV7oT8hP3ys",
				getFieldInstance(t, "TransferRate"): 4234,
				getFieldInstance(t, "Expiration"):   23,
			},
			output: []definitions.FieldInstance{
				getFieldInstance(t, "Expiration"),
				getFieldInstance(t, "TransferRate"),
				getFieldInstance(t, "Account"),
			},
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test %v", i), func(t *testing.T) {
			assert.Equal(t, tc.output, getSortedKeys(tc.input))
		})
	}
}
func TestEncode(t *testing.T) {
	tt := []struct {
		description string
		input       map[string]any
		output      string
		expectedErr error
	}{
		// {
		// 	description: "successfullty serialized signed transaction",
		// 	input: `{
		// 		"Account": "rMBzp8CgpE441cp5PVyA9rpVV7oT8hP3ys",
		// 		"Expiration": 595640108,
		// 		"Fee": "10",
		// 		"Flags": 524288,
		// 		"OfferSequence": 1752791,
		// 		"Sequence": 1752792,
		// 		"SigningPubKey": "03EE83BB432547885C219634A1BC407A9DB0474145D69737D09CCDC63E1DEE7FE3",
		// 		"TakerGets": "15000000000",
		// 		"TakerPays": {
		// 		  "currency": "USD",
		// 		  "issuer": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
		// 		  "value": "7072.8"
		// 		},
		// 		"TransactionType": "OfferCreate",
		// 		"TxnSignature": "30440220143759437C04F7B61F012563AFE90D8DAFC46E86035E1D965A9CED282C97D4CE02204CFD241E86F17E011298FC1A39B63386C74306A5DE047E213B0F29EFA4571C2C",
		// 		"hash": "73734B611DDA23D3F5F62E20A173B78AB8406AC5015094DA53F53D39B9EDB06C"
		// 	  }"`,
		// 	output:      "120007220008000024001ABED82A2380BF2C2019001ABED764D55920AC9391400000000000000000000000000055534400000000000A20B3C85F482532A9578DBB3950B85CA06594D165400000037E11D60068400000000000000A732103EE83BB432547885C219634A1BC407A9DB0474145D69737D09CCDC63E1DEE7FE3744630440220143759437C04F7B61F012563AFE90D8DAFC46E86035E1D965A9CED282C97D4CE02204CFD241E86F17E011298FC1A39B63386C74306A5DE047E213B0F29EFA4571C2C8114DD76483FACDEE26E60D8A586BB58D09F27045C46",
		// 	expectedErr: nil,
		// },
		{
			description: "test Flags",
			input:       map[string]any{"Flags": 524288},
			output:      "2200080000",
			expectedErr: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			got, err := Encode(tc.input)

			if tc.expectedErr != nil {
				assert.EqualError(t, err, tc.expectedErr.Error())
				assert.Empty(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.output, got)
			}
		})
	}

}
