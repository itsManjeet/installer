package app

import "os"

func (app *App) IsDebug() bool {
	return len(os.Getenv("SYS_SETUP_DEBUG")) != 0
}
