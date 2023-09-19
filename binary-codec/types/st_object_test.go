package types

import (
	"fmt"
	"testing"

	"github.com/CreatureDev/xrpl-go/binary-codec/definitions"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/stretchr/testify/require"
)

func TestCreateFieldInstanceMapFromJson(t *testing.T) {

	tt := []struct {
		description string
		input       map[string]any
		output      map[definitions.FieldInstance]any
		expectedErr error
	}{
		{
			description: "convert valid Json",
			input: map[string]any{
				"Fee":           types.XRPCurrencyAmount(10),
				"Flags":         uint(524288),
				"OfferSequence": uint(1752791),
				"TakerGets":     types.XRPCurrencyAmount(150000000000),
			},
			output: map[definitions.FieldInstance]any{
				getFieldInstance(t, "Fee"):           types.XRPCurrencyAmount(10),
				getFieldInstance(t, "Flags"):         uint(524288),
				getFieldInstance(t, "OfferSequence"): uint(1752791),
				getFieldInstance(t, "TakerGets"):     types.XRPCurrencyAmount(150000000000),
			},
			expectedErr: nil,
		},
		// {
		// 	description: "not found error",
		// 	input: &invalidTxWithBase{
		// 		InvalidField: "invalid",
		// 	},
		// 	output:      nil,
		// 	expectedErr: errors.New("FieldName InvalidField not found"),
		// },
		// {
		// 	description: "no base tx",
		// 	input: &invalidTx{
		// 		InvalidField: "invalid",
		// 	},
		// 	output:      nil,
		// 	expectedErr: errors.New("no base tx defined"),
		// },
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {

			got, err := createFieldInstanceMapFromJson(tc.input)
			if tc.expectedErr != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.EqualValues(t, tc.output, got)
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
			require.Equal(t, tc.output, getSortedKeys(tc.input))
		})
	}
}
