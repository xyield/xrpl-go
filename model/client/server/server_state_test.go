package server

import (
	"testing"

	"github.com/xyield/xrpl-go/test"
)

func TestServerStateResponse(t *testing.T) {
	s := ServerStateResponse{
		State: ServerState{
			BuildVersion:    "1.7.2",
			CompleteLedgers: "64572720-65887201",
			IOLatencyMS:     1,
			JQTransOverflow: "0",
			LastClose: &ServerStateClose{
				ConvergeTime: 3005,
				Proposers:    41,
			},
			LoadBase:                256,
			LoadFactor:              256,
			LoadFactorFeeEscelation: 256,
			LoadFactorFeeQueue:      256,
			LoadFactorFeeReference:  256,
			LoadFactorServer:        256,
			Peers:                   216,
			PubkeyNode:              "n9MozjnGB3tpULewtTsVtuudg5JqYFyV3QFdAtVLzJaxHcBaxuXD",
			ServerState:             "full",
			ServerStateDurationUS:   "3588969453592",
			StateAccounting: map[string]ServerStateAccounting{
				"connected": {
					DurationUS:  "301410595",
					Transitions: "2",
				},
				"disconnected": {
					DurationUS:  "1207534",
					Transitions: "2",
				},
				"full": {
					DurationUS:  "3589171798767",
					Transitions: "2",
				},
				"syncing": {
					DurationUS:  "6182323",
					Transitions: "2",
				},
				"tracking": {
					DurationUS:  "43",
					Transitions: "2",
				},
			},
			Time:   "2021-Aug-24 20:44:43.466048 UTC",
			Uptime: 3589480,
			ValidatedLedger: &ServerLedgerState{
				BaseFee:     10,
				CloseTime:   683153081,
				Hash:        "B52AC3876412A152FE9C0442801E685D148D05448D0238587DBA256330A98FD3",
				ReserveBase: 20000000,
				ReserveInc:  5000000,
				Seq:         65887201,
			},
			ValidationQuorum: 33,
		},
	}

	j := `{
	"state": {
		"build_version": "1.7.2",
		"complete_ledgers": "64572720-65887201",
		"io_latency_ms": 1,
		"jq_trans_overflow": "0",
		"last_close": {
			"converge_time": 3005,
			"proposers": 41
		},
		"load_base": 256,
		"load_factor": 256,
		"load_factor_fee_escalation": 256,
		"load_factor_fee_queue": 256,
		"load_factor_fee_reference": 256,
		"load_factor_server": 256,
		"peers": 216,
		"pubkey_node": "n9MozjnGB3tpULewtTsVtuudg5JqYFyV3QFdAtVLzJaxHcBaxuXD",
		"server_state": "full",
		"server_state_duration_us": "3588969453592",
		"state_accounting": {
			"connected": {
				"duration_us": "301410595",
				"transitions": "2"
			},
			"disconnected": {
				"duration_us": "1207534",
				"transitions": "2"
			},
			"full": {
				"duration_us": "3589171798767",
				"transitions": "2"
			},
			"syncing": {
				"duration_us": "6182323",
				"transitions": "2"
			},
			"tracking": {
				"duration_us": "43",
				"transitions": "2"
			}
		},
		"time": "2021-Aug-24 20:44:43.466048 UTC",
		"uptime": 3589480,
		"validated_ledger": {
			"base_fee": 10,
			"close_time": 683153081,
			"hash": "B52AC3876412A152FE9C0442801E685D148D05448D0238587DBA256330A98FD3",
			"reserve_base": 20000000,
			"reserve_inc": 5000000,
			"seq": 65887201
		},
		"validation_quorum": 33
	}
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
