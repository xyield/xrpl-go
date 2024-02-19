package types

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/stretchr/testify/require"
)

func TestHash128FromJson(t *testing.T) {
	tt := []struct {
		description string
		input       any
		expected    []byte
		expectedErr error
	}{
		{
			description: "convert string",
			input:       "00000000000000000000000000000001",
			expected:    []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			expectedErr: nil,
		},
		{
			description: "convert hash128 type",
			input:       types.Hash128("00000000000000000000000000000001"),
			expected:    []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			expectedErr: nil,
		},
		{
			description: "invalid type should error",
			input:       -54,
			expected:    nil,
			expectedErr: ErrInvalidHashType,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			h128 := &Hash128{}
			got, err := h128.FromJson(tc.input)
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
