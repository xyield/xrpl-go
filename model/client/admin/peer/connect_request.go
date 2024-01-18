package peer

import "fmt"

type ConnectRequest struct {
	IP   string `json:"ip"`
	Port int    `json:"port,omitempty"`
}

func (*ConnectRequest) Method() string {
	return "connect"
}

func (c *ConnectRequest) Validate() error {
	if c.IP == "" {
		return fmt.Errorf("connect request: missing ip")
	}
	return nil
}
