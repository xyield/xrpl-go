package types

import (
	"bytes"
	"strings"
	"testing"

	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/stretchr/testify/require"
)

func TestHash256FromJson(t *testing.T) {
	tt := []struct {
		description string
		input       any
		expected    []byte
		expectedErr error
	}{
		{
			description: "convert string",
			input:       strings.Repeat("0", 63) + "1",
			expected:    append(bytes.Repeat([]byte{0}, 31), byte(1)),
			expectedErr: nil,
		},
		{
			description: "convert hash256 type",
			input:       types.Hash256(strings.Repeat("0", 63) + "1"),
			expected:    append(bytes.Repeat([]byte{0}, 31), byte(1)),
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
			h256 := &Hash256{}
			got, err := h256.FromJson(tc.input)
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
