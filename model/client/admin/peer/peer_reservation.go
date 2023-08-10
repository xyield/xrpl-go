package peer

type PeerReservation struct {
	Node        string `json:"node"`
	Description string `json:"description,omitempty"`
}
