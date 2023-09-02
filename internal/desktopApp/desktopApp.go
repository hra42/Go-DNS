package desktopApp

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/hra42/Go-DNS/internal/dns"
	"image/color"
	"log"
)

func RunDesktop() {
	a := app.New()
	w := a.NewWindow("Go-DNS")

	// set the icon
	resourceIconPng, err := fyne.LoadResourceFromURLString("https://gpt-files.postrausch.tech/go-dns.png")
	if err != nil {
		log.Fatal(err)
	}
	w.SetIcon(resourceIconPng)

	// Resize and center the window
	w.Resize(fyne.NewSize(1000, 1000))
	w.CenterOnScreen()

	// menu
	dnsRecords := fyne.NewMenu(
		"DNS Records",
		fyne.NewMenuItem("MX", func() {
			getMX(w)
		}),
		fyne.NewMenuItem("CNAME", func() {
			getCNameRecords(w)
		}),
		fyne.NewMenuItem("TXT", func() {
			getTXT(w)
		}),
	)

	themeSwitcher := fyne.NewMenu(
		"Theme",
		fyne.NewMenuItem("Light", func() {
			a.Settings().SetTheme(theme.LightTheme())
		}),
		fyne.NewMenuItem("Dark", func() {
			a.Settings().SetTheme(theme.DarkTheme())
		}),
	)

	mainMenu := fyne.NewMainMenu(dnsRecords, themeSwitcher)

	w.SetMainMenu(mainMenu)

	// set up the content
	// light blue
	welcomeTitleText := canvas.NewText("Welcome to Go-DNS",
		color.RGBA{R: 4, G: 118, B: 208, A: 255})
	welcomeTitleText.TextStyle.Bold = true
	welcomeTitleText.Alignment = fyne.TextAlignCenter
	welcomeTitleText.TextSize = 60.0
	welcomeTitle := container.NewCenter(welcomeTitleText)

	welcomeText := canvas.NewText("This is a tool to check DNS records for Microsoft 365.",
		color.RGBA{R: 4, G: 118, B: 208, A: 255})
	welcomeText.Alignment = fyne.TextAlignCenter
	welcomeText.TextSize = 20.0
	welcome := container.NewCenter(welcomeText)

	w.SetContent(container.NewVBox(
		welcomeTitle,
		welcome,
	))

	w.ShowAndRun()
}

func getMX(w fyne.Window) {
	// set up the content
	w.SetTitle("Go-DNS - MX Records")
	Title := canvas.NewText("MX Records",
		color.RGBA{R: 4, G: 118, B: 208, A: 255})
	Title.TextStyle.Bold = true
	Title.Alignment = fyne.TextAlignCenter
	Title.TextSize = 60.0
	inputDomain := widget.NewEntry()
	inputDomain.SetPlaceHolder("Enter the domain you want to check")
	placeholder := widget.NewLabel("This is a placeholder")
	placeholder.Hide()

	box := container.NewVBox(
		Title,
		inputDomain,
		placeholder,
		widget.NewButton("Check", func() {
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
	)

	scroll := container.NewScroll(box)
	w.SetContent(scroll)
}

func getCNameRecords(w fyne.Window) {
	w.SetTitle("Go-DNS - CNAME Records")
	Title := canvas.NewText("CNAME Records",
		color.RGBA{R: 4, G: 118, B: 208, A: 255})
	Title.TextStyle.Bold = true
	Title.Alignment = fyne.TextAlignCenter
	Title.TextSize = 60.0
	// set up the content
	inputDomain := widget.NewEntry()
	inputDomain.SetPlaceHolder("Enter the domain you want to check")
	warning := canvas.NewText("Attention! This can take a long time", color.RGBA{R: 255, G: 0, B: 0, A: 255})
	warning.TextStyle.Bold = true
	placeholder := widget.NewLabel("This is a placeholder")
	placeholder.Hide()

	subdomains := []string{"autodiscover", "lyncdiscover", "selector1._domainkey", "selector2._domainkey"}

	box := container.NewVBox(
		Title,
		warning,
		inputDomain,
		placeholder,
		widget.NewButton("Check", func() {
			report := "" // Initialize the report string here
			for _, subdomain := range subdomains {
				FullDomain := fmt.Sprintf("%s.%s", subdomain, inputDomain.Text)
				report += fmt.Sprintf("Domain: %s\n", FullDomain)

				for dnsServerName, dnsServerIP := range dns.GetDNSServers() {
					CnameRecordsFullDomain := dns.GetCNameRecords(FullDomain, dnsServerIP)

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
			warning.Hide()
			placeholder.SetText(report)
			placeholder.Show()
		}),
	)

	scroll := container.NewScroll(box)
	w.SetContent(scroll)
}

func getTXT(w fyne.Window) {
	w.SetTitle("Go-DNS - TXT Records")
	Title := canvas.NewText("TXT Records",
		color.RGBA{R: 4, G: 118, B: 208, A: 255})
	Title.TextStyle.Bold = true
	Title.Alignment = fyne.TextAlignCenter
	Title.TextSize = 60.0
	// set up the content
	inputDomain := widget.NewEntry()
	inputDomain.SetPlaceHolder("Enter the domain you want to check")
	placeholder := widget.NewLabel("This is a placeholder")
	placeholder.Hide()

	box := container.NewVBox(
		Title,
		inputDomain,
		placeholder,
		widget.NewButton("Check", func() {
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
	)

	scroll := container.NewScroll(box)
	w.SetContent(scroll)
}
