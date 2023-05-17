package data

type NodeToShardRequest struct {
	Action string `json:"action"`
}

func (*NodeToShardRequest) Method() string {
	return "node_to_shard"
}
