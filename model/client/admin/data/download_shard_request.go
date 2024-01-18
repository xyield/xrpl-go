package data

import (
	"fmt"

	"github.com/xyield/xrpl-go/model/client/common"
)

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

func (d *DownloadShardRequest) Validate() error {
	if len(d.Shards) == 0 {
		return fmt.Errorf("download shard request: no shard descriptors provided")
	}
	return nil
}
