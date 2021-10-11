package installer

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/rlxos/installer/app"
	"github.com/rlxos/installer/disk"
)

type Installer struct {
	*app.App

	devices disk.BlockDevices
	tempdir string
}

func Init(ui *gtk.Builder) *Installer {
	var in Installer

	in.App = app.Create(ui)

	in.ID = ID
	in.Title = TITLE
	in.WelcomeMesg = WELCOME_MESG
	in.ProcessTitle = PROCESS_TITLE
	in.ProcessMesg = PROCESS_DESC
	in.SuccessMesg = SUCCESS_MESG
	in.SuccessBtn = SUCCESS_BTN

	in.Initialize()

	in.Stages = map[string]func() error{
		"Reading System Configurations": in.StageSysConfig,
		"Verify System Memory":          in.StageVerifyMemory,
		"Installing System":             in.StageInstall,
	}

	return &in
}
