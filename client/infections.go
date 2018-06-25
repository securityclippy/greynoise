package client

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin/json"
	"bytes"
)

var  (
	Infections = "infections"
)

func (c Client) InfectionsCIDR(cidrBlock string) (*http.Response, error) {
	s := strings.Split(cidrBlock, "/")
	block := s[0]
	bits := s[1]
	return c.doRequest("GET", c.buildURL(nil, Infections, CIDR, block, bits), nil)
}

func (c Client) InfectionsASN(asn string) (*http.Response, error) {
	return c.doRequest("GET", c.buildURL(nil, Infections, ASN, asn), nil)
}

func (c Client) InfectionsSearchOrg(orgName string) (*http.Response, error) {
	type body struct {
		Search string `json:"search"`
	}
	b := body{
		Search: orgName,
	}
	js, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	return c.doRequest("GET", c.buildURL(nil, Infections, Search, Org), bytes.NewReader(js))
}



