package definitions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadDefinitions(t *testing.T) {

	err := loadDefinitions()
	assert.NoError(t, err)
	assert.Equal(t, int32(-1), definitions.Types["Done"])
	assert.Equal(t, int32(4), definitions.Types["Hash128"])
	assert.Equal(t, int32(-3), definitions.LedgerEntryTypes["Any"])
	assert.Equal(t, int32(-399), definitions.TransactionResults["telLOCAL_ERROR"])
	assert.Equal(t, int32(1), definitions.TransactionTypes["EscrowCreate"])
	assert.Equal(t, &fieldInfo{Nth: 0, IsVLEncoded: false, IsSerialized: false, IsSigningField: false, Type: "Unknown"}, definitions.Fields["Generic"].fieldInfo)
	assert.Equal(t, &fieldInfo{Nth: 28, IsVLEncoded: false, IsSerialized: true, IsSigningField: true, Type: "Hash256"}, definitions.Fields["NFTokenBuyOffer"].fieldInfo)
	assert.Equal(t, &fieldInfo{Nth: 16, IsVLEncoded: false, IsSerialized: true, IsSigningField: true, Type: "UInt8"}, definitions.Fields["TickSize"].fieldInfo)
	assert.Equal(t, &fieldHeader{TypeCode: 2, FieldCode: 4}, definitions.Fields["Sequence"].FieldHeader)
	assert.Equal(t, &fieldHeader{TypeCode: 18, FieldCode: 1}, definitions.Fields["Paths"].FieldHeader)
	assert.Equal(t, &fieldHeader{TypeCode: 2, FieldCode: 33}, definitions.Fields["SetFlag"].FieldHeader)
	assert.Equal(t, &fieldHeader{TypeCode: 16, FieldCode: 16}, definitions.Fields["TickSize"].FieldHeader)
	assert.Equal(t, "UInt32", definitions.Fields["TransferRate"].Type)
	assert.Equal(t, "Sequence", definitions.FieldIdNameMap[fieldHeader{TypeCode: 2, FieldCode: 4}])
	assert.Equal(t, "OfferSequence", definitions.FieldIdNameMap[fieldHeader{TypeCode: 2, FieldCode: 25}])
	assert.Equal(t, "NFTokenSellOffer", definitions.FieldIdNameMap[fieldHeader{TypeCode: 5, FieldCode: 29}])
	assert.Equal(t, int32(131076), definitions.Fields["Sequence"].Ordinal)
	assert.Equal(t, int32(131097), definitions.Fields["OfferSequence"].Ordinal)
}

// Helper functions to create and test ordinals.
// func CreateOrdinal(fh fieldHeader) int32 {
// 	return fh.TypeCode<<16 | fh.FieldCode
// }

// func TestCreateOrdinal(t *testing.T) {
// 	tt := []struct {
// 		description string
// 		input       fieldHeader
// 	}{
// 		{
// 			description: "test ordinal creation",
// 			input:       fieldHeader{TypeCode: 2, FieldCode: 25},
// 		},
// 	}

// 	for _, tc := range tt {
// 		t.Run(tc.description, func(t *testing.T) {
// 			fmt.Println("Ordinal:", CreateOrdinal(tc.input))
// 		})
// 	}
// }

//nolint
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
			assert.Equal(t, tc.expected, convertIntToBytes(tc.input))
		})
	}
}
