package peer

type PeerReservationDelRequest struct {
	PublicKey string `json:"public_key"`
}

func (*PeerReservationDelRequest) Method() string {
	return "peer_reservations_del"
}
