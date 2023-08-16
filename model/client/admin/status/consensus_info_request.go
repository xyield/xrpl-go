package status

type ConsensusInfoRequest struct {
}

func (*ConsensusInfoRequest) Method() string {
	return "consensus_info"
}
