package status

type GetCountsResponse struct {
	Transaction      int
	Ledger           int
	NodeObject       int
	Uptime           string  `json:"uptime"`
	LedgerHitRate    float32 `json:"ledger_hit_rate"`
	NodeHitRate      float32 `json:"node_hit_rate"`
	NodeReadBytes    int     `json:"node_read_bytes"`
	NodeReadsHit     int     `json:"node_reads_hit"`
	NodeReadsTotal   int     `json:"node_reads_total"`
	NodeWrites       int     `json:"node_writes"`
	NodeWrittenBytes int     `json:"node_written_bytes"`
}
