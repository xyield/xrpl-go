package client

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

type JsonRpcConfig struct {
	HTTPClient HTTPClient
	Url        string
	Headers    map[string][]string
}

func NewJsonRpcConfigWithHttpClient(url string, c HTTPClient) (*JsonRpcConfig, error) {

	cfg, err := NewJsonRpcConfig(url)
	if err != nil {
		return nil, err
	}

	if c != nil {
		cfg.HTTPClient = c
	}

	return cfg, nil
}

func NewJsonRpcConfig(url string) (*JsonRpcConfig, error) {

	// validate a url has been passed in
	if len(url) == 0 {
		return nil, &EmptyUrlError{}
	}
	// add slash if doesn't already end with one
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}

	cfg := &JsonRpcConfig{
		HTTPClient: &http.Client{Timeout: time.Duration(1) * time.Second}, // default timeout value - allow custom timme out?
		Url:        url,
		Headers: map[string][]string{
			"Content-Type": {"application/json"},
		},
	}

	return cfg, nil
}

// Method to set http client used for testing
func (cfg *JsonRpcConfig) AddHttpClient(c HTTPClient) {
	cfg.HTTPClient = c
}
