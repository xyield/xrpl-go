package status

type ValidatorListSitesResponse struct {
	ValidatorSites []ValidatorSite `json:"validator_sites"`
}

type ValidatorSite struct {
	LastRefreshStatus  string `json:"last_refresh_status"`
	LastRefreshTime    string `json:"last_refresh_time"`
	RefreshIntervalMin uint   `json:"refresh_interval_min"`
	URI                string `json:"uri"`
}
