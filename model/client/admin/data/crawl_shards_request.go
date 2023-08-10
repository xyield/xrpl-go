package data

type CrawlShardsRequest struct {
	PublicKey bool `json:"public_key,omitempty"`
	Limit     int  `json:"limit,omitempty"`
}

func (*CrawlShardsRequest) Method() string {
	return "crawl_shards"
}
