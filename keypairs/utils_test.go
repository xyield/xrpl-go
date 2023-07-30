package keypairs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSha512Half(t *testing.T) {
	tt := []struct {
		description string
		input       []byte
		expected    []byte
	}{
		{
			description: "hash of fakeRandomString",
			input:       []byte{102, 97, 107, 101, 82, 97, 110, 100, 111, 109, 83, 116, 114, 105, 110, 103},
			expected:    []byte{187, 62, 202, 137, 133, 225, 72, 79, 166, 162, 140, 75, 48, 251, 0, 66, 162, 204, 93, 243, 236, 141, 195, 123, 95, 61, 18, 109, 223, 211, 202, 20},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			got := sha512Half(tc.input)
			require.Equal(t, tc.expected, got)
		})
	}
}
