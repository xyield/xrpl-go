package ledger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/CreatureDev/xrpl-go/model/client/common"
	"github.com/CreatureDev/xrpl-go/model/ledger"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/CreatureDev/xrpl-go/test"
)

func TestLedgerEntryRequest(t *testing.T) {
	s := LedgerEntryRequest{
		LedgerIndex: common.VALIDATED,
		Directory: &DirectoryEntryReq{
			Owner: "abc",
		},
	}
	j := `{
	"ledger_index": "validated",
	"directory": {
		"owner": "abc"
	}
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestLedgerEntryResponse(t *testing.T) {
	s := LedgerEntryResponse{
		Index:       "13F1A95D7AAB7108D5CE7EEAF504B2894B8C674E6D68499076441C4837282BF8",
		LedgerIndex: 61966146,
		Node: &ledger.AccountRoot{
			Account:           "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
			AccountTxnID:      "4E0AA11CBDD1760DE95B68DF2ABBE75C9698CEB548BEA9789053FCB3EBD444FB",
			Balance:           types.XRPCurrencyAmount(424021949),
			Domain:            "6D64756F31332E636F6D",
			EmailHash:         "98B4375E1D753E5B91627516F6D70977",
			Flags:             types.SetFlag(9568256),
			LedgerEntryType:   ledger.AccountRootEntry,
			MessageKey:        "0000000000000000000000070000000300",
			OwnerCount:        12,
			PreviousTxnID:     "4E0AA11CBDD1760DE95B68DF2ABBE75C9698CEB548BEA9789053FCB3EBD444FB",
			PreviousTxnLgrSeq: 61965653,
			RegularKey:        "rD9iJmieYHn8jTtPjwwkW2Wm9sVDvPXLoJ",
			Sequence:          385,
			TransferRate:      4294967295,
		},
	}
	j := `{
	"index": "13F1A95D7AAB7108D5CE7EEAF504B2894B8C674E6D68499076441C4837282BF8",
	"ledger_index": 61966146,
	"node": {
		"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
		"AccountTxnID": "4E0AA11CBDD1760DE95B68DF2ABBE75C9698CEB548BEA9789053FCB3EBD444FB",
		"Balance": "424021949",
		"Domain": "6D64756F31332E636F6D",
		"EmailHash": "98B4375E1D753E5B91627516F6D70977",
		"Flags": 9568256,
		"LedgerEntryType": "AccountRoot",
		"MessageKey": "0000000000000000000000070000000300",
		"OwnerCount": 12,
		"PreviousTxnID": "4E0AA11CBDD1760DE95B68DF2ABBE75C9698CEB548BEA9789053FCB3EBD444FB",
		"PreviousTxnLgrSeq": 61965653,
		"RegularKey": "rD9iJmieYHn8jTtPjwwkW2Wm9sVDvPXLoJ",
		"Sequence": 385,
		"TransferRate": 4294967295
	}
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestLedgerEntryValidate(t *testing.T) {
	off := &OfferEntryReq{}
	err := off.Validate()
	assert.ErrorContains(t, err, "offer")
	off.Account = "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn"
	err = off.Validate()
	assert.Nil(t, err)

	rs := RippleStateEntryReq{}
	err = rs.Validate()
	assert.ErrorContains(t, err, "requires two accounts")
	rs.Accounts = []types.Address{"rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn", ""}
	err = rs.Validate()
	assert.ErrorContains(t, err, "account 2")
	rs.Accounts[1] = "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCc"
	err = rs.Validate()
	assert.Nil(t, err)

	esc := EscrowEntryReq{}
	err = esc.Validate()
	assert.ErrorContains(t, err, "escrow entry owner")
	esc.Owner = "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn"
	err = esc.Validate()
	assert.Nil(t, err)

	dep := DepositPreauthEntryReq{}
	err = dep.Validate()
	assert.ErrorContains(t, err, "deposit preauth")
	dep.Owner = "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn"
	err = dep.Validate()
	assert.ErrorContains(t, err, "authorized")
	dep.Authorized = "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCC"
	err = dep.Validate()
	assert.Nil(t, err)

	tick := TicketEntryReq{}
	err = tick.Validate()
	assert.ErrorContains(t, err, "ticket")
	tick.Account = "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn"
	err = tick.Validate()
	assert.Nil(t, err)

	s := LedgerEntryRequest{}
	err = s.Validate()
	assert.ErrorContains(t, err, "0")
	s.AccountRoot = "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn"
	err = s.Validate()
	assert.Nil(t, err)
	s.Offer = off
	err = s.Validate()
	assert.ErrorContains(t, err, "2")
	s.AccountRoot = ""
	s.Offer = &OfferEntryReq{}
	err = s.Validate()
	assert.ErrorContains(t, err, "ledger entry offer")
}
