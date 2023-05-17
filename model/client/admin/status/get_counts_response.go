package status

type GetCountsResponse struct {
	Transaction      int     `json:",omitempty"`
	Ledger           int     `json:",omitempty"`
	NodeObject       int     `json:",omitempty"`
	Uptime           string  `json:"uptime,omitempty"`
	LedgerHitRate    float32 `json:"ledger_hit_rate,omitempty"`
	NodeHitRate      float32 `json:"node_hit_rate,omitempty"`
	NodeReadBytes    int     `json:"node_read_bytes,omitempty"`
	NodeReadsHit     int     `json:"node_reads_hit,omitempty"`
	NodeReadsTotal   int     `json:"node_reads_total,omitempty"`
	NodeWrites       int     `json:"node_writes,omitempty"`
	NodeWrittenBytes int     `json:"node_written_bytes,omitempty"`
}
