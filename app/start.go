package app

import (
	"log"
	"time"

	"github.com/gotk3/gotk3/glib"
)

func (app *App) Start() {

	progress := 0.0
	steps := 1.0 / float64(len(app.stages_index))

	for _, id := range app.stages_index {
		app.StartProcess(id)
		time.Sleep(time.Second * 1)
		if err := app.stages[id](); err != nil {
			app.StateProcess("", false)
			app.checkError(err)
		}

		progress += steps
		app.StateProcess("", true)
		app.UpdateProgress(progress)
		log.Println("updating progress", progress)
	}

	glib.IdleAdd(func() {
		app.Stack.SetVisibleChildName("successPage")
	})

}
