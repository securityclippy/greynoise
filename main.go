package main

import (
	"github.com/securityclippy/greynoise/client"
	"github.com/securityclippy/greynoise/conf"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	logger = logrus.WithField("client", "main")
)

func main() {
	//do stuff
	config, err := conf.ReadConfig("config.json")
	if err != nil {
		logger.Fatal(err)
	}
	gn := client.NewClient(config)
	gn.NetClient.Timeout = 200 * time.Second
	//resp, err := gn.MetaPing()
	//resp, err := gn.ResearchTagCombination([]string{"Mirai", "Telnet Scanner"}, nil, 0)
	//resp, err := gn.ResearchRaw("tcp", 23)
	//resp, err := gn.ResearchRawIP("5.221.27.224")
	resp, err := gn.ResearchTimeSeriesScan("tcp", 23)
	//resp, err := gn.ResearchStatsTopScan()
	//resp, err := gn.ResearchStatsTopHTTPPath()
	//resp, err := gn.ResearchStatsTopHTTPUserAgent()
	//resp, err := gn.ResearchStatsTopOrg()
	//resp, err := gn.ResearchStatsTopASN()
	//resp, err := gn.ResearchStatsTopRDNS()
	//resp, err := gn.ResearchSearchOrg("microsoft")
	//resp, err := gn.ResearchActors()
	//resp, err := gn.InfectionsCIDR("5.239.241.48/16")
	//resp, err := gn.InfectionsASN("AS45899")
	//resp, err := gn.InfectionsSearchOrg("Telecommunication Company of Tehran")
	//resp, err := gn.ResearchScannersCIDRBlock("5.239.241.48/16")
	//resp, err := gn.ResearchScannersASN("AS51119")
	//resp, err := gn.ResearchJa3IP("5.239.17.110")
	//resp, err := gn.EnterpriseNoiseQuick("5.239.241.48")
	//resp, err := gn.EnterpriseNoiseMultiQuick([]string{"5.239.241.48", "5.239.17.110"})
	if err != nil {
		logger.Fatal(err)
	}
	body, err := gn.ParseResponse(resp)
	logger.Info(string(body))
}
