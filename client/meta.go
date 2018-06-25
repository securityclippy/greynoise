package client

import "net/http"

var (
	Meta = "meta"
	Ping = "ping"
)

func (c Client) MetaPing() (*http.Response, error) {
	return c.doRequest("GET", c.buildURL(nil, Meta, Ping), nil)
}
