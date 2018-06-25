package client

import (
	"net/http"
	"github.com/gin-gonic/gin/json"
	"bytes"
)

var (
	Enterprise = "enterprise"
	Noise = "noise"
	Quick = "quick"
	Multi = "multi"
)


func (c Client) EnterpriseNoiseQuick(ipAddress string) (*http.Response, error) {
	return c.doRequest("GET", c.buildURL(nil, Enterprise, Noise, Quick, ipAddress), nil)
}

func (c Client) EnterpriseNoiseMultiQuick(ipAddresses []string) (*http.Response, error) {
	type body struct {
		IPS []string `json:"ips"`
	}
	b := body{
		IPS: ipAddresses,
	}
	js, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	return c.doRequest("GET", c.buildURL(nil, Enterprise, Noise, Multi, Quick), bytes.NewReader(js))
}