package installer

import (
	"os/exec"

	"github.com/gotk3/gotk3/gtk"
	"github.com/rlxos/installer/app"
)

type Installer struct {
	*app.App
	WelcomePage *app.Page

	VerifyPage     *app.Page
	verifyProgress *gtk.ProgressBar

	ConfirmPage   *app.Page
	confirmBuffer *gtk.TextBuffer

	InstallPage     *app.Page
	installProgress *gtk.ProgressBar

	FinishedPage *app.Page
}

func Setup(win *gtk.Assistant) error {
	ins := &Installer{}
	var err error
	ins.App, err = app.Setup(win)
	if err != nil {
		return err
	}

	ins.WelcomePage, err = ins.NewTitledPage("Welcome!", "Click on 'Continue' to verify system compatibality with rlxos", "rlxos", nil)
	if err != nil {
		return err
	}

	continueButton, err := gtk.ButtonNewWithLabel("Continue")
	if err != nil {
		return err
	}
	continueButton.SetMarginTop(28)
	continueButton.SetHAlign(gtk.ALIGN_CENTER)
	ins.WelcomePage.Box.PackStart(continueButton, false, false, 0)

	continueButton.Connect("clicked", func() {
		win.SetCurrentPage(1)
		go ins.StartVerification()
	})

	win.AppendPage(ins.WelcomePage)
	win.SetPageType(ins.WelcomePage, gtk.ASSISTANT_PAGE_CUSTOM)
	win.SetPageComplete(ins.WelcomePage, true)

	//
	// Verify Page
	//
	ins.VerifyPage, err = ins.NewPage("System Verification", "Verifying System Compatibility with rlxos", "btsync-gui-0", nil)
	if err != nil {
		return err
	}

	win.AppendPage(ins.VerifyPage)
	win.SetPageType(ins.VerifyPage, gtk.ASSISTANT_PAGE_CONTENT)

	ins.verifyProgress, err = gtk.ProgressBarNew()
	if err != nil {
		return err
	}
	ins.verifyProgress.SetShowText(true)
	ins.verifyProgress.SetMarginStart(28)
	ins.verifyProgress.SetMarginEnd(28)
	ins.VerifyPage.Box.PackStart(ins.verifyProgress, false, false, 128)

	//
	// Confirm Page
	//
	ins.ConfirmPage, err = ins.NewPage("Confirm", "Please confirm the configurations", "configure", nil)
	if err != nil {
		return err
	}
	win.AppendPage(ins.ConfirmPage)
	win.SetPageType(ins.ConfirmPage, gtk.ASSISTANT_PAGE_CONFIRM)
	textView, err := gtk.TextViewNew()
	if err != nil {
		return err
	}
	textView.SetCursorVisible(false)
	textView.SetEditable(false)
	textView.SetTopMargin(15)
	textView.SetLeftMargin(25)
	ins.ConfirmPage.Box.PackStart(textView, true, true, 0)
	ins.confirmBuffer, err = textView.GetBuffer()
	if err != nil {
		return err
	}

	ins.Window.Connect("apply", func() {
		win.SetCurrentPage(2)
		go ins.StartInstallation()
	})

	//
	// Install Page
	//
	ins.InstallPage, err = ins.NewPage("Installation", "Please wait while finishing Installation", "download", nil)
	if err != nil {
		return err
	}
	win.AppendPage(ins.InstallPage)
	win.SetPageType(ins.InstallPage, gtk.ASSISTANT_PAGE_PROGRESS)

	ins.installProgress, err = gtk.ProgressBarNew()
	if err != nil {
		return err
	}
	ins.installProgress.SetShowText(true)
	ins.installProgress.SetMarginStart(28)
	ins.installProgress.SetMarginEnd(28)
	ins.InstallPage.Box.PackStart(ins.installProgress, false, false, 128)

	ins.FinishedPage, err = ins.NewTitledPage("Success", "rlxos is installed successfully, click 'reboot' to reboot into your freshed system", "emblem-checked", nil)
	if err != nil {
		return err
	}
	rebootButton, err := gtk.ButtonNewWithLabel("reboot")
	if err != nil {
		return err
	}
	rebootButton.SetHAlign(gtk.ALIGN_CENTER)
	rebootButton.Connect("clicked", func() {
		if !ins.IsDebug(APPID) {
			exec.Command("reboot").Run()
		} else {
			app, _ := win.GetApplication()
			app.Quit()
		}

	})
	ins.FinishedPage.Box.PackStart(rebootButton, false, false, 28)

	win.AppendPage(ins.FinishedPage)
	win.SetPageType(ins.FinishedPage, gtk.ASSISTANT_PAGE_CUSTOM)

	return err
}
