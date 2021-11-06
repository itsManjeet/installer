package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/rlxos/installer/app"
	"github.com/rlxos/installer/installer"
	"github.com/rlxos/installer/setup"
)

const (
	APPID = "dev.rlxos.setup"
)

var (
	commandLine string
)

func main() {
	application, err := gtk.ApplicationNew(APPID, glib.APPLICATION_FLAGS_NONE)
	checkError(err)

	application.Connect("startup", func() {
		log.Println("Starting up", APPID)
		_cmdline, err := ioutil.ReadFile("/proc/cmdline")
		commandLine = string(_cmdline)
		checkError(err)
		/// TODO pre configurations
	})

	application.Connect("activate", func() {
		log.Println("Activating", APPID)

		builder, err := gtk.BuilderNewFromString(app.UI)
		checkError(err)

		if strings.Contains(string(commandLine), "iso=1") || os.Getenv("SYS_SETUP_MODE") == "installer" {
			installer := installer.Init(builder)
			installer.Window.ShowAll()
			application.AddWindow(installer.Window)
		} else {
			setup := setup.Init(builder)
			setup.Window.ShowAll()
			application.AddWindow(setup.Window)

		}

	})

	application.Connect("shutdown", func() {
		log.Println("Shutting down", APPID)
	})

	os.Exit(application.Run(os.Args))
}

func checkError(err error) {
	if err != nil {
		log.Println("EE", err.Error())
		exec.Command("/bin/zenity", "--error", "--text", err.Error()).Run()
		os.Exit(1)
	}
}
