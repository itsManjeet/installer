package app

import (
	"os/exec"
	"strings"

	"github.com/gotk3/gotk3/gtk"
)

func (app *App) ErrorDialog(mesg string) {
	dialog := gtk.MessageDialogNew(app.Window, gtk.DIALOG_DESTROY_WITH_PARENT, gtk.MESSAGE_ERROR, gtk.BUTTONS_CLOSE, mesg)
	dialog.Run()

	dialog.Destroy()
}

func (app *App) DialogInputText(text string) (string, error) {
	data, err := exec.Command("/bin/zenity", "--entry", "--text="+text).Output()
	return strings.TrimSuffix(string(data), "\n"), err
}

func (app *App) DialogPassword(text string) (string, error) {
	data, err := exec.Command("/bin/zenity", "--password", "--text="+text).Output()
	return strings.TrimSuffix(string(data), "\n"), err
}

func (app *App) DialogAsk(text string) bool {
	return exec.Command("/bin/zenity", "--question", "--title='Confirm'", "--text="+text).Run() == nil
}

func (app *App) DialogList(mesg string, col string, options ...string) (string, error) {
	args := []string{
		"--list",
		"--title=" + mesg,
		"--column=" + col,
	}

	args = append(args, options...)
	data, err := exec.Command("/bin/zenity", args...).Output()
	return strings.TrimSuffix(string(data), "\n"), err
}

func (app *App) DialogUserLogin(text string) (string, string, error) {
	data, err := exec.Command("/bin/zenity", "--password", "--username", "--text="+text).Output()
	data_ := strings.Split(strings.TrimSuffix(string(data), "\n"), "|")
	return data_[0], data_[1], err
}
