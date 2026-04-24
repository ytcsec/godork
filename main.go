package main

import (
	"fmt"
	"image/color"
	"log"
	"net/url"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type DorkSonuc struct {
	Kategori string
	Sorgu    string
	Link     string
}

// dorkları güncellemek isteyenler için dork formatı:"kategori|dork"
// kategori eklersen de KategoriIsımlerı olan mape eklemeyi untuma
var dorkListesi = []string{
	"HassasDosyalar|filetype:pdf",
	"HassasDosyalar|filetype:xls",
	"HassasDosyalar|filetype:xlsx",
	"HassasDosyalar|filetype:doc",
	"HassasDosyalar|filetype:docx",
	"HassasDosyalar|filetype:txt",
	"HassasDosyalar|filetype:sql",
	"HassasDosyalar|filetype:env",
	"HassasDosyalar|filetype:log",
	"HassasDosyalar|filetype:bak",
	"HassasDosyalar|filetype:conf",
	"HassasDosyalar|filetype:ini",
	"HassasDosyalar|filetype:csv",

	"DizinListeleme|intitle:\"index of\"",
	"DizinListeleme|intitle:\"index of\" \"parent directory\"",
	"DizinListeleme|intitle:\"index of /admin\"",
	"DizinListeleme|intitle:\"index of /backup\"",

	"LoginSayfalari|inurl:login",
	"LoginSayfalari|inurl:admin",
	"LoginSayfalari|inurl:admin/login",
	"LoginSayfalari|inurl:signin",
	"LoginSayfalari|intitle:\"admin login\"",

	"HataMesajlari|intext:\"sql syntax near\"",
	"HataMesajlari|intext:\"mysql_fetch\"",
	"HataMesajlari|\"warning: mysql\"",
	"HataMesajlari|intext:\"fatal error\"",
	"HataMesajlari|intext:\"stack trace\"",

	"Paneller|inurl:phpmyadmin",
	"Paneller|inurl:cpanel",
	"Paneller|inurl:wp-admin",
	"Paneller|inurl:webmail",
	"Paneller|inurl:plesk",

	"ConfigDosyalari|inurl:config",
	"ConfigDosyalari|inurl:wp-config",
	"ConfigDosyalari|inurl:.env",
	"ConfigDosyalari|inurl:.git",
	"ConfigDosyalari|filetype:yml inurl:config",

	"YedekDosyalar|inurl:backup",
	"YedekDosyalar|ext:bak",
	"YedekDosyalar|ext:old",
	"YedekDosyalar|inurl:backup filetype:sql",

	"KullaniciBilgileri|intext:\"username\" filetype:log",
	"KullaniciBilgileri|intext:\"password\" filetype:txt",
	"KullaniciBilgileri|intext:\"email\" filetype:csv",
	"KullaniciBilgileri|intext:\"api_key\"",

	"KameraIoT|inurl:view/index.shtml",
	"KameraIoT|intitle:\"webcam\"",
	"KameraIoT|intitle:\"live view\" inurl:axis",
	"KameraIoT|inurl:\"ViewerFrame?Mode=\"",

	"KisiselInfo|intext:\"linkedin.com/in/\"",
	"KisiselInfo|intext:\"contact\" filetype:pdf",
	"KisiselInfo|intext:\"cv\" filetype:pdf",
}

var kategoriIsimleri = map[string]string{
	"HassasDosyalar":     "Gizli Dosyalar",
	"DizinListeleme":     "Dizinler",
	"LoginSayfalari":     "Loginler",
	"HataMesajlari":      "Errorlar",
	"Paneller":           "Paneller",
	"ConfigDosyalari":    "Configler",
	"YedekDosyalar":      "Yedekler",
	"KullaniciBilgileri": "Kullanıcılar",
	"KameraIoT":          "IoT",
	"KisiselInfo":        "Kişisel",
}

func searchURLHazirla(target, dork string) string {
	target = strings.TrimSpace(target)

	q := dork
	if target != "" {
		q = "site:" + target + " " + dork
	}

	return "https://www.google.com/search?q=" + url.QueryEscape(q)
}

func parseDorkSatir(satir string) (string, string, bool) {
	parcalar := strings.Split(satir, "|")

	if len(parcalar) < 2 {
		return "", "", false
	}

	kategori := parcalar[0]
	sorgu := parcalar[1]

	if kategori == "" || sorgu == "" {
		return "", "", false
	}

	return kategori, sorgu, true
}

func dorklariHazirla(hedef string) (map[string][]DorkSonuc, int) {
	sonuc := make(map[string][]DorkSonuc)
	hedef = strings.TrimSpace(hedef)

	var total int
	var bozuk int

	for _, satir := range dorkListesi {
		kategoriKod, dorkSorgu, ok := parseDorkSatir(satir)
		if !ok {
			bozuk++
			log.Printf("sorunlu satırlar geçildi: %q", satir)
			continue
		}

		kategoriAdi := kategoriIsimleri[kategoriKod]
		if kategoriAdi == "" {
			kategoriAdi = kategoriKod
		}

		var displayQuery string
		if hedef == "" {
			displayQuery = dorkSorgu
		} else {
			displayQuery = fmt.Sprintf("site:%s %s", hedef, dorkSorgu)
		}

		sonuc[kategoriAdi] = append(sonuc[kategoriAdi], DorkSonuc{
			Kategori: kategoriAdi,
			Sorgu:    displayQuery,
			Link:     searchURLHazirla(hedef, dorkSorgu),
		})
		total++
	}

	if bozuk > 0 {
		log.Printf("toplam %d dork hazrılandı, %d bozuk dork verdi", total, bozuk)
	} else {
		log.Printf("toplam %d dork üretildi", total)
	}

	return sonuc, total
}

type siyahBeyaz struct{}

func (siyahBeyaz) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch n {
	case theme.ColorNameBackground:
		return color.Black
	case theme.ColorNameForeground:
		return color.White
	case theme.ColorNameInputBackground:
		return color.NRGBA{R: 15, G: 15, B: 15, A: 255}
	case theme.ColorNameButton:
		return color.NRGBA{R: 25, G: 25, B: 25, A: 255}
	case theme.ColorNameDisabled:
		return color.NRGBA{R: 120, G: 120, B: 120, A: 255}
	case theme.ColorNameHover:
		return color.NRGBA{R: 45, G: 45, B: 45, A: 255}
	case theme.ColorNamePlaceHolder:
		return color.NRGBA{R: 130, G: 130, B: 130, A: 255}
	case theme.ColorNamePrimary:
		return color.White
	case theme.ColorNameSeparator:
		return color.NRGBA{R: 60, G: 60, B: 60, A: 255}
	case theme.ColorNameScrollBar:
		return color.NRGBA{R: 80, G: 80, B: 80, A: 255}
	case theme.ColorNameInputBorder:
		return color.NRGBA{R: 70, G: 70, B: 70, A: 255}
	}
	return theme.DefaultTheme().Color(n, v)
}

func (siyahBeyaz) Font(s fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(s)
}

func (siyahBeyaz) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (siyahBeyaz) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}

func kategoriSirasi() []string {
	var sira []string
	gorulen := map[string]bool{}
	for _, satir := range dorkListesi {
		kod, _, ok := parseDorkSatir(satir)
		if !ok {
			continue
		}
		ad := kategoriIsimleri[kod]
		if ad == "" {
			ad = kod
		}
		if !gorulen[ad] {
			gorulen[ad] = true
			sira = append(sira, ad)
		}
	}
	return sira
}

func sonuclariYazdir(liste *fyne.Container, hedef string) {
	liste.RemoveAll()

	hedef = strings.TrimSpace(hedef)
	if hedef == "" {
		uyari := canvas.NewText("Geçerli bir hedef girmediniz.", color.White)
		uyari.TextSize = 14
		liste.Add(uyari)
		liste.Refresh()
		return
	}

	sonuclar, toplam := dorklariHazirla(hedef)

	ozet := canvas.NewText(fmt.Sprintf("Hedef: %s    Toplam: %d dork", hedef, toplam), color.White)
	ozet.TextSize = 15
	ozet.TextStyle = fyne.TextStyle{Bold: true}
	liste.Add(ozet)
	liste.Add(widget.NewSeparator())

	for _, kat := range kategoriSirasi() {
		dorklar, ok := sonuclar[kat]
		if !ok || len(dorklar) == 0 {
			continue
		}

		baslik := canvas.NewText(kat, color.White)
		baslik.TextSize = 15
		baslik.TextStyle = fyne.TextStyle{Bold: true}
		liste.Add(baslik)

		for _, d := range dorklar {
			sorguLbl := widget.NewLabel(d.Sorgu)
			sorguLbl.Wrapping = fyne.TextWrapBreak
			sorguLbl.TextStyle = fyne.TextStyle{Monospace: true}

			parsed, err := url.Parse(d.Link)
			var linkObj fyne.CanvasObject
			if err == nil {
				linkObj = widget.NewHyperlink(d.Link, parsed)
			} else {
				linkObj = widget.NewLabel(d.Link)
			}

			liste.Add(container.NewVBox(sorguLbl, linkObj))
		}

		liste.Add(widget.NewSeparator())
	}

	liste.Refresh()
}

func main() {
	log.SetFlags(log.LstdFlags)

	a := app.New()
	a.Settings().SetTheme(siyahBeyaz{})

	w := a.NewWindow("GoDork")
	w.Resize(fyne.NewSize(900, 720))

	baslik := canvas.NewText("Google Dork Searcher", color.White)
	baslik.TextSize = 22
	baslik.TextStyle = fyne.TextStyle{Bold: true}

	aciklama := widget.NewLabel("Hedef vermesi sizdne dork üretmesi bizden")
	aciklama.Wrapping = fyne.TextWrapWord

	hedefGiris := widget.NewEntry()
	hedefGiris.SetPlaceHolder("site.com")

	sonucKutu := container.NewVBox()
	scroll := container.NewVScroll(sonucKutu)
	scroll.SetMinSize(fyne.NewSize(880, 420))

	ara := func() {
		sonuclariYazdir(sonucKutu, hedefGiris.Text)
	}
	hedefGiris.OnSubmitted = func(string) { ara() }
	araBtn := widget.NewButton("Go", ara)

	form := container.NewBorder(nil, nil, nil, araBtn, hedefGiris)

	ust := container.NewVBox(
		baslik,
		aciklama,
		form,
		widget.NewSeparator(),
	)

	w.SetContent(container.NewBorder(ust, nil, nil, nil, scroll))
	w.ShowAndRun()
}
