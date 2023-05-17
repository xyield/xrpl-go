package peer

import (
	"testing"

	"github.com/xyield/xrpl-go/test"
)

func TestPeersResponse(t *testing.T) {
	s := PeersResponse{
		Peers: []Peer{
			{
				Address:         "5.189.239.203:51235",
				CompleteLedgers: "51813132 - 51815132",
				Latency:         100,
				Ledger:          "99A1E29C9F235DCCBB087F85F11756BECA606A756C22AB826AB1F319C470C3E3",
				Load:            157,
				Metrics: Metrics{
					AvgBpsRecv:     "10255",
					AvgBpsSent:     "2015",
					TotalBytesRecv: "356809",
					TotalBytesSent: "74208",
				},
				PublicKey: "n94ht2A9aBoARRhk1rwypZNVXJDiMN4qzs1Bd5KsQaSnN3WVy8Tw",
				Uptime:    2,
				Version:   "rippled-1.4.0",
			},
			{
				Address:         "[::ffff:50.22.123.222]:51235",
				CompleteLedgers: "32570 - 51815131",
				Latency:         100,
				Ledger:          "99A1E29C9F235DCCBB087F85F11756BECA606A756C22AB826AB1F319C470C3E3",
				Load:            219,
				Metrics: Metrics{
					AvgBpsRecv:     "7223",
					AvgBpsSent:     "6742",
					TotalBytesRecv: "593148",
					TotalBytesSent: "204540",
				},
				PublicKey: "n9LbkoB9ReSbaA9SGL317fm6CvjLcFG8hGoierLYfwiCDsEXHcP3",
				Uptime:    3,
				Version:   "rippled-1.3.1",
			},
		},
	}
	j := `{
	"cluster": {},
	"peers": [
		{
			"address": "5.189.239.203:51235",
			"complete_ledgers": "51813132 - 51815132",
			"latency": 100,
			"ledger": "99A1E29C9F235DCCBB087F85F11756BECA606A756C22AB826AB1F319C470C3E3",
			"load": 157,
			"metrics": {
				"avg_bps_recv": "10255",
				"avg_bps_sent": "2015",
				"total_bytes_recv": "356809",
				"total_bytes_sent": "74208"
			},
			"public_key": "n94ht2A9aBoARRhk1rwypZNVXJDiMN4qzs1Bd5KsQaSnN3WVy8Tw",
			"uptime": 2,
			"version": "rippled-1.4.0"
		},
		{
			"address": "[::ffff:50.22.123.222]:51235",
			"complete_ledgers": "32570 - 51815131",
			"latency": 100,
			"ledger": "99A1E29C9F235DCCBB087F85F11756BECA606A756C22AB826AB1F319C470C3E3",
			"load": 219,
			"metrics": {
				"avg_bps_recv": "7223",
				"avg_bps_sent": "6742",
				"total_bytes_recv": "593148",
				"total_bytes_sent": "204540"
			},
			"public_key": "n9LbkoB9ReSbaA9SGL317fm6CvjLcFG8hGoierLYfwiCDsEXHcP3",
			"uptime": 3,
			"version": "rippled-1.3.1"
		}
	]
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
