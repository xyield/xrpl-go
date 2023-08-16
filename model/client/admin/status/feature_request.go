package status

type FeatureRequest struct {
	Feature string `json:"feature,omitempty"`
	Vetoed  bool   `json:"vetoed,omitempty"`
}

func (*FeatureRequest) Method() string {
	return "feature"
}
