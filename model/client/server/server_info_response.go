package server

import "github.com/xyield/xrpl-go/model/transactions/types"

type ServerInfoResponse struct {
	Info ServerInfo `json:"info"`
}

type ServerInfo struct {
	AmendmentBlocked        bool                            `json:"amendment_blocked,omitempty"`
	BuildVersion            string                          `json:"build_version"`
	ClosedLedger            ServerLedgerInfo                `json:"closed_ledger,omitempty"`
	CompleteLedgers         string                          `json:"complete_ledgers"`
	HostID                  string                          `json:"hostid"`
	IOLatencyMS             uint                            `json:"io_latency_ms"`
	JQTransOverflow         string                          `json:"jq_trans_overflow"`
	LastClose               ServerClose                     `json:"last_close"`
	Load                    ServerLoad                      `json:"load,omitempty"`
	LoadFactor              uint                            `json:"load_factor"`
	LoadFactorLocal         uint                            `json:"load_factor_local,omitempty"`
	LoadFactorNet           uint                            `json:"load_factor_net,omitempty"`
	LoadFactorCluster       uint                            `json:"load_factor_cluster"`
	LoadFactorFeeEscelation uint                            `json:"load_factor_fee_escelation,omitempty"`
	LoadFactorFeeQueue      uint                            `json:"load_factor_fee_queue,omitempty"`
	LoadFactorServer        uint                            `json:"load_factor_server,omitempty"`
	Peers                   uint                            `json:"peers,omitempty"`
	PubkeyNode              string                          `json:"pubkey_node"`
	PubkeyValidator         string                          `json:"pubkey_validator,omitempty"`
	Reporting               ServerReporting                 `json:"reporting,omitempty"`
	ServerState             string                          `json:"server_state"`
	ServerStateDurationUS   string                          `json:"server_state_duration_us"`
	StateAccounting         map[string]ServerInfoAccounting `json:"state_accounting"`
	Time                    string                          `json:"time"`
	Uptime                  uint                            `json:"uptime"`
	ValidatedLedger         ServerLedgerInfo                `json:"validated_ledger,omitempty"`
	ValidationQuorum        uint                            `json:"validation_quorum"`
	ValidatorListExpires    string                          `json:"validator_list_expires,omitempty"`
}

type ServerInfoAccounting struct {
	DurationUS  string `json:"duration_us"`
	Transitions string `json:"transitions"`
}

type ServerReporting struct {
	ETLSources      []ETLSource `json:"etl_sources"`
	IsWriter        bool        `json:"is_writer"`
	LastPublishTime string      `json:"last_publish_time"`
}

type ETLSource struct {
	Connected              bool   `json:"connected"`
	GRPCPort               string `json:"grpc_port"`
	IP                     string `json:"ip"`
	LastMessageArrivalTime string `json:"last_message_arrival_time"`
	ValidatedLedgersRange  string `json:"validated_ledgers_range"`
	WebsocketPort          string `json:"websocket_port"`
}

type ServerLoad struct {
	// TODO determine job types array format
	JobTypes []interface{} `json:"job_types"`
	Threads  uint          `json:"threads"`
}

type ServerClose struct {
	ConvergeTimeS float32 `json:"converge_time_s"`
	Proposers     uint    `json:"proposers"`
}

type ServerLedgerInfo struct {
	Age            uint          `json:"age"`
	BaseFeeXRP     float32       `json:"base_fee_xrp"`
	Hash           types.Hash256 `json:"hash"`
	ReserveBaseXRP float32       `json:"reserve_base_xrp"`
	ReserveIncXRP  float32       `json:"reserve_inc_xrp"`
	Seq            uint          `json:"seq"`
}
