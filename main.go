package main

import (
	"github.com/securityclippy/greynoise/client"
	"github.com/securityclippy/greynoise/conf"
	"github.com/sirupsen/logrus"
)

var (
	logger = logrus.WithField("GreyNoise", "main")
)

func init() {
	logger.Level = logrus.InfoLevel
}

func main() {
	//do stuff
	config, err := conf.ReadConfig("config.json")
	if err != nil {
		logger.Fatal(err)
	}
	greyNoise := client.NewClient(config)

	// testing endpoint call
	//resp, err := greyNoise.MetaPing()
	//resp, err := greyNoise.ResearchTagCombination([]string{"Mirai", "Telnet Scanner"}, nil, 0)
	//resp, err := greyNoise.ResearchRaw("tcp", 23)
	//resp, err := greyNoise.ResearchRawIP("5.221.27.224")
	resp, err := greyNoise.ResearchTimeSeriesScan("tcp", 23)
	//resp, err := greyNoise.ResearchStatsTopScan()
	//resp, err := greyNoise.ResearchStatsTopHTTPPath()
	//resp, err := greyNoise.ResearchStatsTopHTTPUserAgent()
	//resp, err := greyNoise.ResearchStatsTopOrg()
	//resp, err := greyNoise.ResearchStatsTopASN()
	//resp, err := greyNoise.ResearchStatsTopRDNS()
	//resp, err := greyNoise.ResearchSearchOrg("microsoft")
	//resp, err := greyNoise.ResearchActors()
	//resp, err := greyNoise.InfectionsCIDR("5.239.241.48/16")
	//resp, err := greyNoise.InfectionsASN("AS45899")
	//resp, err := greyNoise.InfectionsSearchOrg("Telecommunication Company of Tehran")
	//resp, err := greyNoise.ResearchScannersCIDRBlock("5.239.241.48/16")
	//resp, err := greyNoise.ResearchScannersASN("AS51119")
	//resp, err := greyNoise.ResearchJa3IP("5.239.17.110")
	//resp, err := greyNoise.EnterpriseNoiseQuick("5.239.241.48")
	//resp, err := greyNoise.EnterpriseNoiseMultiQuick([]string{"5.239.241.48", "5.239.17.110"})
	if err != nil {
		logger.Fatal(err)
	}

	// parse response body from *http.response returned by endpoints
	body, err := greyNoise.ParseResponse(resp)
	logger.Info(string(body))
}
