package desktopApp

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/hra42/Go-DNS/internal/dns"
	"image/color"
	"log"
)

func RunDesktop() {
	// set up the app
	a := app.New()
	w := a.NewWindow("Go-DNS")

	// set the icon
	resourceIconPng, err := fyne.LoadResourceFromURLString("https://gpt-files.postrausch.tech/go-dns.png")
	if err != nil {
		log.Fatal(err)
	}
	w.SetIcon(resourceIconPng)

	// set the window size
	w.Resize(fyne.NewSize(400, 400))

	// main menu
	mainMenu := fyne.NewMainMenu(
		fyne.NewMenu("Menu",
			fyne.NewMenuItem("MX Einträge", func() {
				getMX(w)
			}),
			fyne.NewMenuItem("TXT Einträge", func() {
				getTXT(w)
			}),
			fyne.NewMenuItem("CNAME Einträge", func() {
				getCNameRecords(w)
			}),
		),
	)
	w.SetMainMenu(mainMenu)

	w.ShowAndRun()
}

func getMX(w fyne.Window) {
	// set up the content
	inputDomain := widget.NewEntry()
	inputDomain.SetPlaceHolder("Gib die Domäne die du prüfen möchtest ein!")
	placeholder := widget.NewLabel("This is a placeholder")
	placeholder.Hide()

	w.SetContent(container.NewVBox(
		inputDomain,
		placeholder,
		widget.NewButton("Prüfen", func() {
			report := ""
			for dnsServerName, dnsServerIP := range dns.GetDNSServers() {
				MXRecords := dns.GetMXRecords(inputDomain.Text, dnsServerIP)

				report += "DNS server provider: " + dnsServerName + "\nIP: " + dnsServerIP.String() + "\n"
				report += "Domain: " + inputDomain.Text + "\n"
				if len(MXRecords) == 0 {
					report += "No MX records found for " + inputDomain.Text + "\n"
				} else {
					for _, record := range MXRecords {
						report += "MX record: " + record + "\n\n"
					}
				}
			}
			placeholder.SetText(report)
			placeholder.Show()
		}),
	))
}

func getCNameRecords(w fyne.Window) {
	// set up the content
	inputDomain := widget.NewEntry()
	inputDomain.SetPlaceHolder("Gib die Domäne die du prüfen möchtest ein!")
	warning := canvas.NewText("Achtung: Dieser Test kann sehr lange dauern!", color.RGBA{R: 255, G: 0, B: 0, A: 255})
	warning.TextStyle.Bold = true
	placeholder := widget.NewLabel("This is a placeholder")
	placeholder.Hide()

	subdomains := []string{"autodiscover", "lyncdiscover", "selector1._domainkey", "selector2._domainkey"}

	w.SetContent(container.NewVBox(
		warning,
		inputDomain,
		placeholder,
		widget.NewButton("Prüfen", func() {
			report := "" // Initialize the report string here
			for _, subdomain := range subdomains {
				FullDomain := fmt.Sprintf("%s.%s", subdomain, inputDomain.Text)
				report += fmt.Sprintf("Domain: %s\n", FullDomain)

				for dnsServerName, dnsServerIP := range dns.GetDNSServers() {
					CnameRecordsFullDomain := dns.GetCNameRecords(FullDomain, dnsServerIP)

					report += fmt.Sprintf("DNS server provider: %s\nIP: %s\n", dnsServerName, dnsServerIP)

					if len(CnameRecordsFullDomain) == 0 {
						report += fmt.Sprintf("No CNAME records found for %s\n", FullDomain)
					} else {
						for _, record := range CnameRecordsFullDomain {
							if record == "" {
								report += fmt.Sprintf("No CNAME records found for %s\n", FullDomain)
								break
							} else {
								report += fmt.Sprintf("CNAME record for %v: %v\n", FullDomain, record)
							}
						}
					}
				}
			}
			warning.Hide()
			placeholder.SetText(report)
			placeholder.Show()
		}),
	))
}

func getTXT(w fyne.Window) {
	// set up the content
	inputDomain := widget.NewEntry()
	inputDomain.SetPlaceHolder("Gib die Domäne die du prüfen möchtest ein!")
	placeholder := widget.NewLabel("This is a placeholder")
	placeholder.Hide()

	w.SetContent(container.NewVBox(
		inputDomain,
		placeholder,
		widget.NewButton("Prüfen", func() {
			report := ""
			for dnsServerName, dnsServerIP := range dns.GetDNSServers() {
				MXRecords := dns.GetTXTRecords(inputDomain.Text, dnsServerIP)

				report += "DNS server provider: " + dnsServerName + "\nIP: " + dnsServerIP.String() + "\n"
				report += "Domain: " + inputDomain.Text + "\n"
				if len(MXRecords) == 0 {
					report += "No TXT records found for " + inputDomain.Text + "\n"
				} else {
					for _, record := range MXRecords {
						report += "TXT record: " + record + "\n\n"
					}
				}
			}
			placeholder.SetText(report)
			placeholder.Show()
		}),
	))
}
