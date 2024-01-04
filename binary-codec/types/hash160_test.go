package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHash160FromJson(t *testing.T) {
	tt := []struct {
		description string
		input       any
		expected    []byte
		expectedErr error
	}{
		{
			description: "convert string",
			input:       "0000000000000000000000000000000000000001",
			expected:    []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			expectedErr: nil,
		},
		{
			description: "invalid type",
			input:       -54,
			expected:    nil,
			expectedErr: ErrInvalidHashType,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			h160 := &Hash160{}
			got, err := h160.FromJson(tc.input)
			if tc.expectedErr != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
				require.Empty(t, got)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, got)
			}
		})

	}
}
