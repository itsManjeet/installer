package app

import (
	"fmt"
	"log"
	"time"

	"github.com/gotk3/gotk3/glib"
)

func (app *App) DisplayText(mesg string) {
	glib.IdleAdd(func() {
		str, err := app.StatusBuffer.GetText(app.StatusBuffer.GetStartIter(), app.StatusBuffer.GetEndIter(), false)
		app.checkError(err)
		str += mesg
		app.StatusBuffer.SetText(str)
	})
}

func (app *App) StartProcess(mesg string) {
	app.DisplayText(fmt.Sprintf("○ %s - ", mesg))
	log.Println(mesg)
	// FOR BETTER UX...
	time.Sleep(time.Millisecond * 200)
}

func (app *App) UpdateProgress(prog float64) {
	glib.IdleAdd(func() {
		app.ProgressBar.SetFraction(prog)
	})
}

func (app *App) StateProcess(mesg string, success bool) {
	if success {
		app.DisplayText(mesg + "\t✓\n")
	} else {
		app.DisplayText(mesg + "\t✘\n")
	}
}
