package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/CreatureDev/xrpl-go/model/client/common"
	"github.com/CreatureDev/xrpl-go/model/ledger"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/CreatureDev/xrpl-go/test"
)

func TestAccountInfoRequest(t *testing.T) {
	s := AccountInfoRequest{
		Account:     "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
		LedgerIndex: common.CLOSED,
		Queue:       true,
		SignerList:  false,
		Strict:      true,
	}

	// SignerList assigned to default, omitted due to omitempty
	j := `{
	"account": "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
	"ledger_index": "closed",
	"queue": true,
	"strict": true
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestAccountInfoResponse(t *testing.T) {
	s := AccountInfoResponse{
		AccountData: ledger.AccountRoot{
			Account:           "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
			Balance:           types.XRPCurrencyAmount(999999999960),
			Flags:             types.SetFlag(8388608),
			LedgerEntryType:   ledger.AccountRootEntry,
			OwnerCount:        0,
			PreviousTxnID:     "4294BEBE5B569A18C0A2702387C9B1E7146DC3A5850C1E87204951C6FDAA4C42",
			PreviousTxnLgrSeq: 3,
			Sequence:          6,
		},
		LedgerCurrentIndex: 4,
		QueueData: QueueData{
			TxnCount:           5,
			AuthChangeQueued:   true,
			LowestSequence:     6,
			HighestSequence:    10,
			MaxSpendDropsTotal: types.XRPCurrencyAmount(500),
			Transactions: []QueueTransaction{
				{
					AuthChange:    false,
					Fee:           types.XRPCurrencyAmount(100),
					FeeLevel:      types.XRPCurrencyAmount(2560),
					MaxSpendDrops: types.XRPCurrencyAmount(100),
					Seq:           6,
				},
				{
					AuthChange:    true,
					Fee:           types.XRPCurrencyAmount(100),
					FeeLevel:      types.XRPCurrencyAmount(2560),
					MaxSpendDrops: types.XRPCurrencyAmount(100),
					Seq:           10,
				},
			},
		},
		Validated: false,
	}

	j := `{
	"account_data": {
		"Account": "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
		"Balance": "999999999960",
		"Flags": 8388608,
		"LedgerEntryType": "AccountRoot",
		"OwnerCount": 0,
		"PreviousTxnID": "4294BEBE5B569A18C0A2702387C9B1E7146DC3A5850C1E87204951C6FDAA4C42",
		"PreviousTxnLgrSeq": 3,
		"Sequence": 6
	},
	"ledger_current_index": 4,
	"queue_data": {
		"txn_count": 5,
		"auth_change_queued": true,
		"lowest_sequence": 6,
		"highest_sequence": 10,
		"max_spend_drops_total": "500",
		"transactions": [
			{
				"auth_change": false,
				"fee": "100",
				"fee_level": "2560",
				"max_spend_drops": "100",
				"seq": 6
			},
			{
				"auth_change": true,
				"fee": "100",
				"fee_level": "2560",
				"max_spend_drops": "100",
				"seq": 10
			}
		]
	},
	"validated": false
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestAccountInfoValidate(t *testing.T) {
	s := AccountInfoRequest{
		Account: "",
	}

	err := s.Validate()

	assert.ErrorContains(t, err, "missing xrpl address")
}
