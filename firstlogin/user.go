package firstlogin

import (
	"log"
	"os/exec"
	"strings"
	"unicode"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func (f *FirstLogin) createUser(userid, passwd string) {

	showError := func(mesg string) {
		glib.IdleAdd(func() {
			f.Window.SetPageComplete(f.UserAccountPage, false)

			dialog := gtk.MessageDialogNew(f.Window, gtk.DIALOG_DESTROY_WITH_PARENT|gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_CLOSE, mesg)
			dialog.Run()
			dialog.Destroy()
		})
	}

	checkEntry := func(data string, entry string) bool {
		if len(data) == 0 {
			showError("Please Input value in " + entry)
			return false
		}

		if strings.ContainsRune(data, ' ') {
			showError(entry + " should not contain any spaces")
			return false
		}

		if entry == "username" {
			if unicode.IsDigit(rune(data[0])) {
				showError(entry + " should not starts with digit")
				return false
			}
		}

		return true
	}

	if !checkEntry(userid, "username") {
		return
	}

	if !checkEntry(passwd, "password") {
		return
	}

	if !f.IsDebug(APPID) {
		if err := exec.Command("useradd", "-G", "adm", "-g", "users", "-m", userid).Run(); err != nil {
			showError(err.Error())
			return
		}
		if err := exec.Command("sh", "-c", "echo -e \""+passwd+"\n"+passwd+"\" | passwd "+userid).Run(); err != nil {
			showError(err.Error())
			return
		}
	} else {
		log.Println("creating user: ", userid)
	}

	glib.IdleAdd(func() {
		f.Window.SetPageComplete(f.UserAccountPage, true)
		f.Window.SetCurrentPage(3)
	})

}
