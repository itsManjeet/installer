package app

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"os/user"

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

func (app *App) CreateList(list *gtk.ListBox) (*gtk.ScrolledWindow, error) {
	scrolledWidth, err := gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		return nil, err
	}
	scrolledWidth.SetMarginStart(185)
	scrolledWidth.SetMarginEnd(185)

	viewPort, err := gtk.ViewportNew(nil, nil)
	if err != nil {
		return nil, err
	}
	scrolledWidth.Add(viewPort)
	viewPort.Add(list)

	return scrolledWidth, nil
}

func (app *App) UpdateListText(list *gtk.ListBox, data ...string) error {

	widgets := []*gtk.Widget{}
	for _, a := range data {
		wid, err := gtk.LabelNew(a)
		if err != nil {
			return err
		}
		wid.SetHAlign(gtk.ALIGN_CENTER)
		wid.SetHExpand(true)
		widgets = append(widgets, &wid.Widget)
		wid.ShowAll()
	}

	return app.UpdateList(list, widgets...)
}

func (app *App) UpdateList(list *gtk.ListBox, data ...*gtk.Widget) error {

	list.GetChildren().Foreach(func(f interface{}) {
		list.Remove(f.(gtk.IWidget))
	})

	list.SetHeaderFunc(func(row, before *gtk.ListBoxRow) {
		if before != nil {
			sep, _ := gtk.SeparatorNew(gtk.ORIENTATION_HORIZONTAL)
			sep.Show()
			row.SetHeader(sep)
		}
	})

	for pos, i := range data {
		box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
		if err != nil {
			return err
		}
		row, err := gtk.ListBoxRowNew()
		if err != nil {
			return err
		}
		box.SetBorderWidth(16)
		box.Add(i)
		row.Add(box)

		list.Insert(row, pos)
	}

	return nil
}

func (app *App) GetListText(row *gtk.ListBoxRow) (string, error) {
	box, err := row.GetChild()
	if err != nil {
		return "", err
	}
	boxWidget, ok := box.(*gtk.Box)
	if !ok {
		return "", errors.New("not a box widget")
	}

	labelWidget, ok := (boxWidget.GetChildren().First().Data().(*gtk.Widget))
	if !ok {
		return "", errors.New("not a label widget")
	}

	label := &gtk.Label{Widget: *labelWidget}

	return label.GetText()
}

func (app *App) ExecAsUser(uid string, cmd ...string) error {
	u, err := user.LookupId(uid)
	if err != nil {
		return err
	}
	command := []string{
		"pkexec", "--username", u.Username,
	}

	command = append(command, cmd...)
	return exec.Command(command[0], command[1:]...).Run()
}
