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
