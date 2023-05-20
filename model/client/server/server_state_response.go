package server

type ServerStateResponse struct {
	State ServerState `json:"state"`
}

type ServerState struct {
	AmendmentBlocked        bool                             `json:"amendment_blocked,omitempty"`
	BuildVersion            string                           `json:"build_version"`
	CompleteLedgers         string                           `json:"complete_ledgers"`
	ClosedLedger            *ServerLedgerState               `json:"closed_ledger,omitempty"`
	IOLatencyMS             uint                             `json:"io_latency_ms"`
	JQTransOverflow         string                           `json:"jq_trans_overflow"`
	LastClose               *ServerStateClose                `json:"last_close"`
	Load                    *ServerLoad                      `json:"load,omitempty"`
	LoadBase                int                              `json:"load_base"`
	LoadFactor              uint                             `json:"load_factor"`
	LoadFactorFeeEscelation uint                             `json:"load_factor_fee_escalation,omitempty"`
	LoadFactorFeeQueue      uint                             `json:"load_factor_fee_queue,omitempty"`
	LoadFactorFeeReference  uint                             `json:"load_factor_fee_reference,omitempty"`
	LoadFactorServer        uint                             `json:"load_factor_server,omitempty"`
	Peers                   uint                             `json:"peers,omitempty"`
	PubkeyNode              string                           `json:"pubkey_node"`
	PubkeyValidator         string                           `json:"pubkey_validator,omitempty"`
	Reporting               *ServerReporting                 `json:"reporting,omitempty"`
	ServerState             string                           `json:"server_state"`
	ServerStateDurationUS   string                           `json:"server_state_duration_us"`
	StateAccounting         map[string]ServerStateAccounting `json:"state_accounting"`
	Time                    string                           `json:"time"`
	Uptime                  uint                             `json:"uptime"`
	ValidatedLedger         *ServerLedgerState               `json:"validated_ledger,omitempty"`
	ValidationQuorum        uint                             `json:"validation_quorum"`
	ValidatorListExpires    string                           `json:"validator_list_expires,omitempty"`
}

type ServerLedgerState struct {
	BaseFee     uint   `json:"base_fee"`
	CloseTime   uint   `json:"close_time"`
	Hash        string `json:"hash"`
	ReserveBase uint   `json:"reserve_base"`
	ReserveInc  uint   `json:"reserve_inc"`
	Seq         uint   `json:"seq"`
}

type ServerStateClose struct {
	ConvergeTime uint `json:"converge_time"`
	Proposers    uint `json:"proposers"`
}

type ServerStateAccounting struct {
	DurationUS  string `json:"duration_us"`
	Transitions string `json:"transitions"`
}
