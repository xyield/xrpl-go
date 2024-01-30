package status

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/test"
)

func TestValidatorListSitesResponse(t *testing.T) {
	s := ValidatorListSitesResponse{
		ValidatorSites: []ValidatorSite{
			{
				LastRefreshStatus:  "accepted",
				LastRefreshTime:    "2017-Oct-13 21:26:37",
				RefreshIntervalMin: 5,
				URI:                "http://127.0.0.1:51447/validators",
			},
		},
	}

	j := `{
	"validator_sites": [
		{
			"last_refresh_status": "accepted",
			"last_refresh_time": "2017-Oct-13 21:26:37",
			"refresh_interval_min": 5,
			"uri": "http://127.0.0.1:51447/validators"
		}
	]
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
