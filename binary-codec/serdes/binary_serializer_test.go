package serdes

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncodeVariableLength(t *testing.T) {
	tt := []struct {
		description string
		len         int
		expected    []byte
		expectedErr error
	}{
		{
			description: "length less than 193",
			len:         100,
			expected:    []byte{0x64},
			expectedErr: nil,
		},
		{
			description: "length more than 193 and less than 12481",
			len:         1000,
			expected:    []byte{0xC4, 0x27},
			expectedErr: nil,
		},
		{
			description: "length more than 12841 ad less than 918744",
			len:         20000,
			expected:    []byte{0xF1, 0x1D, 0x5F},
			expectedErr: nil,
		},
		{
			description: "length more than 918744",
			len:         1000000,
			expected:    nil,
			expectedErr: ErrLengthPrefixTooLong,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			s := strings.Repeat("A2", tc.len)
			b, _ := hex.DecodeString(s)
			require.Equal(t, tc.len, len(b))
			actual, err := encodeVariableLength(len(b))
			if tc.expectedErr != nil {
				require.Error(t, err, tc.expectedErr.Error())
				require.Nil(t, actual)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, actual)
			}
		})
	}
}
