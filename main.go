package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/rlxos/installer/internal/installer"
)

func main() {
	application, err := gtk.ApplicationNew("dev.rlxos.installer", glib.APPLICATION_FLAGS_NONE)
	checkError(err)

	application.Connect("startup", func() {
		log.Println("Starting up installer")
		/// TODO pre configurations
	})

	application.Connect("activate", func() {
		log.Println("Activating installer")

		builder, err := gtk.BuilderNewFromString(UI)
		checkError(err)

		instlr := installer.Init(builder)
		instlr.Window.ShowAll()
		application.AddWindow(instlr.Window)
	})

	application.Connect("shutdown", func() {
		log.Println("Shutting down installer")
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
