package peer

import (
	"testing"

	"github.com/xyield/xrpl-go/test"
)

func TestPeerReservationsListResponse(t *testing.T) {
	s := PeerReservationsListResponse{
		Reservations: []*PeerReservation{
			{
				Node:        "n9Jt8awsPzWLjBCNKVEEDQnw4bQEPjezfcQ4gttD1UzbLT1FoG99",
				Description: "Ripple s1 server 'WOOL'",
			},
			{
				Node: "n9MZRo92mzYjjsa5XcqnPC7GFYAnENo9VfJzKmpcS9EFZvw5fgwz",
			},
		},
	}

	j := `{
	"reservations": [
		{
			"node": "n9Jt8awsPzWLjBCNKVEEDQnw4bQEPjezfcQ4gttD1UzbLT1FoG99",
			"description": "Ripple s1 server 'WOOL'"
		},
		{
			"node": "n9MZRo92mzYjjsa5XcqnPC7GFYAnENo9VfJzKmpcS9EFZvw5fgwz"
		}
	]
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
