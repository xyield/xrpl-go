package status

import (
	"testing"

	"github.com/xyield/xrpl-go/test"
)

func TestFeatureResponse(t *testing.T) {
	s := FeatureResponse{
		Features: map[string]Feature{
			"42426C4D4F1009EE67080A9B7965B44656D7714D104A72F9B4369F97ABF044EE": {
				Enabled:   false,
				Name:      "FeeEscalation",
				Supported: true,
				Vetoed:    false,
			},
			"4C97EBA926031A7CF7D7B36FDE3ED66DDA5421192D63DE53FFB46E43B9DC8373": {
				Enabled:   false,
				Name:      "MultiSign",
				Supported: true,
				Vetoed:    false,
			},
			"6781F8368C4771B83E8B821D88F580202BCB4228075297B19E4FDC5233F1EFDC": {
				Enabled:   false,
				Name:      "TrustSetAuth",
				Supported: true,
				Vetoed:    false,
			},
			"C1B8D934087225F509BEB5A8EC24447854713EE447D277F69545ABFA0E0FD490": {
				Enabled:   false,
				Name:      "Tickets",
				Supported: true,
				Vetoed:    false,
			},
			"DA1BD556B42D85EA9C84066D028D355B52416734D3283F85E216EA5DA6DB7E13": {
				Enabled:   false,
				Name:      "SusPay",
				Supported: true,
				Vetoed:    false,
			},
		},
	}

	j := `{
	"features": {
		"42426C4D4F1009EE67080A9B7965B44656D7714D104A72F9B4369F97ABF044EE": {
			"enabled": false,
			"name": "FeeEscalation",
			"supported": true,
			"vetoed": false
		},
		"4C97EBA926031A7CF7D7B36FDE3ED66DDA5421192D63DE53FFB46E43B9DC8373": {
			"enabled": false,
			"name": "MultiSign",
			"supported": true,
			"vetoed": false
		},
		"6781F8368C4771B83E8B821D88F580202BCB4228075297B19E4FDC5233F1EFDC": {
			"enabled": false,
			"name": "TrustSetAuth",
			"supported": true,
			"vetoed": false
		},
		"C1B8D934087225F509BEB5A8EC24447854713EE447D277F69545ABFA0E0FD490": {
			"enabled": false,
			"name": "Tickets",
			"supported": true,
			"vetoed": false
		},
		"DA1BD556B42D85EA9C84066D028D355B52416734D3283F85E216EA5DA6DB7E13": {
			"enabled": false,
			"name": "SusPay",
			"supported": true,
			"vetoed": false
		}
	}
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
