package ledger

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/test"
)

// Ledger closed request does not have any fields to test

func TestLedgerClosedResponse(t *testing.T) {
	s := LedgerClosedResponse{
		LedgerHash:  "abc",
		LedgerIndex: 123,
	}
	j := `{
	"ledger_hash": "abc",
	"ledger_index": 123
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
