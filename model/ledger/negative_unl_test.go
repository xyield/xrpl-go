package ledger

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/test"
)

func TestNegativeUNL(t *testing.T) {
	var s LedgerObject = &NegativeUNL{
		DisabledValidators: []DisabledValidatorEntry{
			{
				DisabledValidator: DisabledValidator{
					FirstLedgerSequence: 1609728,
					PublicKey:           "ED6629D456285AE3613B285F65BBFF168D695BA3921F309949AFCD2CA7AFEC16FE",
				},
			},
		},
		Flags:           0,
		LedgerEntryType: NegativeUNLEntry,
	}

	j := `{
	"DisabledValidators": [
		{
			"DisabledValidator": {
				"FirstLedgerSequence": 1609728,
				"PublicKey": "ED6629D456285AE3613B285F65BBFF168D695BA3921F309949AFCD2CA7AFEC16FE"
			}
		}
	],
	"Flags": 0,
	"LedgerEntryType": "NegativeUNL"
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
