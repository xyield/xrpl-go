package peer

import "fmt"

type PeerReservationAddRequest struct {
	PublicKey   string `json:"public_key"`
	Description string `json:"description,omitempty"`
}

func (*PeerReservationAddRequest) Method() string {
	return "peer_reservations_add"
}

func (r *PeerReservationAddRequest) Validate() error {
	if r.PublicKey == "" {
		return fmt.Errorf("peer reservation add request: missing publickey")
	}
	return nil
}
