package serdes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadVariableLength(t *testing.T) {
	tt := []struct {
		description string
		input       []byte
		output      int
	}{
		{
			description: "Length less than 193",
			input:       []byte{190, 230, 131},
			output:      190,
		},
		{
			description: "length > 192 & length < 241",
			input:       []byte{195, 230, 112, 234, 98},
			output:      935,
		},
		{
			description: "length > 240 & length < 255",
			input:       []byte{242, 112, 78, 95, 115},
			output:      106767,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			p := NewBinaryParser(tc.input)
			actual, _ := p.ReadVariableLength()
			require.Equal(t, tc.output, actual)
		})
	}
}
