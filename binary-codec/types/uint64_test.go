package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUint64FromJson(t *testing.T) {
	tt := []struct {
		description string
		input       any
		expected    []byte
		expectedErr error
	}{
		{
			description: "convert uint64",
			input:       uint64(1),
			expected:    []byte{0, 0, 0, 0, 0, 0, 0, 1},
			expectedErr: nil,
		},
		{
			description: "convert uint",
			input:       uint(1),
			expected:    []byte{0, 0, 0, 0, 0, 0, 0, 1},
			expectedErr: nil,
		},
		{
			description: "convert hex encoded string",
			input:       "0000000000000001",
			expected:    []byte{0, 0, 0, 0, 0, 0, 0, 1},
			expectedErr: nil,
		},
		{
			description: "convert short hex encoded string",
			input:       "10",
			expected:    []byte{0, 0, 0, 0, 0, 0, 0, 16},
			expectedErr: nil,
		},
		{
			description: "invalid string should error",
			input:       "invalid",
			expected:    nil,
			expectedErr: ErrInvalidUInt64String,
		},
		{
			description: "invalid type should error",
			input:       -54,
			expected:    nil,
			expectedErr: ErrInvalidUInt64Value,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			u64 := &UInt64{}
			got, err := u64.FromJson(tc.input)
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
