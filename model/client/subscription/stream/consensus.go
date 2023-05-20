package stream

type ConsensusStream struct {
	Type      StreamType `json:"type"`
	Consensus string     `json:"consensus"`
}
