package peer

type PeersResponse struct {
	Cluster Cluster `json:"cluster"`
	Peers   []Peer  `json:"peers"`
}

type Cluster struct {
	Tag string `json:"tag,omitempty"`
	Fee int    `json:"fee,omitempty"`
	Age int    `json:"age,omitempty"`
}

type Peer struct {
	Address         string  `json:"address"`
	Cluster         bool    `json:"cluster,omitempty"`
	Name            string  `json:"name,omitempty"`
	CompleteLedgers string  `json:"complete_ledgers"`
	Inbound         bool    `json:"inbound,omitempty"`
	Latency         int     `json:"latency"`
	Ledger          string  `json:"ledger"`
	Load            int     `json:"load"`
	Protocol        string  `json:"protocol,omitempty"`
	Metrics         Metrics `json:"metrics"`
	PublicKey       string  `json:"public_key,omitempty"`
	Sanity          string  `json:"sanity,omitempty"`
	Status          string  `json:"status,omitempty"`
	Uptime          uint    `json:"uptime"`
	Version         string  `json:"version,omitempty"`
}

type Metrics struct {
	AvgBpsRecv     string `json:"avg_bps_recv"`
	AvgBpsSent     string `json:"avg_bps_sent"`
	TotalBytesRecv string `json:"total_bytes_recv"`
	TotalBytesSent string `json:"total_bytes_sent"`
}
