package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBlobFromJson(t *testing.T) {
	tt := []struct {
		description string
		input       any
		expected    []byte
		expectedErr error
	}{
		{
			description: "convert string",
			input:       "0000000000000001",
			expected:    []byte{0, 0, 0, 0, 0, 0, 0, 1},
			expectedErr: nil,
		},
		{
			description: "invalid type should error",
			input:       -54,
			expected:    nil,
			expectedErr: ErrInvalidBlobType,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			b := &Blob{}
			got, err := b.FromJson(tc.input)
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
