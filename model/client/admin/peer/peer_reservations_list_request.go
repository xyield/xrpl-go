package peer

type PeerReservationsListRequest struct {
}

func (*PeerReservationsListRequest) Method() string {
	return "peer_reservations_list"
}

func (*PeerReservationsListRequest) Validate() error {
	return nil
}
