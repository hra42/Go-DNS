package dns

import (
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
