package dns

import (
	"log"
	"net"

	"github.com/miekg/dns"
)

func getTXTRecords(domain string, dnsServer net.IP) []string {

	dnsClient := dns.Client{}
	dnsMessage := dns.Msg{}
	dnsMessage.SetQuestion(dns.Fqdn(domain), dns.TypeTXT)

	var txtRecords []string

	server := dnsServer.String() + ":53" // Standard DNS port number 53
	reply, _, err := dnsClient.Exchange(&dnsMessage, server)
	if err != nil {
		log.Printf("Failed to get TXT record from %s for %s: %v", dnsServer, domain, err)
		return txtRecords
	}

	for _, ans := range reply.Answer {
		if txt, ok := ans.(*dns.TXT); ok {
			txtRecords = append(txtRecords, txt.String())
		}
	}
	return txtRecords
}

func GetTXTReport(url string) (report string) {
	for dnsServerName, dnsServerIP := range getDNSServers() {
		MXRecords := getTXTRecords(url, dnsServerIP)

		report += "DNS server provider: " + dnsServerName + "\nIP: " + dnsServerIP.String() + "\n"
		report += "Domain: " + url + "\n"
		if len(MXRecords) == 0 {
			report += "No TXT records found for " + url + "\n"
		} else {
			for _, record := range MXRecords {
				report += "TXT record: " + record + "\n\n"
			}
		}
	}
	return
}
