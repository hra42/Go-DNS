package dns

import (
	"fmt"
	"github.com/miekg/dns"
	"log"
	"net"
)

func getDNSServers() map[string]net.IP {
	dnsServers := map[string]net.IP{}
	dnsServers["Google"] = net.ParseIP("8.8.8.8")
	dnsServers["Cloudflare"] = net.ParseIP("1.1.1.1")
	dnsServers["Quad9"] = net.ParseIP("9.9.9.9")

	return dnsServers
}

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
