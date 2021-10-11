package app

func (app *App) onDestroy() {
	// TODO: check if installation is complete or not
}

func (app *App) onContinueBtnClicked() {
	app.Stack.SetVisibleChildName("installPage")
	go app.Start()
}

func (app *App) onSuccessBtnClicked() {
	a, _ := app.Window.GetApplication()
	a.Quit()
}
