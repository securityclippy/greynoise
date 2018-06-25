# Unofficial Greynoise.io go client

This is a basic go client for the greynoise v2 api.  Most endpoints are functional,
but several are still to be implemented on they greynoise platform.

to get started:

```bash
go get -u github.com/securityclippy/greynoise
```

`main.go`

holds some basic examples of the endpoints and their associated functions

### basic usage:

```
package main

import (
	"github.com/securityclippy/greynoise/client"
	"github.com/securityclippy/greynoise/conf"
)

func main() {
	config, err := conf.ReadConfig("config.json")
	if err != nil {
		logger.Fatal(err)
	}
	greyNoise := client.NewClient(config)

	// call endpoint call
	resp, err := greyNoise.MetaPing()
    if err != nil {
        log.Fatal(err)
    }

    // parse response body from *http.response returned by endpoints
    body, err := greyNoise.ParseResponse(resp)
    // print out response body
    log.Info(string(body))
```


### endpoint implementation Status

- [x] /v2/meta/ping
- [x] /v2/research/time_series/scan/{protoco}/{port} __NOTE: currently not responding__
- [ ] /v2/research/time_series/http/path - (Greynoise - Not Yet Implemented)
- [ ] /v2/research/time_series/http/useragent - (Greynoise - Not Yet Implemented)
- [x] /v2/research/tag/list
- [x] /v2/research/tag/combination
- [x] /v2/research/tag/single
- [x] /v2/research/ip/{ip}
- [x] /v2/research/raw/scan/{protocol}/{port} - (Greynoise - Not Yet Implemented)
- [x] /v2/research/raw/ip/{ip} - (Greynoise - Not Yet Implemented)
- [ ] /v2/research/raw/http/path (Greynoise - Not Yet Implemented)
- [ ] /v2/research/raw/http/useragent (Greynoise - Not Yet Implemented)
- [x] /v2/research/stats/top/scan
- [x] /v2/research/stats/top/http/path
- [x] /v2/research/stats/top/http/useragent
- [x] /v2/research/stats/top/org
- [x] /v2/research/stats/top/asn
- [x] /v2/research/stats/top/rdns
- [x] /v2/research/search/org __NOTE: Timing out for all orgs__
- [ ] /v2/research/search/org/historical
- [x] /v2/research/actors __NOTE: Timing out__
- [x] /v2/infections/cidr/{block}/{bits}
- [x] /v2/infections/asn/{asn} __NOTE: timing out__
- [x] /v2/infections/search/org __NOTE: timing out__
- [x] /v2/research/scanners/cidr/{block}/{bits}
- [x] /v2/research/scanners/asn/{asn} __NOTE: timing out__
- [x] /v2/research/ja3/fingerprint/{fingerprint}
- [x] /v2/research/ja3/ip/{ip}
- [x] /v2/enterprise/noise/quick/{ip}
- [x] /v2/enterprise/noise/multi/quick

