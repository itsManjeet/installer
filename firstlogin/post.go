package firstlogin

import (
	"log"
	"os/exec"
	"time"

	"github.com/gotk3/gotk3/glib"
)

func (f *FirstLogin) startPost() {

	updateProgress := func(mesg string, prog float64) {
		glib.IdleAdd(func() {

			f.postProgress.SetText(mesg)
			f.postProgress.SetFraction(prog)
		})

		time.Sleep(time.Millisecond * 150)
	}
	updateProgress("Executing post process", 0.1)

	updateProgress("Updating grub configurations", 0.25)
	if err := exec.Command("update-grub").Run(); err != nil {
		log.Println("Failed to update grub, error: ", err)
	}

	updateProgress("Complete configurations", 1.0)

	glib.IdleAdd(func() {
		f.Window.SetPageComplete(f.ProcessPage, true)
		f.Window.SetCurrentPage(4)
	})
}
