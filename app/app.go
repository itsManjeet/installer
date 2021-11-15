package app

import (
	"log"
	"os"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type App struct {
	Window *gtk.Assistant
}

func Setup(win *gtk.Assistant) (app *App, err error) {
	app = &App{Window: win}
	return
}

func (app App) GetIcon(logo string, size int) (pixbuf *gdk.Pixbuf, err error) {

	theme, err := gtk.IconThemeGetDefault()
	if err != nil {
		return
	}

	pixbuf, err = theme.LoadIcon(logo, size, gtk.ICON_LOOKUP_FORCE_SIZE)
	if err != nil {
		return nil, nil
	}

	return
}

func (app *App) NewPage(title, subtitle, icon string, data interface{}) (page *Page, err error) {
	page = &Page{
		GlobalData: data,
	}

	page.Window = app.Window

	page.Box, err = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return
	}

	pixbuf, err := app.GetIcon(icon, 64)
	if err != nil {
		return
	}
	page.Icon, err = gtk.ImageNewFromPixbuf(pixbuf)
	if err != nil {
		return
	}

	page.Box.PackStart(page.Icon, false, false, 17)

	page.Title, err = gtk.LabelNew("")
	if err != nil {
		return
	}
	page.Title.SetMarkup("<span size=\"xx-large\" weight=\"ultrabold\">" + title + "</span>")
	page.Box.PackStart(page.Title, false, false, 0)

	page.SubTitle, err = gtk.LabelNew("")
	if err != nil {
		return
	}
	page.SubTitle.SetMarkup("<span weight=\"bold\">" + subtitle + "</span>")
	page.SubTitle.SetMarginBottom(27)
	page.Box.PackStart(page.SubTitle, false, false, 0)

	return
}

func (app *App) NewTitledPage(title, subtitle, icon string, data interface{}) (page *Page, err error) {
	page, err = app.NewPage(title, subtitle, icon, data)
	if err != nil {
		return
	}

	pixbuf, err := app.GetIcon(icon, 128)
	if err != nil {
		return
	}

	page.Icon.SetFromPixbuf(pixbuf)
	page.Box.SetVAlign(gtk.ALIGN_CENTER)
	page.Box.SetMarginBottom(150)

	return
}

func (app *App) IsDebug(appID string) bool {
	if env := os.Getenv(appID + "_DEBUG"); len(env) == 0 {
		log.Println("Debug is disabled")
		return false
	}
	log.Println("Debug is enabled")
	return true
}
