package status

type GetCountsRequest struct {
	MinCount int `json:"min_count,omitempty"`
}

func (*GetCountsRequest) Method() string {
	return "get_counts"
}
