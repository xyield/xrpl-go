package definitions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadDefinitions(t *testing.T) {

	err := loadDefinitions()
	assert.NoError(t, err)
	assert.Equal(t, -1, definitions.Types["Done"])
	assert.Equal(t, 4, definitions.Types["Hash128"])
	assert.Equal(t, -3, definitions.LedgerEntryTypes["Any"])
	assert.Equal(t, -399, definitions.TransactionResults["telLOCAL_ERROR"])
	assert.Equal(t, 1, definitions.TransactionTypes["EscrowCreate"])
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
}

func BenchmarkLoadDefinitions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loadDefinitions()
	}
}
