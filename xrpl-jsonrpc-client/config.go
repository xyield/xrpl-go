package jsonrpcclient

import (
	"net/http"
	"strings"
	"time"
)

type EmptyUrlError struct {
}

func (e *EmptyUrlError) Error() string {
	return "Empty port and IP provided"
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Config struct {
	HTTPClient HTTPClient
	Url        string
	Headers    map[string][]string
}

func NewConfigWithHttpClient(url string, c HTTPClient) (*Config, error) {

	cfg, err := NewConfig(url)
	if err != nil {
		return nil, err
	}

	if c != nil {
		cfg.HTTPClient = c
	}

	return cfg, nil
}

func NewConfig(url string) (*Config, error) {

	// validate a url has been passed in
	if len(url) == 0 {
		return nil, &EmptyUrlError{}
	}
	// add slash if doesn't already end with one
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}

	cfg := &Config{
		HTTPClient: &http.Client{Timeout: time.Duration(1) * time.Second}, // default timeout value - allow custom timme out?
		Url:        url,
		Headers: map[string][]string{
			"Content-Type": {"application/json"},
		},
	}

	return cfg, nil
}

// Method to set http client used for testing
func (cfg *Config) AddHttpClient(c HTTPClient) {
	cfg.HTTPClient = c
}
