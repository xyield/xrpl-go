package client

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

type customHttpClient struct{}

func (c customHttpClient) Do(req *http.Request) (*http.Response, error) {
	return nil, nil
}

func TestConfigCreation(t *testing.T) {

	t.Run("Set config with valid port + ip", func(t *testing.T) {
		cfg, _ := NewJsonRpcConfig("http://s1.ripple.com:51234/")

		req, err := http.NewRequest(http.MethodPost, "http://s1.ripple.com:51234/", nil)

		req.Header = cfg.Headers
		require.Equal(t, "http://s1.ripple.com:51234/", cfg.Url)
		require.NoError(t, err)
	})
	t.Run("No port + IP provided", func(t *testing.T) {
		cfg, err := NewJsonRpcConfig("")

		require.Nil(t, cfg)
		require.EqualError(t, err, "empty port and IP provided")
	})
	t.Run("Format root path - add /", func(t *testing.T) {
		cfg, _ := NewJsonRpcConfig("http://s1.ripple.com:51234")

		req, err := http.NewRequest(http.MethodPost, "http://s1.ripple.com:51234/", nil)

		req.Header = cfg.Headers
		require.Equal(t, "http://s1.ripple.com:51234/", cfg.Url)
		require.NoError(t, err)
	})
	t.Run("Pass in custom HTTP client", func(t *testing.T) {

		c := customHttpClient{}
		cfg, _ := NewJsonRpcConfig("http://s1.ripple.com:51234", WithHttpClient(c))

		req, err := http.NewRequest(http.MethodPost, "http://s1.ripple.com:51234/", nil)
		headers := map[string][]string{
			"Content-Type": {"application/json"},
		}
		req.Header = cfg.Headers
		require.Equal(t, &JsonRpcConfig{HTTPClient: customHttpClient{}, Url: "http://s1.ripple.com:51234/", Headers: headers}, cfg)
		require.NoError(t, err)
	})
}
