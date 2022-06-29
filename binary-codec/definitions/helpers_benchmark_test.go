package definitions

import (
	"fmt"
	"testing"
)

func BenchmarkGetTypeNameByFieldName(b *testing.B) {

	tt := []struct {
		input string
	}{
		{
			input: "Generic",
		},
		{
			input: "yurt",
		},
	}

	for _, test := range tt {
		b.Run(fmt.Sprintf("input_name_%v", test.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				definitions.GetTypeNameByFieldName(test.input)
			}
		})
	}
}

func BenchmarkGetTypeCodeByTypeName(b *testing.B) {

	tt := []struct {
		input string
	}{
		{
			input: "Validation",
		},
		{
			input: "yurt",
		},
	}

	for _, test := range tt {
		b.Run(fmt.Sprintf("input_name_%v", test.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				definitions.GetTypeCodeByTypeName(test.input)
			}
		})
	}
}

func BenchmarkGetTypeCodeByFieldName(b *testing.B) {
	tt := []struct {
		input string
	}{
		{
			input: "Generic",
		},
		{
			input: "yurt",
		},
	}

	for _, test := range tt {
		b.Run(fmt.Sprintf("input_name_%v", test.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				definitions.GetTypeCodeByFieldName(test.input)
			}
		})
	}
}

func BenchmarkGetFieldCodeByFieldName(b *testing.B) {
	tt := []struct {
		input string
	}{
		{
			input: "Generic",
		},
		{
			input: "yurt",
		},
	}

	for _, test := range tt {
		b.Run(fmt.Sprintf("input_name_%v", test.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				definitions.GetFieldCodeByFieldName(test.input)
			}
		})
	}
}

func BenchmarkGetFieldHeaderByFieldName(b *testing.B) {
	tt := []struct {
		input string
	}{
		{
			input: "Generic",
		},
		{
			input: "yurt",
		},
	}

	for _, test := range tt {
		b.Run(fmt.Sprintf("input_name_%v", test.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				definitions.GetFieldHeaderByFieldName(test.input)
			}
		})
	}
}

// func BenchmarkGetFieldNameByFieldHeader(b *testing.B) {
// 	tt := []struct {
// 		input fieldHeader
// 	}{
// 		{
// 			input: fieldHeader{
// 				TypeCode:  []byte{1},
// 				FieldCode: []byte{1},
// 			},
// 		},
// 		{
// 			input: fieldHeader{
// 				TypeCode: []byte() 0000000000111,
// 				FieldCode: 00000000000000111,
// 			},
// 		},
// 	}

// 	for _, test := range tt {
// 		b.Run(fmt.Sprintf("input_name_%v", test.input), func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				definitions.GetFieldNameByFieldHeader(test.input)
// 			}
// 		})
// 	}
// }

func BenchmarkGetFieldInfoByFieldName(b *testing.B) {
	tt := []struct {
		input string
	}{
		{
			input: "Generic",
		},
		{
			input: "yurt",
		},
	}

	for _, test := range tt {
		b.Run(fmt.Sprintf("input_name_%v", test.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				definitions.GetFieldInfoByFieldName(test.input)
			}
		})
	}
}

func BenchmarkGetFieldInstanceByFieldName(b *testing.B) {
	tt := []struct {
		input string
	}{
		{
			input: "Generic",
		},
		{
			input: "yurt",
		},
	}

	for _, test := range tt {
		b.Run(fmt.Sprintf("input_name_%v", test.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				definitions.GetFieldInstanceByFieldName(test.input)
			}
		})
	}
}

func BenchmarkGetTransactionTypeCodeByTransactionTypeName(b *testing.B) {
	tt := []struct {
		input string
	}{
		{
			input: "Payment",
		},
		{
			input: "yurt",
		},
	}

	for _, test := range tt {
		b.Run(fmt.Sprintf("input_name_%v", test.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				definitions.GetTransactionTypeCodeByTransactionTypeName(test.input)
			}
		})
	}
}

func BenchmarkGetTransactionTypeNameByTransactionTypeCode(b *testing.B) {
	tt := []struct {
		input int32
	}{
		{
			input: 1,
		},
		{
			input: 999999999,
		},
	}

	for _, test := range tt {
		b.Run(fmt.Sprintf("input_code_%v", test.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				definitions.GetTransactionTypeNameByTransactionTypeCode(test.input)
			}
		})
	}
}

func BenchmarkGetTransactionResultNameByTransactionResultTypeCode(b *testing.B) {
	tt := []struct {
		input int32
	}{
		{
			input: 100,
		},
		{
			input: 999999999,
		},
	}

	for _, test := range tt {
		b.Run(fmt.Sprintf("input_code_%v", test.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				definitions.GetTransactionResultNameByTransactionResultTypeCode(test.input)
			}
		})
	}
}

func BenchmarkGetTransactionResultTypeCodeByTransactionResultName(b *testing.B) {
	tt := []struct {
		input string
	}{
		{
			input: "tesSUCCESS",
		},
		{
			input: "yurt",
		},
	}

	for _, test := range tt {
		b.Run(fmt.Sprintf("input_name_%v", test.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				definitions.GetTransactionResultTypeCodeByTransactionResultName(test.input)
			}
		})
	}
}

func BenchmarkGetLedgerEntryTypeCodeByLedgerEntryTypeName(b *testing.B) {
	tt := []struct {
		input string
	}{
		{
			input: "Any",
		},
		{
			input: "yurt",
		},
	}

	for _, test := range tt {
		b.Run(fmt.Sprintf("input_name_%v", test.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				definitions.GetLedgerEntryTypeCodeByLedgerEntryTypeName(test.input)
			}
		})
	}
}

func BenchmarkGetLedgerEntryTypeNameByLedgerEntryTypeCode(b *testing.B) {
	tt := []struct {
		input int32
	}{
		{
			input: 100,
		},
		{
			input: 999999999,
		},
	}

	for _, test := range tt {
		b.Run(fmt.Sprintf("input_code_%v", test.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				definitions.GetLedgerEntryTypeNameByLedgerEntryTypeCode(test.input)
			}
		})
	}
}
