package dns

import (
	"fmt"
	"github.com/miekg/dns"
	"log"
	"net"
)

func GetDNSServers() map[string]net.IP {
	dnsServers := map[string]net.IP{}
	dnsServers["Google"] = net.ParseIP("8.8.8.8")
	dnsServers["Cloudflare"] = net.ParseIP("1.1.1.1")
	dnsServers["Quad9"] = net.ParseIP("9.9.9.9")

	return dnsServers
}

func GetMXRecords(domain string, dnsServer net.IP) []string {

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

func GetCNameRecords(domain string, dnsServer net.IP) []string {
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

func GetTXTRecords(domain string, dnsServer net.IP) []string {

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

func PrintMXRecords(domain string) {
	fmt.Printf("Domain: %s\n", domain)
	for dnsServerName, dnsServerIP := range GetDNSServers() {
		MXRecords := GetMXRecords(domain, dnsServerIP)

		fmt.Printf("DNS server provider: %s\nIP: %s\n", dnsServerName, dnsServerIP)
		if len(MXRecords) == 0 {
			fmt.Printf("No MX records found for %s\n", domain)
		} else {
			for _, record := range MXRecords {
				fmt.Printf("MX for record: %v\n", record)
			}
		}
	}
	fmt.Println("------------------------------------")
}

func PrintTXTRecords(domain string) {
	fmt.Printf("Domain: %s\n", domain)
	for dnsServerName, dnsServerIP := range GetDNSServers() {
		TXTRecords := GetTXTRecords(domain, dnsServerIP)
		fmt.Printf("DNS server provider: %s\nIP: %s\n", dnsServerName, dnsServerIP)
		if len(TXTRecords) == 0 {
			fmt.Printf("No TXT records found for %s\n", domain)
		} else {
			for _, record := range TXTRecords {
				if record == "" {
					fmt.Printf("No TXT records found for %s\n", domain)
					break
				} else {
					fmt.Printf("TXT for record: %v\n", record)
				}
			}
		}
	}
	fmt.Println("------------------------------------")
}

func PrintCNameRecords(domain string) {
	subdomains := []string{"autodiscover", "lyncdiscover", "selector1._domainkey", "selector2._domainkey"}
	for _, subdomain := range subdomains {
		FullDomain := fmt.Sprintf("%s.%s", subdomain, domain)
		fmt.Printf("Domain: %s\n", FullDomain)
		for dnsServerName, dnsServerIP := range GetDNSServers() {
			CnameRecordsFullDomain := GetCNameRecords(FullDomain, dnsServerIP)
			fmt.Printf("DNS server provider: %s\nIP: %s\n", dnsServerName, dnsServerIP)
			if len(CnameRecordsFullDomain) == 0 {
				fmt.Printf("No CNAME records found for %s\n", FullDomain)
			} else {
				for _, record := range CnameRecordsFullDomain {
					if record == "" {
						fmt.Printf("No CNAME records found for %s\n", FullDomain)
						break
					} else {
						fmt.Printf("CNAME record for %v: %v\n", FullDomain, record)
					}
				}
			}
		}
		fmt.Println("------------------------------------")
	}
}
