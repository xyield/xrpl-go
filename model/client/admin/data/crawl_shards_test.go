package data

import (
	"testing"

	"github.com/xyield/xrpl-go/test"
)

func TestCrawlShardsRequest(t *testing.T) {
	s := CrawlShardsRequest{
		PublicKey: true,
		Limit:     1,
	}

	j := `{
	"public_key": true,
	"limit": 1
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestCrawlShardsResponse(t *testing.T) {
	s := CrawlShardsResponse{
		CompleteShards: "1-2,5,8-9,584,1973,2358",
		Peers: []PeerShards{
			{
				CompleteShards: "1-2,8,47,371,464,554,653,857,1076,1402,1555,1708,1813,1867",
				PublicKey:      "n9LxFZiySnfDSvfh23N94UxsFkCjWyrchTeKHcYE6tJJQL5iejb2",
			},
			{
				CompleteShards: "8-9,584",
				PublicKey:      "n9MN5xwYqbrj64rtfZAXQy7Y3sNxXZJeLt7Lj61a9DYEZ4SE2tQQ",
			},
		},
	}

	j := `{
	"complete_shards": "1-2,5,8-9,584,1973,2358",
	"peers": [
		{
			"complete_shards": "1-2,8,47,371,464,554,653,857,1076,1402,1555,1708,1813,1867",
			"public_key": "n9LxFZiySnfDSvfh23N94UxsFkCjWyrchTeKHcYE6tJJQL5iejb2"
		},
		{
			"complete_shards": "8-9,584",
			"public_key": "n9MN5xwYqbrj64rtfZAXQy7Y3sNxXZJeLt7Lj61a9DYEZ4SE2tQQ"
		}
	]
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
