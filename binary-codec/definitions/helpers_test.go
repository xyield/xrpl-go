package definitions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTypeNameByFieldName(t *testing.T) {

	tt := []struct {
		description   string
		input         string
		expected      string
		expectedError error
	}{
		{
			description:   "test that `TransferRate` gives `UInt32`",
			input:         "TransferRate",
			expected:      "UInt32",
			expectedError: nil,
		},
		{
			description: "test that invalid value gives an error",
			input:       "yurt",
			expected:    "",
			expectedError: &NotFoundError{
				Instance: "FieldName",
				Input:    "yurt",
			},
		},
	}

	for _, test := range tt {

		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetTypeNameByFieldName(test.input)
			if test.expectedError != nil {
				assert.Error(t, test.expectedError, err.Error())
				assert.Zero(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})

	}

}

// func BenchmarkGetTypeNameByFieldName(b *testing.B) {

// 	tt := []struct {
// 		input string
// 	}{
// 		{
// 			input: "Generic",
// 		},
// 		{
// 			input: "yurt",
// 		},
// 	}

// 	for _, test := range tt {
// 		b.Run(fmt.Sprintf("input_name_%v", test.input), func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				definitions.GetTypeNameByFieldName(test.input)
// 			}
// 		})
// 	}
// }

func TestGetTypeCodeByTypeName(t *testing.T) {

	tt := []struct {
		description   string
		input         string
		expected      int
		expectedError error
	}{
		{
			description:   "test that `Done` gives correct code",
			input:         "Done",
			expected:      -1,
			expectedError: nil,
		},
		{
			description:   "test that `Hash128` gives correct code",
			input:         "Hash128",
			expected:      4,
			expectedError: nil,
		},
		{
			description: "test that incorrect value gives an error",
			input:       "yurt",
			expected:    0,
			expectedError: &NotFoundError{
				Instance: "TypeName",
				Input:    "yurt",
			},
		},
	}

	for _, test := range tt {

		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetTypeCodeByTypeName(test.input)
			if test.expectedError != nil {
				assert.Error(t, test.expectedError, err.Error())
				assert.Zero(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})

	}

}

// func BenchmarkGetTypeCodeByTypeName(b *testing.B) {

// 	tt := []struct {
// 		input string
// 	}{
// 		{
// 			input: "Validation",
// 		},
// 		{
// 			input: "yurt",
// 		},
// 	}

// 	for _, test := range tt {
// 		b.Run(fmt.Sprintf("input_name_%v", test.input), func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				definitions.GetTypeCodeByTypeName(test.input)
// 			}
// 		})
// 	}
// }

func TestGetTypeCodeByFieldName(t *testing.T) {
	tt := []struct {
		description   string
		input         string
		expected      int
		expectedError error
	}{
		{
			description:   "test that `TransferRate` gives 2",
			input:         "TransferRate",
			expected:      2,
			expectedError: nil,
		},
		{
			description:   "test that `OwnerNode` gives 3",
			input:         "OwnerNode",
			expected:      3,
			expectedError: nil,
		},
		{
			description: "test that non-existent value gives error",
			input:       "yurt",
			expected:    0,
			expectedError: &NotFoundError{
				Instance: "FieldName",
				Input:    "yurt",
			},
		},
	}

	for _, test := range tt {

		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetTypeCodeByFieldName(test.input)
			if test.expectedError != nil {
				assert.Error(t, test.expectedError, err.Error())
				assert.Zero(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})
	}
}

// func BenchmarkGetTypeCodeByFieldName(b *testing.B) {
// 	tt := []struct {
// 		input string
// 	}{
// 		{
// 			input: "Generic",
// 		},
// 		{
// 			input: "yurt",
// 		},
// 	}

// 	for _, test := range tt {
// 		b.Run(fmt.Sprintf("input_name_%v", test.input), func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				definitions.GetTypeCodeByFieldName(test.input)
// 			}
// 		})
// 	}
// }
func TestGetFieldCodeByFieldName(t *testing.T) {

	tt := []struct {
		description   string
		input         string
		expected      int
		expectedError error
	}{
		{
			description:   "correct FieldCode",
			input:         "TransferRate",
			expected:      11,
			expectedError: nil,
		},
		{
			description: "Invalid FieldName",
			input:       "yurt",
			expected:    0,
			expectedError: &NotFoundError{
				Instance: "FieldName",
				Input:    "yurt",
			},
		},
	}

	for _, test := range tt {

		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetFieldCodeByFieldName(test.input)
			if test.expectedError != nil {
				assert.Error(t, test.expectedError, err.Error())
				assert.Zero(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})
	}
}

// func BenchmarkGetFieldCodeByFieldName(b *testing.B) {
// 	tt := []struct {
// 		input string
// 	}{
// 		{
// 			input: "Generic",
// 		},
// 		{
// 			input: "yurt",
// 		},
// 	}

// 	for _, test := range tt {
// 		b.Run(fmt.Sprintf("input_name_%v", test.input), func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				definitions.GetFieldCodeByFieldName(test.input)
// 			}
// 		})
// 	}
// }
func TestGetFieldHeaderByFieldName(t *testing.T) {
	tt := []struct {
		description   string
		input         string
		expected      *fieldHeader
		expectedError error
	}{
		{
			description: "correct FieldHeader",
			input:       "TransferRate",
			expected: &fieldHeader{
				TypeCode:  2,
				FieldCode: 11,
			},
			expectedError: nil,
		},
		{
			description: "Invalid FieldName",
			input:       "yurt",
			expected:    nil,
			expectedError: &NotFoundError{
				Instance: "FieldName",
				Input:    "yurt",
			},
		},
	}

	for _, test := range tt {

		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetFieldHeaderByFieldName(test.input)
			if test.expectedError != nil {
				assert.Error(t, test.expectedError, err.Error())
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})
	}
}

// func BenchmarkGetFieldHeaderByFieldName(b *testing.B) {
// 	tt := []struct {
// 		input string
// 	}{
// 		{
// 			input: "Generic",
// 		},
// 		{
// 			input: "yurt",
// 		},
// 	}

// 	for _, test := range tt {
// 		b.Run(fmt.Sprintf("input_name_%v", test.input), func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				definitions.GetFieldHeaderByFieldName(test.input)
// 			}
// 		})
// 	}
// }
func TestGetFieldInfoByFieldName(t *testing.T) {
	tt := []struct {
		description   string
		input         string
		expected      *fieldInfo
		expectedError error
	}{
		{
			description: "correct FieldInfo",
			input:       "TransferRate",
			expected: &fieldInfo{
				Nth:            11,
				IsVLEncoded:    false,
				IsSerialized:   true,
				IsSigningField: true,
				Type:           "UInt32",
			},
			expectedError: nil,
		},
		{
			description: "invalid FieldInfo",
			input:       "yurt",
			expected:    nil,
			expectedError: &NotFoundError{
				Instance: "FieldName",
				Input:    "yurt",
			},
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetFieldInfoByFieldName(test.input)
			if test.expectedError != nil {
				assert.Error(t, test.expectedError, err.Error())
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}

		})
	}
}

// func BenchmarkGetFieldInfoByFieldName(b *testing.B) {
// 	tt := []struct {
// 		input string
// 	}{
// 		{
// 			input: "Generic",
// 		},
// 		{
// 			input: "yurt",
// 		},
// 	}

// 	for _, test := range tt {
// 		b.Run(fmt.Sprintf("input_name_%v", test.input), func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				definitions.GetFieldInfoByFieldName(test.input)
// 			}
// 		})
// 	}
// }
func TestGetFieldInstanceByFieldName(t *testing.T) {
	tt := []struct {
		description   string
		input         string
		expected      *fieldInstance
		expectedError error
	}{
		{
			description: "correct FieldInstance",
			input:       "TransferRate",
			expected: &fieldInstance{
				FieldName: "TransferRate",
				fieldInfo: &fieldInfo{
					Nth:            11,
					IsVLEncoded:    false,
					IsSerialized:   true,
					IsSigningField: true,
					Type:           "UInt32",
				},
				FieldHeader: &fieldHeader{
					TypeCode:  2,
					FieldCode: 11,
				},
			},
		},
		{
			description: "invalid FieldName",
			input:       "yurt",
			expected:    nil,
			expectedError: &NotFoundError{
				Instance: "FieldName",
				Input:    "yurt",
			},
		},
	}

	for _, test := range tt {

		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetFieldInstanceByFieldName(test.input)
			if test.expectedError != nil {
				assert.Error(t, test.expectedError, err.Error())
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})
	}
}

// func BenchmarkGetFieldInstanceByFieldName(b *testing.B) {
// 	tt := []struct {
// 		input string
// 	}{
// 		{
// 			input: "Generic",
// 		},
// 		{
// 			input: "yurt",
// 		},
// 	}

// 	for _, test := range tt {
// 		b.Run(fmt.Sprintf("input_name_%v", test.input), func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				definitions.GetFieldInstanceByFieldName(test.input)
// 			}
// 		})
// 	}
// }
func TestGetTransactionTypeCodeByTransactionTypeName(t *testing.T) {
	tt := []struct {
		description   string
		input         string
		expected      int
		expectedError error
	}{
		{
			description:   "correct TransactionTypeCode",
			input:         "EscrowCreate",
			expected:      1,
			expectedError: nil,
		},
		{
			description: "invalid TransactionTypeName",
			input:       "yurt",
			expected:    0,
			expectedError: &NotFoundError{
				Instance: "TransactionTypeName",
				Input:    "yurt",
			},
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetTransactionTypeCodeByTransactionTypeName(test.input)
			if test.expectedError != nil {
				assert.Error(t, test.expectedError, err.Error())
				assert.Zero(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})
	}
}

// func BenchmarkGetTransactionTypeCodeByTransactionTypeName(b *testing.B) {
// 	tt := []struct {
// 		input string
// 	}{
// 		{
// 			input: "Payment",
// 		},
// 		{
// 			input: "yurt",
// 		},
// 	}

// 	for _, test := range tt {
// 		b.Run(fmt.Sprintf("input_name_%v", test.input), func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				definitions.GetTransactionTypeCodeByTransactionTypeName(test.input)
// 			}
// 		})
// 	}
// }
func TestGetTransactionTypeNameByTransactionTypeCode(t *testing.T) {
	tt := []struct {
		description   string
		input         int
		expected      string
		expectedError error
	}{
		{
			description:   "correct TypeName",
			input:         1,
			expected:      "EscrowCreate",
			expectedError: nil,
		},
		{
			description: "invalid TransactionTypeCode",
			input:       999999999999,
			expected:    "",
			expectedError: &NotFoundErrorInt{
				Instance: "TransactionTypeCode",
				Input:    999999999999,
			},
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetTransactionTypeNameByTransactionTypeCode(test.input)
			if test.expectedError != nil {
				assert.Error(t, test.expectedError, err.Error())
				assert.Zero(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})
	}
}

// func BenchmarkGetTransactionTypeNameByTransactionTypeCode(b *testing.B) {
// 	tt := []struct {
// 		input int
// 	}{
// 		{
// 			input: 1,
// 		},
// 		{
// 			input: 999999999999999999,
// 		},
// 	}

// 	for _, test := range tt {
// 		b.Run(fmt.Sprintf("input_code_%v", test.input), func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				definitions.GetTransactionTypeNameByTransactionTypeCode(test.input)
// 			}
// 		})
// 	}
// }
func TestGetTransactionResultNameByTransactionResultTypeCode(t *testing.T) {
	tt := []struct {
		description   string
		input         int
		expected      string
		expectedError error
	}{
		{
			description:   "correct TransactionResultName",
			input:         100,
			expected:      "tecCLAIM",
			expectedError: nil,
		},
		{
			description: "invalid txResultTypeCode",
			input:       999999999999999,
			expected:    "",
			expectedError: &NotFoundErrorInt{
				Instance: "TransactionResultTypeCode",
				Input:    999999999999999,
			},
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetTransactionResultNameByTransactionResultTypeCode(test.input)
			if test.expectedError != nil {
				assert.Error(t, test.expectedError, err.Error())
				assert.Zero(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})
	}
}

// func BenchmarkGetTransactionResultNameByTransactionResultTypeCode(b *testing.B) {
// 	tt := []struct {
// 		input int
// 	}{
// 		{
// 			input: 100,
// 		},
// 		{
// 			input: 999999999999999999,
// 		},
// 	}

// 	for _, test := range tt {
// 		b.Run(fmt.Sprintf("input_code_%v", test.input), func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				definitions.GetTransactionResultNameByTransactionResultTypeCode(test.input)
// 			}
// 		})
// 	}
// }
func TestGetTransactionResultTypeCodeByTransactionResultName(t *testing.T) {
	tt := []struct {
		description   string
		input         string
		expected      int
		expectedError error
	}{
		{
			description:   "correct TransactionResultTypeCode",
			input:         "tecCLAIM",
			expected:      100,
			expectedError: nil,
		},
		{
			description: "invalid TransactionResultName",
			input:       "yurt",
			expected:    0,
			expectedError: &NotFoundError{
				Instance: "TransactionResultName",
				Input:    "yurt",
			},
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetTransactionResultTypeCodeByTransactionResultName(test.input)
			if test.expectedError != nil {
				assert.Error(t, test.expectedError, err.Error())
				assert.Zero(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})
	}
}

// func BenchmarkGetTransactionResultTypeCodeByTransactionResultName(b *testing.B) {
// 	tt := []struct {
// 		input string
// 	}{
// 		{
// 			input: "tesSUCCESS",
// 		},
// 		{
// 			input: "yurt",
// 		},
// 	}

// 	for _, test := range tt {
// 		b.Run(fmt.Sprintf("input_name_%v", test.input), func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				definitions.GetTransactionResultTypeCodeByTransactionResultName(test.input)
// 			}
// 		})
// 	}
// }
func TestGetLedgerEntryTypeCodeByLedgerEntryTypeName(t *testing.T) {
	tt := []struct {
		description   string
		input         string
		expected      int
		expectedError error
	}{
		{
			description:   "correct LedgerEntryTypeCode",
			input:         "Any",
			expected:      -3,
			expectedError: nil,
		},
		{
			description: "invalid LedgerEntryTypeName",
			input:       "yurt",
			expected:    0,
			expectedError: &NotFoundError{
				Instance: "LedgerEntryTypeName",
				Input:    "yurt",
			},
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetLedgerEntryTypeCodeByLedgerEntryTypeName(test.input)
			if test.expectedError != nil {
				assert.Error(t, test.expectedError, err.Error())
				assert.Zero(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})
	}

}

// func BenchmarkGetLedgerEntryTypeCodeByLedgerEntryTypeName(b *testing.B) {
// 	tt := []struct {
// 		input string
// 	}{
// 		{
// 			input: "Any",
// 		},
// 		{
// 			input: "yurt",
// 		},
// 	}

// 	for _, test := range tt {
// 		b.Run(fmt.Sprintf("input_name_%v", test.input), func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				definitions.GetLedgerEntryTypeCodeByLedgerEntryTypeName(test.input)
// 			}
// 		})
// 	}
// }
func TestGetLedgerEntryTypeNameByLedgerEntryTypeCode(t *testing.T) {
	tt := []struct {
		description   string
		input         int
		expected      string
		expectedError error
	}{
		{
			description:   "correct LedgerEntryTypeName",
			input:         -3,
			expected:      "Any",
			expectedError: nil,
		},
		{
			description: "invalid LedgerEntryTypeCode",
			input:       999999999999,
			expected:    "",
			expectedError: &NotFoundErrorInt{
				Instance: "LedgerEntryTypeCode",
				Input:    999999999999,
			},
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetLedgerEntryTypeNameByLedgerEntryTypeCode(test.input)
			if test.expectedError != nil {
				assert.Error(t, test.expectedError, err.Error())
				assert.Zero(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})
	}
}

// func BenchmarkGetLedgerEntryTypeNameByLedgerEntryTypeCode(b *testing.B) {
// 	tt := []struct {
// 		input int
// 	}{
// 		{
// 			input: 100,
// 		},
// 		{
// 			input: 999999999999999999,
// 		},
// 	}

// 	for _, test := range tt {
// 		b.Run(fmt.Sprintf("input_code_%v", test.input), func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				definitions.GetLedgerEntryTypeNameByLedgerEntryTypeCode(test.input)
// 			}
// 		})
// 	}
// }
