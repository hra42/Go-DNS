package dns

import (
	"fmt"
	"github.com/miekg/dns"
	"log"
	"net"
)

func getCNameRecords(domain string, dnsServer net.IP) []string {
	dnsClient := dns.Client{}
	dnsMessage := dns.Msg{}
	dnsMessage.SetQuestion(dns.Fqdn(domain), dns.TypeCNAME)

	var cnameRecords []string

	server := dnsServer.String() + ":53"
	reply, _, err := dnsClient.Exchange(&dnsMessage, server)
	if err != nil {
		log.Printf("Failed to get CNAME record from %s for %s: %v", dnsServer, domain, err)
		return cnameRecords
	}

	for _, ans := range reply.Answer {
		if cname, ok := ans.(*dns.CNAME); ok {
			cnameRecords = append(cnameRecords, cname.Target)
		}
	}
	return cnameRecords
}

func GetCnameReport(url string) (report string) {
	subdomains := []string{"autodiscover", "lyncdiscover", "selector1._domainkey", "selector2._domainkey"}
	for _, subdomain := range subdomains {
		FullDomain := fmt.Sprintf("%s.%s", subdomain, url)
		report += fmt.Sprintf("Domain: %s\n", FullDomain)

		for dnsServerName, dnsServerIP := range getDNSServers() {
			CnameRecordsFullDomain := getCNameRecords(FullDomain, dnsServerIP)

			report += fmt.Sprintf("DNS server provider: %s\nIP: %s\n", dnsServerName, dnsServerIP)

			if len(CnameRecordsFullDomain) == 0 {
				report += fmt.Sprintf("No CNAME records found for %s\n\n", FullDomain)
			} else {
				for _, record := range CnameRecordsFullDomain {
					if record == "" {
						report += fmt.Sprintf("No CNAME records found for %s\n\n", FullDomain)
						break
					} else {
						report += fmt.Sprintf("CNAME record for %v: %v\n\n", FullDomain, record)
					}
				}
			}
		}
	}
	return
}
