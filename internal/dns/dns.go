package dns

import (
	"net"
)

func getDNSServers() map[string]net.IP {
	dnsServers := map[string]net.IP{}
	dnsServers["Google"] = net.ParseIP("8.8.8.8")
	dnsServers["Cloudflare"] = net.ParseIP("1.1.1.1")
	dnsServers["Quad9"] = net.ParseIP("9.9.9.9")

	return dnsServers
}
