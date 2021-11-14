package app

import (
	"errors"
	"log"
	"os"
	"os/exec"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type App struct {
	ui *gtk.Builder

	ID          string
	Title       string
	WelcomeMesg string

	ProcessTitle string
	ProcessMesg  string

	SuccessMesg string
	SuccessBtn  string

	Stack  *gtk.Stack
	Window *gtk.Window

	welcomeMesgLbl *gtk.Label
	processNameLbl *gtk.Label
	processDescLbl *gtk.Label
	successMesgLbl *gtk.Label
	successBtn     *gtk.Button

	StatusBuffer *gtk.TextBuffer
	ProgressBar  *gtk.ProgressBar

	stages_index []string
	stages       map[string]func() error

	SuccessHandler func()

	maxMessageSize int
}

func Create(ui *gtk.Builder) *App {
	var app App

	app.ui = ui

	app.Window = app.getWidget("mainWindow").(*gtk.Window)
	app.Window.SetTitle(app.Title)

	app.Stack = app.getWidget("stack").(*gtk.Stack)

	app.StatusBuffer = app.getWidget("statusBuffer").(*gtk.TextBuffer)
	app.ProgressBar = app.getWidget("progressBar").(*gtk.ProgressBar)
	app.welcomeMesgLbl = app.getWidget("welcomeMesgLbl").(*gtk.Label)
	app.processNameLbl = app.getWidget("processNameLbl").(*gtk.Label)
	app.processDescLbl = app.getWidget("processDescLbl").(*gtk.Label)
	app.successMesgLbl = app.getWidget("successDescLbl").(*gtk.Label)
	app.successBtn = app.getWidget("successBtn").(*gtk.Button)

	signals := map[string]interface{}{
		"onDestroy":            app.onDestroy,
		"onSuccessBtnClicked":  app.onSuccessBtnClicked,
		"onContinueBtnClicked": app.onContinueBtnClicked,
	}

	app.stages = map[string]func() error{
		"Initializing": func() error {
			return errors.New("not yet implemented")
		},
	}

	app.ui.ConnectSignals(signals)

	app.Window.SetPosition(gtk.WIN_POS_CENTER_ALWAYS)
	return &app
}

func (app *App) Initialize() {
	app.successBtn.SetLabel(app.SuccessBtn)
	app.successMesgLbl.SetText(app.SuccessMesg)
	app.processDescLbl.SetText(app.ProcessMesg)
	app.processNameLbl.SetText(app.ProcessTitle)
	app.welcomeMesgLbl.SetText(app.WelcomeMesg)

}

func (app *App) getWidget(id string) glib.IObject {
	obj, err := app.ui.GetObject(id)
	app.checkError(err)
	return obj
}

func (applr App) checkError(err error) {
	if err != nil {
		log.Println("EE", err.Error())
		exec.Command("/bin/zenity", "--error", "--text", err.Error()).Run()
		os.Exit(1)
	}
}

func (app *App) AddStage(id string, callback func() error) {

	if app.maxMessageSize < len(id) {
		app.maxMessageSize = len(id)
	}

	if app.stages_index == nil {
		app.stages_index = make([]string, 0)
	}
	app.stages_index = append(app.stages_index, id)

	if app.stages == nil {
		app.stages = make(map[string]func() error)
	}
	app.stages[id] = callback
}
