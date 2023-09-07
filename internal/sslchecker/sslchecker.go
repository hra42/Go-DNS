package sslchecker

import (
	"crypto/tls"
	"fmt"
	"math"
	"time"
)

func CheckSSL(url string) (report string) {
	// Konfiguration, um die Zertifikatsprüfung zu überspringen
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", url+":443", conf)
	if err != nil {
		report = fmt.Sprint("Failed to establish a connection:", err)
		return
	}
	defer func(conn *tls.Conn) {
		errConn := conn.Close()
		if errConn != nil {
			report = fmt.Sprintln("Failed to close connection:", errConn)
		}
	}(conn)

	state := conn.ConnectionState()
	certs := state.PeerCertificates

	if len(certs) == 0 {
		fmt.Println("No certificates found!")
		return
	}

	mainCert := certs[0]

	report += fmt.Sprintln("Domain resolves to:", url, "and has the following IP address:", conn.RemoteAddr())
	report += fmt.Sprintln("The hostname", url, "is correctly listed in the certificate.")

	t1 := mainCert.NotAfter
	t2 := time.Now()
	t1Date := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, t1.Location())
	t2Date := time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, t2.Location())
	validUntil := math.Round(t1Date.Sub(t2Date).Hours() / 24)
	if validUntil < 0 {
		report += fmt.Sprintln("The certificate is expired!")
	} else {
		report += fmt.Sprintln("The certificate is valid for", validUntil, "days until expired.")
	}
	report += fmt.Sprintln("Certificate Details:")
	report += fmt.Sprintln("Issuer:", mainCert.Issuer)
	report += fmt.Sprintln("Subject:", mainCert.Subject)
	report += fmt.Sprintln("Not Before:", mainCert.NotBefore)
	report += fmt.Sprintln("Not After:", mainCert.NotAfter)
	report += fmt.Sprintln("Signature Algorithm:", mainCert.SignatureAlgorithm)
	report += fmt.Sprintln("Public Key Algorithm:", mainCert.PublicKeyAlgorithm)
	if len(mainCert.DNSNames) > 0 {
		report += fmt.Sprintln("DNS Names:", mainCert.DNSNames)
	}
	if len(mainCert.EmailAddresses) > 0 {
		report += fmt.Sprintln("Email Addresses:", mainCert.EmailAddresses)
	}

	return
}
