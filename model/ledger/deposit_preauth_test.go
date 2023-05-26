package ledger

import (
	"testing"

	"github.com/xyield/xrpl-go/test"
)

func TestDepositPreauth(t *testing.T) {
	var s LedgerObject
	s = &DepositPreauthObj{
		LedgerEntryType:   DepositPreauthObjEntry,
		Account:           "rsUiUMpnrgxQp24dJYZDhmV4bE3aBtQyt8",
		Authorize:         "rEhxGqkqPPSxQ3P25J66ft5TwpzV14k2de",
		Flags:             0,
		OwnerNode:         "0000000000000000",
		PreviousTxnID:     "3E8964D5A86B3CD6B9ECB33310D4E073D64C865A5B866200AD2B7E29F8326702",
		PreviousTxnLgrSeq: 7,
	}

	j := `{
	"Account": "rsUiUMpnrgxQp24dJYZDhmV4bE3aBtQyt8",
	"Authorize": "rEhxGqkqPPSxQ3P25J66ft5TwpzV14k2de",
	"Flags": 0,
	"LedgerEntryType": "DepositPreauth",
	"OwnerNode": "0000000000000000",
	"PreviousTxnID": "3E8964D5A86B3CD6B9ECB33310D4E073D64C865A5B866200AD2B7E29F8326702",
	"PreviousTxnLgrSeq": 7
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
