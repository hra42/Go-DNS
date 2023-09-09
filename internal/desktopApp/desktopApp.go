package desktopApp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	catppuccin "github.com/catppuccin/go"
	"github.com/hra42/Go-DNS/internal/dns"
	"github.com/hra42/Go-DNS/internal/sslchecker"
	"image/color"
	"log"
)

func RunDesktop() {
	// create a new app and window
	a := app.New()
	w := a.NewWindow("Go-DNS")

	// set the icon from cloudflare r2 bucket
	resourceIconPng, err := fyne.LoadResourceFromURLString("https://gpt-files.postrausch.tech/go-dns.png")
	if err != nil {
		log.Fatal(err)
	}
	w.SetIcon(resourceIconPng)

	// Resize and center the window
	w.Resize(fyne.NewSize(1000, 1000))
	w.CenterOnScreen()

	// main menu
	dnsRecords := fyne.NewMenu(
		"DNS Records",
		fyne.NewMenuItem("CNAME", func() {
			getCNameRecords(w)
		}),
		fyne.NewMenuItem("MX", func() {
			getMX(w)
		}),
		fyne.NewMenuItem("SSL Checker", func() {
			checkSSL(w)
		}),
		fyne.NewMenuItem("TXT", func() {
			getTXT(w)
		}),
	)

	// theme menu
	themeSwitcher := fyne.NewMenu(
		"Theme",
		// TODO: replace this with a function that is not deprecated & define catppuccin theme
		fyne.NewMenuItem("Light", func() {
			a.Settings().SetTheme(theme.LightTheme())
		}),
		fyne.NewMenuItem("Dark", func() {
			a.Settings().SetTheme(theme.DarkTheme())
		}),
	)

	// make menu
	mainMenu := fyne.NewMainMenu(dnsRecords, themeSwitcher)
	w.SetMainMenu(mainMenu)

	// Set the welcome title when the app starts
	welcomeTitleText := getTitle(w, "")
	welcomeTitle := container.NewCenter(welcomeTitleText)

	// Set the welcome text when the app starts
	// make text light blue, bold and centered with text size of 20.
	welcomeText := canvas.NewText("This is a tool to check DNS records for Microsoft 365.",
		getCatppucinColor(catppuccin.Mocha.Blue()))
	welcomeText.Alignment = fyne.TextAlignCenter
	welcomeText.TextSize = 20.0
	welcome := container.NewCenter(welcomeText)

	// set the content on canvas
	w.SetContent(container.NewVBox(
		welcomeTitle,
		welcome,
	))

	// start the main app
	w.ShowAndRun()
}

func getCNameRecords(w fyne.Window) {
	// Set up the window for CNAME Records Lookup
	title := getTitle(w, "CNAME Records")
	inputDomain := widget.NewEntry()
	inputDomain.SetPlaceHolder("Enter the domain you want to check")
	warning := canvas.NewText("Attention! This can take a long time", color.RGBA{R: 255, G: 0, B: 0, A: 255})
	warning.TextStyle.Bold = true
	placeholder := widget.NewLabel("This is a placeholder")
	placeholder.Hide()

	box := container.NewVBox(
		title,
		warning,
		inputDomain,
		placeholder,
		widget.NewButton("Check", func() {
			report := ""
			report += dns.GetCnameReport(inputDomain.Text)
			warning.Hide()
			placeholder.SetText(report)
			placeholder.Show()
		}),
	)

	scroll := container.NewScroll(box)
	w.SetContent(scroll)
}

func getMX(w fyne.Window) {
	// set up the window for MX Record Lookup
	title := getTitle(w, "MX Records")
	inputDomain := widget.NewEntry()
	inputDomain.SetPlaceHolder("Enter the domain you want to check")
	placeholder := widget.NewLabel("This is a placeholder")
	placeholder.Hide()

	box := container.NewVBox(
		title,
		inputDomain,
		placeholder,
		widget.NewButton("Check", func() {
			report := ""
			report += dns.GetMXReport(inputDomain.Text)
			placeholder.SetText(report)
			placeholder.Show()
		}),
	)

	scroll := container.NewScroll(box)
	w.SetContent(scroll)
}

func checkSSL(w fyne.Window) {
	// Set up the window for SSL Checker
	title := getTitle(w, "SSL Checker")
	inputDomain := widget.NewEntry()
	inputDomain.SetPlaceHolder("Enter the domain you want to check - Format: domain.tld:Port")
	placeholder := widget.NewLabel("This is a placeholder")
	placeholder.Hide()

	// Set the content on canvas
	box := container.NewVBox(
		title,
		inputDomain,
		placeholder,
		widget.NewButton("Check", func() {
			report := ""
			report += sslchecker.CheckSSL(inputDomain.Text)
			placeholder.SetText(report)
			placeholder.Show()
		}),
	)
	scroll := container.NewScroll(box)
	w.SetContent(scroll)
}

func getTXT(w fyne.Window) {
	// set up the window for TXT Record Lookup
	title := getTitle(w, "TXT Records")
	inputDomain := widget.NewEntry()
	inputDomain.SetPlaceHolder("Enter the domain you want to check")
	placeholder := widget.NewLabel("This is a placeholder")
	placeholder.Hide()

	box := container.NewVBox(
		title,
		inputDomain,
		placeholder,
		widget.NewButton("Check", func() {
			report := ""
			report += dns.GetTXTReport(inputDomain.Text)
			placeholder.SetText(report)
			placeholder.Show()
		}),
	)

	scroll := container.NewScroll(box)
	w.SetContent(scroll)
}

func getTitle(w fyne.Window, titleText string) (title *canvas.Text) {
	blueTitle := getCatppucinColor(catppuccin.Mocha.Blue())
	// set the title of a window
	if titleText == "" {
		w.SetTitle("Go-DNS")
		title = canvas.NewText("Welcome to Go-DNS", blueTitle)
		title.TextStyle.Bold = true
		title.Alignment = fyne.TextAlignCenter
		title.TextSize = 60.0
	} else {
		w.SetTitle("Go-DNS - " + titleText)
		// define the title text
		title = canvas.NewText(titleText, blueTitle)
		title.TextStyle.Bold = true
		title.Alignment = fyne.TextAlignCenter
		title.TextSize = 60.0
	}
	return
}

func getCatppucinColor(c catppuccin.Color) (col color.Color) {
	// change opacity of the color to 255
	col = color.RGBA{R: uint8(c.RGB[0]), G: uint8(c.RGB[1]), B: uint8(c.RGB[2]), A: 255}
	return
}
