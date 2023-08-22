package status

type ConsensusInfoRequest struct {
}

func (*ConsensusInfoRequest) Method() string {
	return "consensus_info"
}

func (*ConsensusInfoRequest) Validate() error {
	return nil
}
