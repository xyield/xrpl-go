package status

type FetchInfoResponse struct {
	Info map[string]FetchInfo `json:"info"`
}

type FetchInfo struct {
	Hash              string   `json:"hash"`
	HaveHeader        bool     `json:"have_header"`
	HaveTransactions  bool     `json:"have_transactions"`
	NeededStateHashes []string `json:"needed_state_hashes"`
	Peers             int      `json:"peers"`
	Timeouts          int      `json:"timeouts"`
}
