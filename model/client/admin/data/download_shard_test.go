package data

import (
	"testing"

	"github.com/xyield/xrpl-go/test"
)

func TestDownloadShardsRequest(t *testing.T) {
	s := DownloadShardRequest{
		Shards: []ShardDescriptor{
			{
				Index: 1,
				URL:   "https://example.com/1.tar.lz4",
			},
			{
				Index: 2,
				URL:   "https://example.com/2.tar.lz4",
			},
			{
				Index: 5,
				URL:   "https://example.com/5.tar.lz4",
			},
		},
	}

	j := `{
	"shards": [
		{
			"index": 1,
			"url": "https://example.com/1.tar.lz4"
		},
		{
			"index": 2,
			"url": "https://example.com/2.tar.lz4"
		},
		{
			"index": 5,
			"url": "https://example.com/5.tar.lz4"
		}
	]
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestDownloadShardsResponse(t *testing.T) {
	s := DownloadShardResponse{
		Message: "downloading shards 1-3",
	}
	j := `{
	"message": "downloading shards 1-3"
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
