package ledger

import (
	"testing"

	"github.com/xyield/xrpl-go/test"
)

func TestFeeSettings(t *testing.T) {
	var s LedgerObject
	s = &FeeSettings{
		BaseFee:           "000000000000000A",
		Flags:             0,
		LedgerEntryType:   FeeSettingsEntry,
		ReferenceFeeUnits: 10,
		ReserveBase:       20000000,
		ReserveIncrement:  5000000,
	}

	j := `{
	"BaseFee": "000000000000000A",
	"Flags": 0,
	"LedgerEntryType": "FeeSettings",
	"ReferenceFeeUnits": 10,
	"ReserveBase": 20000000,
	"ReserveIncrement": 5000000
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
