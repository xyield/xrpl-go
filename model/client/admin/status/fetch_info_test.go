package status

import (
	"testing"

	"github.com/xyield/xrpl-go/test"
)

func TestFetchInfoResponse(t *testing.T) {
	s := FetchInfoResponse{
		Info: map[string]FetchInfo{
			"348928": {
				Hash:             "C26D432B06F84861BCACD7942EDC3FE0B2E1DEB966A9E516A0FD275A375C2010",
				HaveHeader:       true,
				HaveTransactions: true,
				NeededStateHashes: []string{
					"BF8DC6B1E10D1D3565BF0649075D22EBFD34F751AFCC0E53E81D74786BC88922",
					"34E37A71CB51A12C73A435250E6A6349F7884C7EEBA6B88FA31F0244E967E88F",
					"BFB7D3008A7D61FD6A0538D1C2E70CFB94CE8DC66606319C372F278A48629765",
					"41C0C61D701FB1EA586F0EF1FC7A91FEC476D979589DA60507F05C13F7C21975",
					"6DDE8840A2C3C7FF05E5FFEE4D06408694C16A8357338FE0C4581DC3D8A00BBA",
					"6C69D833B582C849917806FA009518832BB50E900E43716FD7CC1966428DD0CF",
					"1EDC020CFC4AF19B625C52E20B66D6AE672821CCC461E8A9C457A3B2955657F7",
					"FC0616A66A2B0589CA513F3341D4EA51E782C4601E5072308478E3CC19264640",
					"19FC607B5DE1B64681A676EC1ED5507B9555B0E098CD9D898320297DE1A64033",
					"5E128D3FC990074E35687387A14AA12D9FD287E5AB57CB9B2FD83DE635DF5CA9",
					"DE72820F3981770F2AA8770BC233B80661F1A452819D8529008875FF8DED87A9",
					"3ACB84BEE2C45556351FF60FD787D235C9CF5623FB8A35B01446B773598E7CC0",
					"0DD3A8DF69874148057F1F2BF305442FF2E89A76A08B4CC8C051E2ED69B874F3",
					"4AE9A9C4F12A5BD0355037DA40A0B145420A2168A9FEDE43E643BD13062F8ECE",
					"08CBF8CFFEC207F5AC4E4F24BC447011FD8C79D25B344281FBFB4732D7058ED4",
					"779B2577C5C4BAED6657421448EA506BBF50F86BE363E0924127C4EA17A58BBE",
				},
				Peers:    2,
				Timeouts: 0,
			},
		},
	}

	j := `{
	"info": {
		"348928": {
			"hash": "C26D432B06F84861BCACD7942EDC3FE0B2E1DEB966A9E516A0FD275A375C2010",
			"have_header": true,
			"have_transactions": true,
			"needed_state_hashes": [
				"BF8DC6B1E10D1D3565BF0649075D22EBFD34F751AFCC0E53E81D74786BC88922",
				"34E37A71CB51A12C73A435250E6A6349F7884C7EEBA6B88FA31F0244E967E88F",
				"BFB7D3008A7D61FD6A0538D1C2E70CFB94CE8DC66606319C372F278A48629765",
				"41C0C61D701FB1EA586F0EF1FC7A91FEC476D979589DA60507F05C13F7C21975",
				"6DDE8840A2C3C7FF05E5FFEE4D06408694C16A8357338FE0C4581DC3D8A00BBA",
				"6C69D833B582C849917806FA009518832BB50E900E43716FD7CC1966428DD0CF",
				"1EDC020CFC4AF19B625C52E20B66D6AE672821CCC461E8A9C457A3B2955657F7",
				"FC0616A66A2B0589CA513F3341D4EA51E782C4601E5072308478E3CC19264640",
				"19FC607B5DE1B64681A676EC1ED5507B9555B0E098CD9D898320297DE1A64033",
				"5E128D3FC990074E35687387A14AA12D9FD287E5AB57CB9B2FD83DE635DF5CA9",
				"DE72820F3981770F2AA8770BC233B80661F1A452819D8529008875FF8DED87A9",
				"3ACB84BEE2C45556351FF60FD787D235C9CF5623FB8A35B01446B773598E7CC0",
				"0DD3A8DF69874148057F1F2BF305442FF2E89A76A08B4CC8C051E2ED69B874F3",
				"4AE9A9C4F12A5BD0355037DA40A0B145420A2168A9FEDE43E643BD13062F8ECE",
				"08CBF8CFFEC207F5AC4E4F24BC447011FD8C79D25B344281FBFB4732D7058ED4",
				"779B2577C5C4BAED6657421448EA506BBF50F86BE363E0924127C4EA17A58BBE"
			],
			"peers": 2,
			"timeouts": 0
		}
	}
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}