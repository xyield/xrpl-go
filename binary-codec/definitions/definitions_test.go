package definitions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadDefinitions(t *testing.T) {

	err := loadDefinitions()
	require.NoError(t, err)
	require.Equal(t, int32(-1), definitions.Types["Done"])
	require.Equal(t, int32(4), definitions.Types["Hash128"])
	require.Equal(t, int32(-3), definitions.LedgerEntryTypes["Any"])
	require.Equal(t, int32(-399), definitions.TransactionResults["telLOCAL_ERROR"])
	require.Equal(t, int32(1), definitions.TransactionTypes["EscrowCreate"])
	require.Equal(t, &fieldInfo{Nth: 0, IsVLEncoded: false, IsSerialized: false, IsSigningField: false, Type: "Unknown"}, definitions.Fields["Generic"].fieldInfo)
	require.Equal(t, &fieldInfo{Nth: 28, IsVLEncoded: false, IsSerialized: true, IsSigningField: true, Type: "Hash256"}, definitions.Fields["NFTokenBuyOffer"].fieldInfo)
	require.Equal(t, &fieldInfo{Nth: 16, IsVLEncoded: false, IsSerialized: true, IsSigningField: true, Type: "UInt8"}, definitions.Fields["TickSize"].fieldInfo)
	require.Equal(t, &FieldHeader{TypeCode: 2, FieldCode: 4}, definitions.Fields["Sequence"].FieldHeader)
	require.Equal(t, &FieldHeader{TypeCode: 18, FieldCode: 1}, definitions.Fields["Paths"].FieldHeader)
	require.Equal(t, &FieldHeader{TypeCode: 2, FieldCode: 33}, definitions.Fields["SetFlag"].FieldHeader)
	require.Equal(t, &FieldHeader{TypeCode: 16, FieldCode: 16}, definitions.Fields["TickSize"].FieldHeader)
	require.Equal(t, "UInt32", definitions.Fields["TransferRate"].Type)
	require.Equal(t, "Sequence", definitions.FieldIdNameMap[FieldHeader{TypeCode: 2, FieldCode: 4}])
	require.Equal(t, "OfferSequence", definitions.FieldIdNameMap[FieldHeader{TypeCode: 2, FieldCode: 25}])
	require.Equal(t, "NFTokenSellOffer", definitions.FieldIdNameMap[FieldHeader{TypeCode: 5, FieldCode: 29}])
	require.Equal(t, int32(131076), definitions.Fields["Sequence"].Ordinal)
	require.Equal(t, int32(131097), definitions.Fields["OfferSequence"].Ordinal)
}

// Helper functions to create and test ordinals.
// func CreateOrdinal(fh FieldHeader) int32 {
// 	return fh.TypeCode<<16 | fh.FieldCode
// }

// func TestCreateOrdinal(t *testing.T) {
// 	tt := []struct {
// 		description string
// 		input       FieldHeader
// 	}{
// 		{
// 			description: "test ordinal creation",
// 			input:       FieldHeader{TypeCode: 2, FieldCode: 25},
// 		},
// 	}

// 	for _, tc := range tt {
// 		t.Run(tc.description, func(t *testing.T) {
// 			fmt.Println("Ordinal:", CreateOrdinal(tc.input))
// 		})
// 	}
// }

// nolint
func BenchmarkLoadDefinitions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loadDefinitions()
	}
}

func TestConvertIntToBytes(t *testing.T) {
	tt := []struct {
		description string
		input       int32
		expected    []byte
	}{
		{
			description: "Convert int < 256 to bytes",
			input:       3,
			expected:    []byte{3},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			require.Equal(t, tc.expected, convertIntToBytes(tc.input))
		})
	}
}
