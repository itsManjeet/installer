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

		var window *gtk.Window

		if strings.Contains(string(commandLine), "iso=1") || os.Getenv("SYS_SETUP_MODE") == "installer" {
			installer := installer.Init(builder)
			window = installer.Window
		} else {
			setup := setup.Init(builder)
			window = setup.Window
		}

		window.Fullscreen()
		window.ShowAll()
		application.AddWindow(window)

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
