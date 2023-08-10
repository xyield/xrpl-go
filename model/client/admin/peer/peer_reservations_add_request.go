package peer

type PeerReservationAddRequest struct {
	PublicKey   string `json:"public_key"`
	Description string `json:"description,omitempty"`
}

func (*PeerReservationAddRequest) Method() string {
	return "peer_reservations_add"
}
