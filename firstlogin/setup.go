package firstlogin

import (
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/gotk3/gotk3/gtk"
	"github.com/rlxos/installer/app"
)

type FirstLogin struct {
	*app.App

	WelcomePage *app.Page
	langList    *gtk.ListBox

	TimeZonePage *app.Page
	timeZoneList *gtk.ListBox

	UserAccountPage *app.Page
	FinishedPage    *app.Page
}

func Setup(win *gtk.Assistant) error {

	f := &FirstLogin{}
	var err error

	f.App, err = app.Setup(win)
	if err != nil {
		return err
	}

	f.WelcomePage, err = f.NewPage("Welcome!", "Thanks for choosing rlxos", "rlxos", nil)
	if err != nil {
		return err
	}
	f.langList, err = gtk.ListBoxNew()
	if err != nil {
		return err
	}

	f.UpdateListText(f.langList, "en_IN.UTF-8", "en_US.UTF-8")
	f.langList.Connect("row-activated", func(list *gtk.ListBox, row *gtk.ListBoxRow) {
		SelectedLocale, _ = f.GetListText(row)
		log.Println("Selected Locale: ", SelectedLocale)
		go f.generateLocale(SelectedLocale)
		win.SetPageComplete(f.WelcomePage, true)
	})
	wid, err := f.CreateList(f.langList)
	if err != nil {
		return err
	}
	f.WelcomePage.Box.PackStart(wid, true, true, 0)

	win.AppendPage(f.WelcomePage)
	win.SetPageType(f.WelcomePage, gtk.ASSISTANT_PAGE_INTRO)

	//
	// Timezone
	//
	f.TimeZonePage, err = f.NewPage("TimeZone", "Select your timezone", "time", nil)
	if err != nil {
		return err
	}
	f.timeZoneList, err = gtk.ListBoxNew()
	if err != nil {
		return err
	}
	f.UpdateListText(f.timeZoneList, f.getTimeZoneList()...)
	f.timeZoneList.Connect("row-activated", func(list *gtk.ListBox, row *gtk.ListBoxRow) {

	})
	timeZoneListWidget, err := f.CreateList(f.timeZoneList)
	if err != nil {
		return err
	}
	f.TimeZonePage.Box.PackStart(timeZoneListWidget, true, true, 0)
	f.timeZoneList.Connect("row-activated", func(list *gtk.ListBox, row *gtk.ListBoxRow) {
		SelectedTimeZone, _ = f.GetListText(row)
		log.Println("Selected Timezone: ", SelectedTimeZone)
		go os.Link(path.Join(TIMEZONE_DIR, SelectedTimeZone), "/etc/localtime")
		win.SetPageComplete(f.TimeZonePage, true)
	})

	win.AppendPage(f.TimeZonePage)
	win.SetPageType(f.TimeZonePage, gtk.ASSISTANT_PAGE_CONTENT)

	//
	// User Account
	//
	f.UserAccountPage, err = f.NewTitledPage("UserAccount", "Create Primary user", "im-user", nil)
	if err != nil {
		return err
	}
	win.AppendPage(f.UserAccountPage)
	win.SetPageType(f.UserAccountPage, gtk.ASSISTANT_PAGE_CONTENT)
	useridBox, _ := gtk.EntryNew()
	useridBox.SetPlaceholderText("Username")
	useridBox.SetMarginStart(178)
	useridBox.SetMarginEnd(178)
	useridBox.SetMarginBottom(12)
	f.UserAccountPage.Box.PackStart(useridBox, false, false, 0)

	passwdBox, _ := gtk.EntryNew()
	passwdBox.SetPlaceholderText("Password")
	passwdBox.SetMarginBottom(16)
	passwdBox.SetMarginStart(178)
	passwdBox.SetMarginEnd(178)
	passwdBox.SetVisibility(false)
	passwdBox.SetIconFromIconName(gtk.ENTRY_ICON_SECONDARY, "view-hidden")
	passwdBox.Connect("icon-press", func(entry *gtk.Entry, iconType gtk.EntryIconPosition) {
		passwdBox.SetVisibility(!passwdBox.GetVisibility())
		if passwdBox.GetVisibility() {
			passwdBox.SetIconFromIconName(gtk.ENTRY_ICON_SECONDARY, "image-red-eye")
		} else {
			passwdBox.SetIconFromIconName(gtk.ENTRY_ICON_SECONDARY, "view-hidden")
		}
	})
	f.UserAccountPage.Box.PackStart(passwdBox, false, false, 0)

	createUserBtn, _ := gtk.ButtonNewWithLabel("Create")
	createUserBtn.SetHAlign(gtk.ALIGN_CENTER)
	f.UserAccountPage.PackStart(createUserBtn, false, false, 0)
	createUserBtn.Connect("clicked", func() {

		userid, _ := useridBox.GetText()
		passwd, _ := passwdBox.GetText()

		go f.createUser(userid, passwd)
	})

	//
	// Finished Page
	//
	f.FinishedPage, err = f.NewTitledPage("Success", "First login tasks done successfully, Enjoy", "emblem-checked", nil)
	if err != nil {
		return err
	}
	finishedBtn, err := gtk.ButtonNewWithLabel("Done")
	if err != nil {
		return err
	}
	f.FinishedPage.Box.PackStart(finishedBtn, false, false, 0)
	finishedBtn.SetHAlign(gtk.ALIGN_CENTER)
	finishedBtn.Connect("clicked", func() {
		if !f.IsDebug(APPID) {
			exec.Command("xfce4-session-logout", "--logout").Run()
		} else {
			app, _ := win.GetApplication()
			app.Quit()
		}
	})
	win.AppendPage(f.FinishedPage)
	win.SetPageType(f.FinishedPage, gtk.ASSISTANT_PAGE_CUSTOM)

	return nil

}
