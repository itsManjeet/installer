package setup

import (
	"log"
	"os/exec"

	"github.com/gotk3/gotk3/gtk"
	"github.com/rlxos/installer/app"
)

type Setup struct {
	*app.App
}

func Init(ui *gtk.Builder) *Setup {
	var setup Setup

	setup.App = app.Create(ui)

	setup.ID = ID
	setup.Title = TITLE
	setup.WelcomeMesg = WELCOME_MESG
	setup.ProcessTitle = PROCESS_TITLE
	setup.ProcessMesg = PROCESS_DESC
	setup.SuccessMesg = SUCCESS_MESG
	setup.SuccessBtn = SUCCESS_BTN

	setup.Initialize()

	setup.AddStage("Setting up locales", setup.StageLocale)
	setup.AddStage("Setting up TimeZone", setup.StageTimezone)
	setup.AddStage("Setting up Account", setup.StageAccount)
	setup.AddStage("Post cleanup processes", setup.StagePost)

	setup.SuccessHandler = func() {
		log.Println("logging out from system")
		if !setup.IsDebug() {
			exec.Command("xfce4-session-logout", "--logout").Run()
		}
	}

	return &setup
}
