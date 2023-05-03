package status

import "github.com/xyield/xrpl-go/model/client/common"

type ConsensusInfoResponse struct {
	Info ConsensusInfo `json:"info"`
}

type ConsensusInfo struct {
	Consensus         string              `json:"consensus,omitempty"`
	Acquired          map[string]string   `json:"acquired,omitempty"`
	CloseGranularity  int                 `json:"close_granularity,omitempty"`
	ClosePercent      int                 `json:"close_percent,omitempty"`
	CloseResolution   int                 `json:"close_resolution,omitempty"`
	CloseTimes        map[string]int      `json:"close_times,omitempty"`
	CurrentMs         int                 `json:"current_ms,omitempty"`
	HaveTimeConsensus bool                `json:"have_time_consensus"`
	LedgerSeq         common.LedgerIndex  `json:"ledger_seq,omitempty"`
	OurPosition       Position            `json:"our_position"`
	PeerPositions     map[string]Position `json:"peer_positions,omitempty"`
	Proposers         int                 `json:"proposers,omitempty"`
	State             string              `json:"state,omitempty"`
}

type Position struct {
	CloseTime       uint              `json:"close_time"`
	PeerId          string            `json:"peer_id"`
	PreviousLedger  common.LedgerHash `json:"previous_ledger"`
	ProposeSeq      int               `json:"propose_seq"`
	TransactionHash string            `json:"transaction_hash"`
}
