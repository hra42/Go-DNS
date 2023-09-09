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

type catppuccinTheme struct{}

func (c catppuccinTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		if variant == theme.VariantLight {
			return getCatppucinColor(catppuccin.Latte.Base())
		} else {
			return getCatppucinColor(catppuccin.Mocha.Base())
		}
	case theme.ColorNamePlaceHolder:
		if variant == theme.VariantLight {
			return getCatppucinColor(catppuccin.Latte.Text())
		} else {
			return getCatppucinColor(catppuccin.Mocha.Blue())
		}
	case theme.ColorNameInputBackground:
		if variant == theme.VariantLight {
			return getCatppucinColor(catppuccin.Latte.Base())
		} else {
			return getCatppucinColor(catppuccin.Mocha.Base())
		}
	case theme.ColorNameSeparator:
		if variant == theme.VariantLight {
			return getCatppucinColor(catppuccin.Latte.Red())
		} else {
			return getCatppucinColor(catppuccin.Mocha.Red())
		}
	case theme.ColorNamePrimary:
		if variant == theme.VariantLight {
			return getCatppucinColor(catppuccin.Latte.Subtext0())
		} else {
			return getCatppucinColor(catppuccin.Mocha.Subtext0())
		}
	case theme.ColorNameSuccess:
		if variant == theme.VariantLight {
			return getCatppucinColor(catppuccin.Latte.Blue())
		} else {
			return getCatppucinColor(catppuccin.Mocha.Blue())
		}
	case theme.ColorNameButton:
		if variant == theme.VariantLight {
			return getCatppucinColor(catppuccin.Latte.Surface0())
		} else {
			return getCatppucinColor(catppuccin.Mocha.Surface0())
		}
	case theme.ColorNameMenuBackground:
		if variant == theme.VariantLight {
			return getCatppucinColor(catppuccin.Latte.Surface1())
		} else {
			return getCatppucinColor(catppuccin.Mocha.Surface1())
		}
	case theme.ColorNameHover:
		if variant == theme.VariantLight {
			return getCatppucinColor(catppuccin.Latte.Surface2())
		} else {
			return getCatppucinColor(catppuccin.Mocha.Surface2())
		}
	case theme.ColorNameError:
		if variant == theme.VariantLight {
			return getCatppucinColor(catppuccin.Latte.Red())
		} else {
			return getCatppucinColor(catppuccin.Mocha.Red())
		}
	case theme.ColorNameSelection:
		if variant == theme.VariantLight {
			return getCatppucinColor(catppuccin.Latte.Lavender())
		} else {
			return getCatppucinColor(catppuccin.Mocha.Lavender())
		}
	case theme.ColorNameDisabled:
		if variant == theme.VariantLight {
			return getCatppucinColor(catppuccin.Latte.Crust())
		} else {
			return getCatppucinColor(catppuccin.Mocha.Crust())
		}
	case theme.ColorNameFocus:
		if variant == theme.VariantLight {
			return getCatppucinColor(catppuccin.Latte.Lavender())
		} else {
			return getCatppucinColor(catppuccin.Mocha.Lavender())
		}
	default:
		if variant == theme.VariantLight {
			return getCatppucinColor(catppuccin.Latte.Text())
		} else {
			return getCatppucinColor(catppuccin.Mocha.Text())
		}
	}
}

func (c catppuccinTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (c catppuccinTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (c catppuccinTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

type forceVariantTheme struct {
	fyne.Theme

	variant fyne.ThemeVariant
}

func (f *forceVariantTheme) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return f.Theme.Color(name, f.variant)
}

func RunDesktop() {
	// create a new app and window
	a := app.New()
	w := a.NewWindow("Go-DNS")

	// set the default theme to latte or mocha
	a.Settings().SetTheme(&catppuccinTheme{})

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
			getCNameRecords(a, w)
		}),
		fyne.NewMenuItem("MX", func() {
			getMX(a, w)
		}),
		fyne.NewMenuItem("SSL Checker", func() {
			checkSSL(a, w)
		}),
		fyne.NewMenuItem("TXT", func() {
			getTXT(a, w)
		}),
	)

	// theme menu – override the default theme
	themeSwitcher := fyne.NewMenu(
		"Theme",
		fyne.NewMenuItem("Light", func() {
			a.Settings().SetTheme(&forceVariantTheme{Theme: catppuccinTheme{}, variant: theme.VariantLight})
		}),
		fyne.NewMenuItem("Dark", func() {
			a.Settings().SetTheme(&forceVariantTheme{Theme: catppuccinTheme{}, variant: theme.VariantDark})
		}),
	)

	// make menu
	mainMenu := fyne.NewMainMenu(dnsRecords, themeSwitcher)
	w.SetMainMenu(mainMenu)

	// Set the welcome title when the app starts
	welcomeTitleText := getTitle(a, w, "")
	welcomeTitle := container.NewCenter(welcomeTitleText)

	// Set the welcome text when the app starts
	// make text light blue, bold and centered with text size of 20.
	var welcomeTextColor color.Color
	if a.Settings().ThemeVariant() == theme.VariantLight {
		welcomeTextColor = getCatppucinColor(catppuccin.Latte.Blue())
	} else {
		welcomeTextColor = getCatppucinColor(catppuccin.Mocha.Blue())
	}
	welcomeText := canvas.NewText("This is a tool to check DNS records for Microsoft 365.",
		welcomeTextColor)
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

func getCNameRecords(a fyne.App, w fyne.Window) {
	// Set up the window for CNAME Records Lookup
	title := getTitle(a, w, "CNAME Records")
	inputDomain := widget.NewEntry()
	inputDomain.SetPlaceHolder("Enter the domain you want to check")
	warning := canvas.NewText("Attention! This can take a long time", getCatppucinColor(catppuccin.Mocha.Red()))
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

func getMX(a fyne.App, w fyne.Window) {
	// set up the window for MX Record Lookup
	title := getTitle(a, w, "MX Records")
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

func checkSSL(a fyne.App, w fyne.Window) {
	// Set up the window for SSL Checker
	title := getTitle(a, w, "SSL Checker")
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

func getTXT(a fyne.App, w fyne.Window) {
	// set up the window for TXT Record Lookup
	title := getTitle(a, w, "TXT Records")
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

func getTitle(a fyne.App, w fyne.Window, titleText string) (title *canvas.Text) {
	// define title color
	var blueTitle color.Color
	if a.Settings().ThemeVariant() == theme.VariantLight {
		blueTitle = getCatppucinColor(catppuccin.Latte.Blue())
	} else {
		blueTitle = getCatppucinColor(catppuccin.Mocha.Blue())
	}
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
