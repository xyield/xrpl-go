package definitions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadDefinitions(t *testing.T) {

	err := LoadDefinitions()
	assert.NoError(t, err)
	assert.Equal(t, int64(-1), definitions.Types["Done"])
	assert.Equal(t, int64(4), definitions.Types["Hash128"])
	assert.Equal(t, int64(-3), definitions.LedgerEntryTypes["Any"])
	assert.Equal(t, int64(-399), definitions.TransactionResults["telLOCAL_ERROR"])
	assert.Equal(t, int64(1), definitions.TransactionTypes["EscrowCreate"])
	assert.Equal(t, fieldInfo{Nth: int64(0), IsVLEncoded: false, IsSerialized: false, IsSigningField: false, Type: "Unknown"}, definitions.Fields["Generic"].FieldInfo)
	assert.Equal(t, fieldInfo{Nth: int64(28), IsVLEncoded: false, IsSerialized: true, IsSigningField: true, Type: "Hash256"}, definitions.Fields["NFTokenBuyOffer"].FieldInfo)
	assert.Equal(t, fieldInfo{Nth: int64(16), IsVLEncoded: false, IsSerialized: true, IsSigningField: true, Type: "UInt8"}, definitions.Fields["TickSize"].FieldInfo)
}
