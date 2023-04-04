package clio

import "github.com/xyield/xrpl-go/model/transactions/types"

type ServerInfoResponse struct {
	Info ClioLedgerInfo `json:"info"`
}

type ClioServerInfo struct {
	CompletedLedgers     string         `json:"completed_ledgers"`
	Counters             ClioCounters   `json:"counters,omitempty"`
	LoadFactor           int            `json:"load_factor"`
	ClioVersion          string         `json:"clio_version"`
	ValidationQuorum     int            `json:"validation_quorum,omitempty"`
	RippledVersion       string         `json:"rippled_version,omitempty"`
	ValidatedLedger      ClioLedgerInfo `json:"validated_ledger,omitempty"`
	ValidatorListExpires string         `json:"validator_list_expires,omitempty"`
	Cache                ClioCache      `json:"cache"`
	ETL                  ClioETL        `json:"etl,omitempty"`
	Validated            bool           `json:"validated"`
	Status               string         `json:"status,omitempty"`
}

type ClioCounters struct {
	RPC           map[string]ClioRPC `json:"rpc"`
	Subscriptions ClioSubscriptions  `json:"subscriptions"`
}

type ClioRPC struct {
	Started    string `json:"started"`
	Finished   string `json:"finished"`
	Errored    string `json:"errored"`
	Forwarded  string `json:"forwarded"`
	DurationUS string `json:"duration_us"`
}

type ClioSubscriptions struct {
	Ledger               int `json:"ledger"`
	Transactions         int `json:"transactions"`
	TransactionsProposed int `json:"transactions_proposed"`
	Manifests            int `json:"manifests"`
	Validations          int `json:"validations"`
	Account              int `json:"account"`
	AccountsProposed     int `json:"accounts_proposed"`
	Books                int `json:"books"`
}

type ClioLedgerInfo struct {
	Age            uint          `json:"age"`
	BaseFeeXRP     float32       `json:"base_fee_xrp"`
	Hash           types.Hash256 `json:"hash"`
	ReserveBaseXRP float32       `json:"reserve_base_xrp"`
	ReserveIncXRP  float32       `json:"reserve_inc_xrp"`
	Seq            uint          `json:"seq"`
}

type ClioCache struct {
	Size            int  `json:"size"`
	IsFull          bool `json:"is_full"`
	LatestLedgerSeq int  `json:"latest_ledger_seq"`
}

type ClioETL struct {
	ETLSources            []ClioETLSource `json:"etl_sources"`
	IsWriter              bool            `json:"is_writer"`
	ReadOnly              bool            `json:"read_only"`
	LastPublishAgeSeconds string          `json:"last_publish_age_seconds"`
}

type ClioETLSource struct {
	ValidatedRange    string `json:"validated_range"`
	IsConnected       string `json:"is_connected"`
	IP                string `json:"ip"`
	WSPort            string `json:"ws_port"`
	GRPCPort          string `json:"grpc_port"`
	LastMsgAgeSeconds string `json:"last_msg_age_seconds"`
}
