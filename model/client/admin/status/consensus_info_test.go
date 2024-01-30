package status

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/test"
)

func TestConsensusInfoResponse(t *testing.T) {
	s := ConsensusInfoResponse{
		Info: ConsensusInfo{
			Acquired: map[string]string{
				"4BC2CE596CBD1321775320E2067F9C06D3862826212C16EF42ABB6A2B0414306": "acquired",
			},
			CloseGranularity: 10,
			ClosePercent:     50,
			CloseResolution:  10,
			CloseTimes: map[string]int{
				"486082972": 1,
				"486082973": 4,
			},
			CurrentMs:         1003,
			HaveTimeConsensus: false,
			LedgerSeq:         13701086,
			OurPosition: Position{
				CloseTime:       486082973,
				PreviousLedger:  "0BB01379B51234BAAF501A71C7AB147F595460B689BB9E8252A0B87B5A483623",
				ProposeSeq:      0,
				TransactionHash: "4BC2CE596CBD1321775320E2067F9C06D3862826212C16EF42ABB6A2B0414306",
			},
			PeerPositions: map[string]Position{
				"0A2EAF919033A036D363D4E5610A66209DDBE8EE": {
					CloseTime:       486082972,
					PeerId:          "n9KiYM9CgngLvtRCQHZwgC2gjpdaZcCcbt3VboxiNFcKuwFVujzS",
					PreviousLedger:  "0BB01379B51234BAAF501A71C7AB147F595460B689BB9E8252A0B87B5A483623",
					ProposeSeq:      0,
					TransactionHash: "4BC2CE596CBD1321775320E2067F9C06D3862826212C16EF42ABB6A2B0414306",
				},
				"1567A8C953A86F8428C7B01641D79BBF2FD508F3": {
					CloseTime:       486082973,
					PeerId:          "n9LdgEtkmGB9E2h3K4Vp7iGUaKuq23Zr32ehxiU8FWY7xoxbWTSA",
					PreviousLedger:  "0BB01379B51234BAAF501A71C7AB147F595460B689BB9E8252A0B87B5A483623",
					ProposeSeq:      0,
					TransactionHash: "4BC2CE596CBD1321775320E2067F9C06D3862826212C16EF42ABB6A2B0414306",
				},
				"202397A81F20B44CF44EA99AF761295E5A8397D2": {
					CloseTime:       486082973,
					PeerId:          "n9MD5h24qrQqiyBC8aeqqCWvpiBiYQ3jxSr91uiDvmrkyHRdYLUj",
					PreviousLedger:  "0BB01379B51234BAAF501A71C7AB147F595460B689BB9E8252A0B87B5A483623",
					ProposeSeq:      0,
					TransactionHash: "4BC2CE596CBD1321775320E2067F9C06D3862826212C16EF42ABB6A2B0414306",
				},
				"5C29005CF4FB479FC49EEFB4A5B075C86DD963CC": {
					CloseTime:       486082973,
					PeerId:          "n9L81uNCaPgtUJfaHh89gmdvXKAmSt5Gdsw2g1iPWaPkAHW5Nm4C",
					PreviousLedger:  "0BB01379B51234BAAF501A71C7AB147F595460B689BB9E8252A0B87B5A483623",
					ProposeSeq:      0,
					TransactionHash: "4BC2CE596CBD1321775320E2067F9C06D3862826212C16EF42ABB6A2B0414306",
				},
				"EFC49EB648E557CC50A72D715249B80E071F7705": {
					CloseTime:       486082973,
					PeerId:          "n949f75evCHwgyP4fPVgaHqNHxUVN15PsJEZ3B3HnXPcPjcZAoy7",
					PreviousLedger:  "0BB01379B51234BAAF501A71C7AB147F595460B689BB9E8252A0B87B5A483623",
					ProposeSeq:      0,
					TransactionHash: "4BC2CE596CBD1321775320E2067F9C06D3862826212C16EF42ABB6A2B0414306",
				},
			},
			PreviousMseconds:  2005,
			PreviousProposers: 5,
			Proposers:         5,
			State:             "consensus",
			Synched:           true,
		},
	}

	j := `{
	"info": {
		"acquired": {
			"4BC2CE596CBD1321775320E2067F9C06D3862826212C16EF42ABB6A2B0414306": "acquired"
		},
		"close_granularity": 10,
		"close_percent": 50,
		"close_resolution": 10,
		"close_times": {
			"486082972": 1,
			"486082973": 4
		},
		"current_ms": 1003,
		"have_time_consensus": false,
		"ledger_seq": 13701086,
		"our_position": {
			"close_time": 486082973,
			"previous_ledger": "0BB01379B51234BAAF501A71C7AB147F595460B689BB9E8252A0B87B5A483623",
			"propose_seq": 0,
			"transaction_hash": "4BC2CE596CBD1321775320E2067F9C06D3862826212C16EF42ABB6A2B0414306"
		},
		"peer_positions": {
			"0A2EAF919033A036D363D4E5610A66209DDBE8EE": {
				"close_time": 486082972,
				"peer_id": "n9KiYM9CgngLvtRCQHZwgC2gjpdaZcCcbt3VboxiNFcKuwFVujzS",
				"previous_ledger": "0BB01379B51234BAAF501A71C7AB147F595460B689BB9E8252A0B87B5A483623",
				"propose_seq": 0,
				"transaction_hash": "4BC2CE596CBD1321775320E2067F9C06D3862826212C16EF42ABB6A2B0414306"
			},
			"1567A8C953A86F8428C7B01641D79BBF2FD508F3": {
				"close_time": 486082973,
				"peer_id": "n9LdgEtkmGB9E2h3K4Vp7iGUaKuq23Zr32ehxiU8FWY7xoxbWTSA",
				"previous_ledger": "0BB01379B51234BAAF501A71C7AB147F595460B689BB9E8252A0B87B5A483623",
				"propose_seq": 0,
				"transaction_hash": "4BC2CE596CBD1321775320E2067F9C06D3862826212C16EF42ABB6A2B0414306"
			},
			"202397A81F20B44CF44EA99AF761295E5A8397D2": {
				"close_time": 486082973,
				"peer_id": "n9MD5h24qrQqiyBC8aeqqCWvpiBiYQ3jxSr91uiDvmrkyHRdYLUj",
				"previous_ledger": "0BB01379B51234BAAF501A71C7AB147F595460B689BB9E8252A0B87B5A483623",
				"propose_seq": 0,
				"transaction_hash": "4BC2CE596CBD1321775320E2067F9C06D3862826212C16EF42ABB6A2B0414306"
			},
			"5C29005CF4FB479FC49EEFB4A5B075C86DD963CC": {
				"close_time": 486082973,
				"peer_id": "n9L81uNCaPgtUJfaHh89gmdvXKAmSt5Gdsw2g1iPWaPkAHW5Nm4C",
				"previous_ledger": "0BB01379B51234BAAF501A71C7AB147F595460B689BB9E8252A0B87B5A483623",
				"propose_seq": 0,
				"transaction_hash": "4BC2CE596CBD1321775320E2067F9C06D3862826212C16EF42ABB6A2B0414306"
			},
			"EFC49EB648E557CC50A72D715249B80E071F7705": {
				"close_time": 486082973,
				"peer_id": "n949f75evCHwgyP4fPVgaHqNHxUVN15PsJEZ3B3HnXPcPjcZAoy7",
				"previous_ledger": "0BB01379B51234BAAF501A71C7AB147F595460B689BB9E8252A0B87B5A483623",
				"propose_seq": 0,
				"transaction_hash": "4BC2CE596CBD1321775320E2067F9C06D3862826212C16EF42ABB6A2B0414306"
			}
		},
		"previous_mseconds": 2005,
		"previous_proposers": 5,
		"proposers": 5,
		"state": "consensus",
		"synched": true
	}
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
