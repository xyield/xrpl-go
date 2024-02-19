package types

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/model/ledger"
	"github.com/CreatureDev/xrpl-go/model/transactions"
	"github.com/stretchr/testify/require"
)

func TestUInt16FromJson(t *testing.T) {
	tt := []struct {
		description string
		input       any
		expected    []byte
		expectedErr error
	}{
		{
			description: "convert uint16",
			input:       uint16(1),
			expected:    []byte{0, 1},
			expectedErr: nil,
		},
		{
			description: "convert uint",
			input:       uint(1),
			expected:    []byte{0, 1},
			expectedErr: nil,
		},
		{
			description: "convert int",
			input:       int(1),
			expected:    []byte{0, 1},
			expectedErr: nil,
		},
		{
			description: "convert TxType",
			input:       transactions.PaymentTx,
			expected:    []byte{0, 0},
			expectedErr: nil,
		},
		{
			description: "convert LedgerEntryType",
			input:       ledger.AccountRootEntry,
			expected:    []byte{0, 97},
			expectedErr: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			u16 := &UInt16{}
			got, err := u16.FromJson(tc.input)
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
