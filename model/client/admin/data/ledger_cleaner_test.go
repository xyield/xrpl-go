package data

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/test"
)

func TestLedgerCleanerRequest(t *testing.T) {
	s := LedgerCleanerRequest{
		Ledger:    7,
		MaxLedger: 5,
		MinLedger: 3,
		Stop:      true,
	}

	j := `{
	"ledger": 7,
	"max_ledger": 5,
	"min_ledger": 3,
	"stop": true
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestLedgerCleanerResponse(t *testing.T) {
	s := LedgerCleanerResponse{
		Message: "Cleaner configured",
	}

	j := `{
	"message": "Cleaner configured"
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
