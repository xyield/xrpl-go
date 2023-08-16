package peer

type PeerReservationsListResponse struct {
	Reservations []*PeerReservation `json:"reservations"`
}
