package main

import (
	"flag"
	"github.com/hra42/Go-DNS/internal/desktopApp"
	"github.com/hra42/Go-DNS/internal/dns"
)

func main() {
	domain := flag.String("domain", "", "Domain to query")
	mx := flag.Bool("mx", false, "Get MX records for a domain")
	cname := flag.Bool("cname", false, "Get CNAME records for a domain")
	txt := flag.Bool("txt", false, "Get TXT records for a domain")
	all := flag.Bool("all", false, "Get all records for a domain")
	flag.Parse()

	switch true {
	case *mx == true:
		dns.PrintMXRecords(*domain)
	case *all == true:
		dns.PrintMXRecords(*domain)
		dns.PrintCNameRecords(*domain)
		dns.PrintTXTRecords(*domain)
	case *cname == true:
		dns.PrintCNameRecords(*domain)
	case *txt == true:
		dns.PrintTXTRecords(*domain)
	default:
		desktopApp.RunDesktop()
	}
}
