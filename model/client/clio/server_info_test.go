package clio

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/test"
)

func TestServerInfoResponseLocalhost(t *testing.T) {
	s := ServerInfoResponse{
		Info: ClioServerInfo{
			CompleteLedgers: "19499132-19977628",
			Counters: &ClioCounters{
				RPC: map[string]ClioRPC{
					"account_objects": {
						Started:    "1",
						Finished:   "1",
						Errored:    "0",
						Forwarded:  "0",
						DurationUS: "991",
					},
					"account_tx": {
						Started:    "1",
						Finished:   "1",
						Errored:    "0",
						Forwarded:  "0",
						DurationUS: "91633",
					},
					"account_lines": {
						Started:    "5",
						Finished:   "5",
						Errored:    "0",
						Forwarded:  "0",
						DurationUS: "4915159",
					},
					"submit_multisigned": {
						Started:    "2",
						Finished:   "2",
						Errored:    "0",
						Forwarded:  "2",
						DurationUS: "4823",
					},
					"ledger_entry": {
						Started:    "1",
						Finished:   "1",
						Errored:    "0",
						Forwarded:  "0",
						DurationUS: "17806",
					},
					"server_info": {
						Started:    "5",
						Finished:   "5",
						Errored:    "0",
						Forwarded:  "0",
						DurationUS: "2375580",
					},
					"account_info": {
						Started:    "5",
						Finished:   "5",
						Errored:    "0",
						Forwarded:  "5",
						DurationUS: "9256",
					},
					"account_currencies": {
						Started:    "4",
						Finished:   "4",
						Errored:    "0",
						Forwarded:  "0",
						DurationUS: "517302",
					},
					"noripple_check": {
						Started:    "1",
						Finished:   "1",
						Errored:    "0",
						Forwarded:  "1",
						DurationUS: "2218",
					},
					"tx": {
						Started:    "1",
						Finished:   "1",
						Errored:    "0",
						Forwarded:  "0",
						DurationUS: "562",
					},
					"gateway_balances": {
						Started:    "6",
						Finished:   "6",
						Errored:    "0",
						Forwarded:  "0",
						DurationUS: "1395156",
					},
					"channel_authorize": {
						Started:    "1",
						Finished:   "1",
						Errored:    "0",
						Forwarded:  "1",
						DurationUS: "2017",
					},
					"manifest": {
						Started:    "1",
						Finished:   "1",
						Errored:    "0",
						Forwarded:  "1",
						DurationUS: "1707",
					},
					"subscribe": {
						Started:    "6",
						Finished:   "6",
						Errored:    "0",
						Forwarded:  "0",
						DurationUS: "116",
					},
					"random": {
						Started:    "1",
						Finished:   "1",
						Errored:    "0",
						Forwarded:  "0",
						DurationUS: "111",
					},
					"ledger_data": {
						Started:    "14",
						Finished:   "3",
						Errored:    "11",
						Forwarded:  "0",
						DurationUS: "6179145",
					},
					"ripple_path_find": {
						Started:    "1",
						Finished:   "1",
						Errored:    "0",
						Forwarded:  "1",
						DurationUS: "1409563",
					},
					"account_channels": {
						Started:    "14",
						Finished:   "14",
						Errored:    "0",
						Forwarded:  "0",
						DurationUS: "1062692",
					},
					"submit": {
						Started:    "6",
						Finished:   "6",
						Errored:    "0",
						Forwarded:  "6",
						DurationUS: "11383",
					},
					"transaction_entry": {
						Started:    "8",
						Finished:   "5",
						Errored:    "3",
						Forwarded:  "0",
						DurationUS: "494131",
					},
				},
				Subscriptions: ClioSubscriptions{
					Ledger:               0,
					Transactions:         0,
					TransactionsProposed: 0,
					Manifests:            2,
					Validations:          2,
					Account:              0,
					AccountsProposed:     0,
					Books:                0,
				},
			},
			LoadFactor:       1,
			ClioVersion:      "0.3.0-b2",
			ValidationQuorum: 8,
			RippledVersion:   "1.9.1-rc1",
			ValidatedLedger: &ClioLedgerInfo{
				Age:            4,
				Hash:           "4CD25FB70D45646EE5822E76E58B66D39D5AE6BA0F70491FA803DA0DA218F434",
				Seq:            19977628,
				BaseFeeXRP:     1e-5,
				ReserveBaseXRP: 1e1,
				ReserveIncXRP:  2e0,
			},
			Cache: ClioCache{
				Size:            8812733,
				IsFull:          true,
				LatestLedgerSeq: 19977629,
			},
			ETL: &ClioETL{
				ETLSources: []ClioETLSource{
					{
						ValidatedRange:    "19405538-19977629",
						IsConnected:       "1",
						IP:                "52.36.182.38",
						WSPort:            "6005",
						GRPCPort:          "50051",
						LastMsgAgeSeconds: "0",
					},
				},
				IsWriter:              true,
				ReadOnly:              false,
				LastPublishAgeSeconds: "2",
			},
		},
		Validated: true,
	}
	j := `{
	"info": {
		"complete_ledgers": "19499132-19977628",
		"counters": {
			"rpc": {
				"account_channels": {
					"started": "14",
					"finished": "14",
					"errored": "0",
					"forwarded": "0",
					"duration_us": "1062692"
				},
				"account_currencies": {
					"started": "4",
					"finished": "4",
					"errored": "0",
					"forwarded": "0",
					"duration_us": "517302"
				},
				"account_info": {
					"started": "5",
					"finished": "5",
					"errored": "0",
					"forwarded": "5",
					"duration_us": "9256"
				},
				"account_lines": {
					"started": "5",
					"finished": "5",
					"errored": "0",
					"forwarded": "0",
					"duration_us": "4915159"
				},
				"account_objects": {
					"started": "1",
					"finished": "1",
					"errored": "0",
					"forwarded": "0",
					"duration_us": "991"
				},
				"account_tx": {
					"started": "1",
					"finished": "1",
					"errored": "0",
					"forwarded": "0",
					"duration_us": "91633"
				},
				"channel_authorize": {
					"started": "1",
					"finished": "1",
					"errored": "0",
					"forwarded": "1",
					"duration_us": "2017"
				},
				"gateway_balances": {
					"started": "6",
					"finished": "6",
					"errored": "0",
					"forwarded": "0",
					"duration_us": "1395156"
				},
				"ledger_data": {
					"started": "14",
					"finished": "3",
					"errored": "11",
					"forwarded": "0",
					"duration_us": "6179145"
				},
				"ledger_entry": {
					"started": "1",
					"finished": "1",
					"errored": "0",
					"forwarded": "0",
					"duration_us": "17806"
				},
				"manifest": {
					"started": "1",
					"finished": "1",
					"errored": "0",
					"forwarded": "1",
					"duration_us": "1707"
				},
				"noripple_check": {
					"started": "1",
					"finished": "1",
					"errored": "0",
					"forwarded": "1",
					"duration_us": "2218"
				},
				"random": {
					"started": "1",
					"finished": "1",
					"errored": "0",
					"forwarded": "0",
					"duration_us": "111"
				},
				"ripple_path_find": {
					"started": "1",
					"finished": "1",
					"errored": "0",
					"forwarded": "1",
					"duration_us": "1409563"
				},
				"server_info": {
					"started": "5",
					"finished": "5",
					"errored": "0",
					"forwarded": "0",
					"duration_us": "2375580"
				},
				"submit": {
					"started": "6",
					"finished": "6",
					"errored": "0",
					"forwarded": "6",
					"duration_us": "11383"
				},
				"submit_multisigned": {
					"started": "2",
					"finished": "2",
					"errored": "0",
					"forwarded": "2",
					"duration_us": "4823"
				},
				"subscribe": {
					"started": "6",
					"finished": "6",
					"errored": "0",
					"forwarded": "0",
					"duration_us": "116"
				},
				"transaction_entry": {
					"started": "8",
					"finished": "5",
					"errored": "3",
					"forwarded": "0",
					"duration_us": "494131"
				},
				"tx": {
					"started": "1",
					"finished": "1",
					"errored": "0",
					"forwarded": "0",
					"duration_us": "562"
				}
			},
			"subscriptions": {
				"ledger": 0,
				"transactions": 0,
				"transactions_proposed": 0,
				"manifests": 2,
				"validations": 2,
				"account": 0,
				"accounts_proposed": 0,
				"books": 0
			}
		},
		"load_factor": 1,
		"clio_version": "0.3.0-b2",
		"validation_quorum": 8,
		"rippled_version": "1.9.1-rc1",
		"validated_ledger": {
			"age": 4,
			"base_fee_xrp": 0.00001,
			"hash": "4CD25FB70D45646EE5822E76E58B66D39D5AE6BA0F70491FA803DA0DA218F434",
			"reserve_base_xrp": 10,
			"reserve_inc_xrp": 2,
			"seq": 19977628
		},
		"cache": {
			"size": 8812733,
			"is_full": true,
			"latest_ledger_seq": 19977629
		},
		"etl": {
			"etl_sources": [
				{
					"validated_range": "19405538-19977629",
					"is_connected": "1",
					"ip": "52.36.182.38",
					"ws_port": "6005",
					"grpc_port": "50051",
					"last_msg_age_seconds": "0"
				}
			],
			"is_writer": true,
			"read_only": false,
			"last_publish_age_seconds": "2"
		}
	},
	"validated": true
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestServerInfoResponseRemote(t *testing.T) {
	s := ServerInfoResponse{
		Info: ClioServerInfo{
			CompleteLedgers:  "32570-73737719",
			LoadFactor:       1,
			ClioVersion:      "1.0.2",
			ValidationQuorum: 28,
			RippledVersion:   "1.9.1",
			ValidatedLedger: &ClioLedgerInfo{
				Age:            7,
				BaseFeeXRP:     0.00001,
				Hash:           "4ECDEAF9E6F8B37EFDE297953168AAB42DEED1082A565639EBB2D29E047341B4",
				ReserveBaseXRP: 10,
				ReserveIncXRP:  2,
				Seq:            73737719,
			},
			Cache: ClioCache{
				Size:            15258947,
				IsFull:          true,
				LatestLedgerSeq: 73737719,
			},
		},
		Validated: true,
		Status:    "success",
	}

	j := `{
	"info": {
		"complete_ledgers": "32570-73737719",
		"load_factor": 1,
		"clio_version": "1.0.2",
		"validation_quorum": 28,
		"rippled_version": "1.9.1",
		"validated_ledger": {
			"age": 7,
			"base_fee_xrp": 0.00001,
			"hash": "4ECDEAF9E6F8B37EFDE297953168AAB42DEED1082A565639EBB2D29E047341B4",
			"reserve_base_xrp": 10,
			"reserve_inc_xrp": 2,
			"seq": 73737719
		},
		"cache": {
			"size": 15258947,
			"is_full": true,
			"latest_ledger_seq": 73737719
		}
	},
	"validated": true,
	"status": "success"
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
