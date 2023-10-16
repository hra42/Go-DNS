package dns

import (
	"log"
	"net"

	"github.com/miekg/dns"
)

func getMXRecords(domain string, dnsServer net.IP) []string {

	dnsClient := dns.Client{}
	dnsMessage := dns.Msg{}
	dnsMessage.SetQuestion(dns.Fqdn(domain), dns.TypeMX)

	var mxRecords []string

	server := dnsServer.String() + ":53" // Standard DNS port number 53
	reply, _, err := dnsClient.Exchange(&dnsMessage, server)
	if err != nil {
		log.Printf("Failed to get MX record from %s for %s: %v", dnsServer, domain, err)
		return mxRecords
	}

	for _, ans := range reply.Answer {
		if mx, ok := ans.(*dns.MX); ok {
			mxRecords = append(mxRecords, mx.Mx)
		}
	}
	return mxRecords
}

func GetMXReport(url string) (report string) {
	for dnsServerName, dnsServerIP := range getDNSServers() {
		MXRecords := getMXRecords(url, dnsServerIP)

		report += "DNS server provider: " + dnsServerName + "\nIP: " + dnsServerIP.String() + "\n"
		report += "Domain: " + url + "\n"
		if len(MXRecords) == 0 {
			report += "No MX records found for " + url + "\n"
		} else {
			for _, record := range MXRecords {
				report += "MX record: " + record + "\n\n"
			}
		}
	}
	return
}
