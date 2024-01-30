package peer

type PeersRequest struct {
}

func (*PeersRequest) Method() string {
	return "peers"
}

func (*PeersRequest) Validate() error {
	return nil
}
