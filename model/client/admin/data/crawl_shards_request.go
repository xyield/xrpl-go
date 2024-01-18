package data

import "fmt"

type CrawlShardsRequest struct {
	PublicKey bool `json:"public_key,omitempty"`
	Limit     int  `json:"limit,omitempty"`
}

func (*CrawlShardsRequest) Method() string {
	return "crawl_shards"
}

func (r *CrawlShardsRequest) Validate() error {
	if r.Limit < 0 || r.Limit > 3 {
		return fmt.Errorf("crawl shards request: invalid limit, must be 0 <= limit <= 3")
	}

	return nil
}
