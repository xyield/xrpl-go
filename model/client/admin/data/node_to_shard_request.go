package data

import "fmt"

type NodeToShardRequest struct {
	Action string `json:"action"`
}

func (*NodeToShardRequest) Method() string {
	return "node_to_shard"
}

func (r *NodeToShardRequest) Validate() error {
	if r.Action != "start" && r.Action != "stop" && r.Action != "status" {
		return fmt.Errorf("node to shard request: invalid action '%s'", r.Action)
	}

	return nil
}
