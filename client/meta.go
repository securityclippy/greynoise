package client

import "net/http"

var (
	Meta = "meta"
	Ping = "ping"
)

func (c Client) MetaPing() (*http.Response, error) {
	return c.DoRequest("GET", c.BuildURL(nil, Meta, Ping), nil)
}
