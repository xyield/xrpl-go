package addresscodec

import (
	"testing"
)

func BenchmarkEncodeBase58Original(b *testing.B) {
	tt := []struct {
		description string
		input       []byte
	}{
		{
			description: "Benchmark successful encode XRP address",
			input:       []byte("rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN"),
		},
	}

	for _, tc := range tt {
		b.Run(tc.description, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				EncodeBase58Original(tc.input)
			}
		})
	}
}

func BenchmarkEncodeBase58(b *testing.B) {
	tt := []struct {
		description string
		input       []byte
	}{
		{
			description: "Benchmark successful encode XRP address",
			input:       []byte("rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN"),
		},
	}

	for _, tc := range tt {
		b.Run(tc.description, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				EncodeBase58(tc.input)
			}
		})
	}
}

func BenchmarkDecodeBase58(b *testing.B) {

	tt := []struct {
		description string
		input       string
	}{
		{
			description: "Benchmark successful encode - long string",
			input:       "s2Fku4vaPpFiqqXdAD3V5rYrSx5a9h9qvUJW3423akZSCeD",
		},
	}

	for _, tc := range tt {
		b.Run(tc.description, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				DecodeBase58(tc.input)
			}
		})
	}
}
