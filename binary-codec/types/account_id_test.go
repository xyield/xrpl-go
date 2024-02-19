package types

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/stretchr/testify/require"
)

func TestAccountIDFromJson(t *testing.T) {
	tt := []struct {
		description string
		input       any
		expected    []byte
		expectedErr error
	}{
		{
			description: "convert address",
			input:       types.Address("r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59"),
			expected:    []byte{0x5e, 0x7b, 0x11, 0x25, 0x23, 0xf6, 0x8d, 0x2f, 0x5e, 0x87, 0x9d, 0xb4, 0xea, 0xc5, 0x1c, 0x66, 0x98, 0xa6, 0x93, 0x4},
			expectedErr: nil,
		},
		{
			description: "convert address from string type",
			input:       "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
			expected:    []byte{0x5e, 0x7b, 0x11, 0x25, 0x23, 0xf6, 0x8d, 0x2f, 0x5e, 0x87, 0x9d, 0xb4, 0xea, 0xc5, 0x1c, 0x66, 0x98, 0xa6, 0x93, 0x4},
			expectedErr: nil,
		},
		{
			description: "invalid type should error",
			input:       false,
			expected:    nil,
			expectedErr: ErrInvalidAccountID,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			a := &AccountID{}
			got, err := a.FromJson(tc.input)
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
