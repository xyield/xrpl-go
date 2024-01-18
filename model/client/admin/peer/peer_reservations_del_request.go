package peer

import "fmt"

type PeerReservationDelRequest struct {
	PublicKey string `json:"public_key"`
}

func (*PeerReservationDelRequest) Method() string {
	return "peer_reservations_del"
}

func (r *PeerReservationDelRequest) Validate() error {
	if r.PublicKey == "" {
		return fmt.Errorf("peer reservation del request: missing publickey")
	}
	return nil
}
