package ledger

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/test"
)

// Ledger Current request has no fields to test

func TestLedgerCurrentResponse(t *testing.T) {
	s := LedgerCurrentResponse{
		LedgerCurrentIndex: 123,
	}
	j := `{
	"ledger_current_index": 123
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
