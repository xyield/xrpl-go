package data

type CrawlShardsResponse struct {
	CompleteShards string       `json:"complete_shards,omitempty"`
	Peers          []PeerShards `json:"peers,omitempty"`
}

type PeerShards struct {
	CompleteShards   string `json:"complete_shards"`
	IncompleteShards string `json:"incomplete_shards,omitempty"`
	PublicKey        string `json:"public_key,omitempty"`
}
