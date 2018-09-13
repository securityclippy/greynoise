package client

import (
	"net/http"
	"encoding/json"
	"bytes"
	"strconv"
	"fmt"
	"strings"
)

var (
	Research = "research"
	Tag = "tag"
	List = "list"
	Combination = "combination"
	TimeSeries = "time_series"
	Single = "single"
	IP = "ip"
	Raw = "raw"
	Scan = "scan"
	Stats = "stats"
	Top = "top"
	HTTP = "http"
	Path = "path"
	UserAgent = "useragent"
	Org = "org"
	ASN = "asn"
	RDNS = "rdns"
	Search = "search"
	Actors = "actors"
	CIDR = "cidr"
	Scanners = "scanners"
	Ja3 = "ja3"
	Fingerprint = "fingerprint"

)

func (c Client) ResearchTagList() (*http.Response, error) {
	return c.DoRequest("GET", c.BuildURL(nil, Research, Tag, List), nil)
}


// ResearchTagCombination combines include and exclude lists to return all IPs that match said criteria
func (c Client) ResearchTagCombination(includeTags, excludeTags []string, offset int) (*http.Response, error) {
	type body struct {
		Query []string `json:"query"`
	}

	b := body{}

	for _, t := range includeTags {
		s := fmt.Sprintf("+%s", t)
		b.Query = append(b.Query, s)
	}

	for _, t := range excludeTags {
		s := fmt.Sprintf("-%s", t)
		b.Query = append(b.Query, s)
	}
	js, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	logger.Infof("body: %s", string(js))
	return c.DoRequest("GET", c.BuildURL(nil, Research, Tag, Combination), bytes.NewReader(js))
}


// ResearchTagSingle returns the values for the api endpoint /research/tag/single.  tag and offset are passed as strings
func (c Client) ResearchTagSingle(tag string, offset int) (*http.Response, error)  {
	type body struct {
		Tag string `json:"tag"`
	}
	b := body{
		Tag: tag,
	}
	js, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	if offset > 0 {
		params := map[string]string{"offset": strconv.Itoa(offset)}
		return c.DoRequest("GET", c.BuildURL(params, Research, Tag, Single), bytes.NewReader(js))
	}
	return c.DoRequest("GET", c.BuildURL(nil, Research, Tag, Single), bytes.NewReader(js))
}

func (c Client) ResearchTimeSeriesScan(protocol string, port int) (*http.Response,  error) {
	return c.DoRequest("GET", c.BuildURL(nil, Research, TimeSeries, Scan, protocol, strconv.Itoa(port)),nil)
}

func (c Client) ResearchIP(ipAddress string) (*http.Response, error) {
	return c.DoRequest("GET", c.BuildURL(nil, Research,  IP, ipAddress), nil)
}

func (c Client) ResearchRaw(protocol string, port int) (*http.Response, error) {
	return c.DoRequest("GET", c.BuildURL(nil, Research, Raw, Scan, protocol, strconv.Itoa(port)), nil)
}

func (c Client) ResearchRawIP(ipAddress string) (*http.Response, error) {
	return c.DoRequest("GET", c.BuildURL(nil, Research, Raw, IP, ipAddress), nil)
}

func (c Client) ResearchStatsTopScan() (*http.Response, error) {
	return c.DoRequest("GET", c.BuildURL(nil, Research, Stats, Top, Scan), nil)
}

func (c Client) ResearchStatsTopHTTPPath() (*http.Response, error) {
	return c.DoRequest("GET", c.BuildURL(nil, Research, Stats, Top, HTTP, Path), nil)
}

func (c Client) ResearchStatsTopHTTPUserAgent() (*http.Response, error) {
	return c.DoRequest("GET", c.BuildURL(nil, Research, Stats, Top, HTTP, UserAgent), nil)
}

func (c Client) ResearchStatsTopOrg() (*http.Response, error) {
	return c.DoRequest("GET", c.BuildURL(nil, Research, Stats, Top, Org), nil)
}

func (c Client) ResearchStatsTopASN() (*http.Response, error) {
	return c.DoRequest("GET", c.BuildURL(nil, Research, Stats, Top, ASN), nil)
}

func (c Client) ResearchStatsTopRDNS() (*http.Response, error) {
	return c.DoRequest("GET", c.BuildURL(nil, Research, Stats, Top,  RDNS), nil)
}

func (c Client) ResearchSearchOrg(orgName string) (*http.Response, error) {
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
	return c.DoRequest("GET",  c.BuildURL(nil, Research, Search, Org), bytes.NewReader(js))
}

func (c Client) ResearchActors()  (*http.Response, error) {
	return c.DoRequest("GET", c.BuildURL(nil, Research, Actors), nil)
}

func (c Client) ResearchScannersCIDRBlock(cidrBlock string) (*http.Response, error) {
	s := strings.Split(cidrBlock, "/")
	block := s[0]
	bits := s[1]
	return c.DoRequest("GET", c.BuildURL(nil, Research, Scanners, CIDR, block, bits), nil)
}

func (c Client) ResearchScannersASN(asn string) (*http.Response, error) {
	return c.DoRequest("GET", c.BuildURL(nil, Research, Scanners, ASN, asn), nil)
}

func (c Client) ResearchJa3Fingerprint(fingerprint string) (*http.Response, error) {
	return c.DoRequest("GET", c.BuildURL(nil, Research, Ja3, Fingerprint, fingerprint), nil)
}

func (c Client) ResearchJa3IP(ipAddress string) (*http.Response, error) {
	return c.DoRequest("GET", c.BuildURL(nil, Research, Ja3, IP, ipAddress), nil)
}
