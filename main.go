package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/rlxos/installer/firstlogin"
	"github.com/rlxos/installer/installer"
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

		window, err := gtk.AssistantNew()
		checkError(err)

		window.SetDefaultSize(800, 600)

		if strings.Contains(commandLine, "iso") || os.Getenv("SYS_SETUP_MODE") == "installer" {
			if err := installer.Setup(window); err != nil {
				checkError(err)
			}
		} else {
			if err := firstlogin.Setup(window); err != nil {
				checkError(err)
			}
		}

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
