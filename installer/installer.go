package installer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const CONFIG_FILE = "/etc/installer.json"

type Installer struct {
	ui *gtk.Builder

	stack  *gtk.Stack
	Window *gtk.Window

	statusBuffer *gtk.TextBuffer
	progressBar  *gtk.ProgressBar

	RootDevice string `json:"root"`
	BootDevice string `json:"boot"`
	RootUUID   string

	cleanup func()

	IsEfi bool `json:"is-efi"`

	SystemImage   string `json:"system-image"`
	SystemVersion string `json:"system-version"`

	RootDeviceLabel string `json:"root-label"`
	BootDeviceLabel string `json:"boot-label"`

	RootFSType string `json:"root-fs-type"`
	BootFSType string `json:"boot-fs-type"`

	MinimumRam  int `json:"mimimum-ram"`
	MinimumDisk int `json:"mimimum-disk"`

	signals map[string]interface{}
}

func Init(ui *gtk.Builder) *Installer {
	var installer Installer

	if _, err := os.Stat(CONFIG_FILE); os.IsNotExist(err) {
		installer.checkError(fmt.Errorf("no configuration file found %v", err))

	} else {
		if data, err := os.ReadFile(CONFIG_FILE); err == nil {
			json.Unmarshal(data, &installer)
		} else {
			installer.checkError(err)
		}
	}

	installer.ui = ui
	installer.Window = installer.getWidget("mainWindow").(*gtk.Window)
	installer.stack = installer.getWidget("stack").(*gtk.Stack)

	installer.statusBuffer = installer.getWidget("statusBuffer").(*gtk.TextBuffer)
	installer.progressBar = installer.getWidget("progressBar").(*gtk.ProgressBar)

	installer.signals = map[string]interface{}{
		"onDestroy":            installer.onDestroy,
		"onContinueBtnClicked": installer.onContinueBtnClicked,
		"onRebootBtnClicked":   installer.onRebootBtnClicked,
	}

	installer.ui.ConnectSignals(installer.signals)
	installer.Window.SetPosition(gtk.WIN_POS_CENTER_ALWAYS)

	return &installer
}

func (inst *Installer) getWidget(id string) glib.IObject {
	obj, err := inst.ui.GetObject(id)
	inst.checkError(err)
	return obj
}

func (instlr Installer) checkError(err error) {
	if err != nil {
		log.Println("EE", err.Error())
		exec.Command("/bin/zenity", "--error", "--text", err.Error()).Run()
		if instlr.cleanup != nil {
			instlr.cleanup()
		}
		os.Exit(1)
	}
}
