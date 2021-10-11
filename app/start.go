package app

import (
	"time"

	"github.com/gotk3/gotk3/glib"
)

func (app *App) Start() {

	progress := 0.0
	steps := float64(1.0 / len(app.Stages))
	for mesg, fn := range app.Stages {
		app.StartProcess(mesg)
		time.Sleep(time.Second * 1)
		if err := fn(); err != nil {
			app.StateProcess("", false)
			app.checkError(err)
		}

		progress += steps
		app.StateProcess("", true)
	}

	glib.IdleAdd(func() {
		app.Stack.SetVisibleChildName("successPage")
	})

}
