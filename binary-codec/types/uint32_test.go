package types

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/stretchr/testify/require"
)

func TestUint32FromJson(t *testing.T) {
	tt := []struct {
		description string
		input       any
		expected    []byte
		expectedErr error
	}{
		{
			description: "convert uint32",
			input:       uint32(1),
			expected:    []byte{0, 0, 0, 1},
			expectedErr: nil,
		},
		{
			description: "convert uint",
			input:       uint(1),
			expected:    []byte{0, 0, 0, 1},
			expectedErr: nil,
		},
		{
			description: "convert flag",
			input:       types.SetFlag(1),
			expected:    []byte{0, 0, 0, 1},
			expectedErr: nil,
		},
		{
			description: "invalid type should error",
			input:       "invalid",
			expected:    nil,
			expectedErr: ErrInvalidUInt32,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			u32 := &UInt32{}
			got, err := u32.FromJson(tc.input)
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
