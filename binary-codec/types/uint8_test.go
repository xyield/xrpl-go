package types

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/binary-codec/definitions"
	"github.com/stretchr/testify/require"
)

func TestUint8FromJson(t *testing.T) {
	tt := []struct {
		description string
		input       any
		expected    []byte
		expectedErr error
	}{
		{
			description: "find transaction code",
			input:       "tecAMM_ACCOUNT",
			expected:    []byte{168},
			expectedErr: nil,
		},
		{
			description: "regular uint8",
			input:       uint8(30),
			expected:    []byte{30},
			expectedErr: nil,
		},
		{
			description: "regular int",
			input:       int(30),
			expected:    []byte{30},
			expectedErr: nil,
		},
		{
			description: "invalid transaction result",
			input:       "invalid",
			expected:    nil,
			expectedErr: &definitions.NotFoundError{Instance: "TransactionResultName", Input: "invalid"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			u8 := &UInt8{}
			got, err := u8.FromJson(tc.input)
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
