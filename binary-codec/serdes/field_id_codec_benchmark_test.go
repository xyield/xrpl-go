package serdes

import (
	"encoding/hex"
	"fmt"
	"testing"
)

// nolint
func BenchmarkEncode(b *testing.B) {

	tt := []struct {
		input string
	}{
		{
			input: "LedgerEntry",
		},
		{
			input: "yurt",
		},
	}

	for _, test := range tt {
		b.Run(fmt.Sprintf("input_name_%v", test.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				encodeFieldID(test.input)
			}
		})
	}
}

// nolint
func BenchmarkDecode(b *testing.B) {

	tt := []struct {
		input []byte
	}{
		{
			input: []byte{1, 18},
		},
		{
			input: []byte{255},
		},
	}

	for _, test := range tt {
		b.Run(fmt.Sprintf("input_name_%v", test.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				hex := hex.EncodeToString(test.input)
				decodeFieldID(hex)
			}
		})
	}
}
