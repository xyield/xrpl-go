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
