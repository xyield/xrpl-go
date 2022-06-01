package addresscodec

import (
	"testing"
)

func BenchmarkEncodeBase58(b *testing.B) {
	tt := []struct {
		description string
		input       []byte
	}{
		{
			description: "Benchmark XRP encode",
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
			description: "Benchmark XRP decode",
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

func BenchmarkEncodeClassicAddressFromPublicKeyHex(b *testing.B) {

	tt := []struct {
		description string
		input       string
		prefix      []byte
	}{
		{
			description: "Benchmark encode classic address",
			input:       "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
			prefix:      []byte{AccountAddressPrefix},
		},
	}

	for _, tc := range tt {
		b.Run(tc.description, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				EncodeClassicAddressFromPublicKeyHex(tc.input, tc.prefix)
			}
		})
	}

}
