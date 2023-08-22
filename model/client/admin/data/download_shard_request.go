package data

import "github.com/xyield/xrpl-go/model/client/common"

type DownloadShardRequest struct {
	Shards []ShardDescriptor `json:"shards"`
}

type ShardDescriptor struct {
	Index common.LedgerIndex `json:"index"`
	URL   string             `json:"url"`
}

func (*DownloadShardRequest) Method() string {
	return "download_shard"
}

func (*DownloadShardRequest) Validate() error {
	return nil
}
