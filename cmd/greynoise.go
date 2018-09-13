package main

import (
	"github.com/integrii/flaggy"
	"github.com/sirupsen/logrus"
	"github.com/securityclippy/greynoise/client"
	"github.com/securityclippy/greynoise/conf"
	"os"
	"bufio"
	"io/ioutil"
)

var (
	log = logrus.WithField("greynoise", "client")
)

func readFileLines(filePath string) ([]string) {
	lines := []string{}
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func getSubcommands(command *flaggy.Subcommand) ([]*flaggy.Subcommand) {
	subs := []*flaggy.Subcommand{}
	for _, c := range command.Subcommands {
		if c.Used {
			subs = append(subs, c)
		} }
	for _, s := range subs{
		subs = append(subs, getSubcommands(s)...)
	}
	return subs
}

func main() {
	config, err := conf.ReadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}
	greyNoise := client.NewClient(config)

	var readFile string
	flaggy.String(&readFile, "f", "file", "read input from file")

	bits := flaggy.NewSubcommand("bits")
	bitsVar := ""
	bits.AddPositionalValue(&bitsVar, "bits", 1, true, "bits")

	block := flaggy.NewSubcommand("block")
	blockVar := ""
	block.AddPositionalValue(&blockVar, "block", 1, true, "block")
	block.AttachSubcommand(bits, 2)
	cidr := flaggy.NewSubcommand("cidr")
	rdns := flaggy.NewSubcommand("rdns")
	asn := flaggy.NewSubcommand("asn")
	org := flaggy.NewSubcommand("org")
	top := flaggy.NewSubcommand("top")
	ja3 := flaggy.NewSubcommand("ja3")
	scanners := flaggy.NewSubcommand("scanners")
	search := flaggy.NewSubcommand("search")
	actors := flaggy.NewSubcommand("actors")
	scan := flaggy.NewSubcommand("scan")
	raw := flaggy.NewSubcommand("raw")
	stats := flaggy.NewSubcommand("stats")
	userAgent := flaggy.NewSubcommand("useragent")
	pathFlag := flaggy.NewSubcommand("path")
	list := flaggy.NewSubcommand("list")
	combination := flaggy.NewSubcommand("combination")
	single := flaggy.NewSubcommand("single")
	port := flaggy.NewSubcommand("port")
	portVar := ""
	port.AddPositionalValue(&portVar, "port", 1, true, "port")

	proto := flaggy.NewSubcommand("protocol")
	protoVar := ""
	proto.AddPositionalValue(&protoVar, "proto", 1, true, "protocol")
	proto.AttachSubcommand(port, 2)

	ip := flaggy.NewSubcommand("ip")
	ipVar := ""
	ip.AddPositionalValue(&ipVar, "ipaddress", 1, false, "ip address")

	httpFlag := flaggy.NewSubcommand("http")
	httpFlag.AttachSubcommand(pathFlag, 1)
	httpFlag.AttachSubcommand(userAgent, 1)

	tag := flaggy.NewSubcommand("tag")
	tag.AttachSubcommand(list, 1)
	tag.AttachSubcommand(combination, 1)
	tag.AttachSubcommand(single, 1)

	//research api endpoints
	research := flaggy.NewSubcommand("research")
	research.ShortName = "r"
	flaggy.AttachSubcommand(research, 1)

	//endpoint paths

	timeSeries := flaggy.NewSubcommand("timeseries")
	timeSeries.AttachSubcommand(scan, 1)
	timeSeries.AttachSubcommand(httpFlag, 1)

	scan.AttachSubcommand(proto, 1)

	raw.AttachSubcommand(scan, 1)
	raw.AttachSubcommand(httpFlag, 1)

	stats.AttachSubcommand(top, 1)

	top.AttachSubcommand(scan, 1)
	top.AttachSubcommand(httpFlag, 1)
	top.AttachSubcommand(org, 1)
	top.AttachSubcommand(asn, 1)
	top.AttachSubcommand(rdns, 1)

	search.AttachSubcommand(org, 1)

	scanners.AttachSubcommand(cidr, 1)

	cidr.AttachSubcommand(block, 1)
	//stats := flaggy.NewSubcommand("stats")

	//research/ip

	research.AttachSubcommand(timeSeries, 1)
	research.AttachSubcommand(tag, 1)
	research.AttachSubcommand(ip, 1)
	research.AttachSubcommand(raw,1)
	research.AttachSubcommand(stats, 1)
	research.AttachSubcommand(actors, 1)
	research.AttachSubcommand(search, 1)
	research.AttachSubcommand(scanners, 1)
	research.AttachSubcommand(ja3, 1)


	//scan.AttachSubcommand(proto, 1)

	buildCommands := func(c *flaggy.Subcommand) ([]string, []string) {
		subcommands := getSubcommands(c)
		data := []string{}
		vars := []string{}
		for _, s := range subcommands {
			data = append(data, s.Name)
			pf := s.PositionalFlags
			for _, p := range pf {
				vars = append(vars, *p.AssignmentVar)
			}
		}
		data = append([]string{c.Name}, data...)
		return data, vars
	}


	flaggy.Parse()

	if len(readFile) > 0 {
		lines := readFileLines(readFile)
		cmd, _ := buildCommands(research)
		for _, l := range lines {
			paths := []string{}
			for _, p := range cmd {
				paths = append(paths, p)
			}
			paths = append(paths, l)
			url := greyNoise.BuildURL(nil, paths...)
			log.Info(url)
			resp, err := greyNoise.DoRequest("GET", url, nil)
			if err != nil {
				log.Error(err)
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil{
				log.Error(err)
			}
			log.Info(string(body))
		}
	}


	//if research.Used {
		//url := greyNoise.BuildURL(nil, data...)
		//logrus.Info(url)
	//}

}
