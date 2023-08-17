package main

import (
	"flag"
	"github.com/hra42/Go-DNS/internal/dns"
)

func main() {
	domain := flag.String("domain", "", "Domain to query")
	mx := flag.Bool("mx", false, "Get MX records for a domain")
	cname := flag.Bool("cname", false, "Get CNAME records for a domain")
	txt := flag.Bool("txt", false, "Get TXT records for a domain")
	all := flag.Bool("all", false, "Get all records for a domain")
	flag.Parse()

	dnsServers := dns.GetDNSServers()
	switch true {
	case *mx == true:
		dns.PrintMXRecords(*domain, dnsServers)
	case *all == true:
		dns.PrintMXRecords(*domain, dnsServers)
		dns.PrintCNameRecords(*domain, dnsServers)
		dns.PrintTXTRecords(*domain, dnsServers)
	case *cname == true:
		dns.PrintCNameRecords(*domain, dnsServers)
	case *txt == true:
		dns.PrintTXTRecords(*domain, dnsServers)
	default:
		flag.Usage()
	}
}
