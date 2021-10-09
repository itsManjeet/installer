package installer

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/gotk3/gotk3/glib"
)

func (instlr *Installer) askUser() error {
	return exec.Command("/bin/zenity", "--question").Run()
}

func (instlr *Installer) startInstall() {
	instlr.verifySystemMemory()
	instlr.checkRootDevice()
	instlr.checkPartSize()
	instlr.checkBootMode()

	if instlr.IsEfi {
		instlr.searchEfiDevice()
	} else {
		instlr.checkBootDevice()
	}

	instlr.verifySystemImage()

	instlr.startProcess("User confirmation")
	if err := instlr.askUser(); err != nil {
		instlr.stateProcess("denied", false)
		instlr.checkError(fmt.Errorf("user denied, aborting"))
	}
	instlr.stateProcess("", true)

	log.Println("Creating temporary dir")
	workdir, err := os.MkdirTemp(os.TempDir(), "installer-*")
	instlr.checkError(err)

	log.Printf("Mounting %s on %s\n", instlr.RootDevice, workdir)
	instlr.checkError(exec.Command("/bin/mount", instlr.RootDevice, workdir).Run())

	instlr.cleanup = func() {
		log.Println("unmounting", workdir)
		syscall.Unmount(workdir, 0)
	}

	instlr.installImage(workdir)
	instlr.installBootloader(workdir)

	instlr.configureBootloader(workdir)

	glib.IdleAdd(func() {
		instlr.stack.SetVisibleChildName("successPage")
	})

	instlr.cleanup()
}

func (instlr *Installer) displayText(mesg string) {
	glib.IdleAdd(func() {
		str, err := instlr.statusBuffer.GetText(instlr.statusBuffer.GetStartIter(), instlr.statusBuffer.GetEndIter(), false)
		instlr.checkError(err)
		str += mesg
		instlr.statusBuffer.SetText(str)
	})
}

func (instlr *Installer) startProcess(mesg string) {
	instlr.displayText(fmt.Sprintf("○ %s - ", mesg))
	log.Println(mesg)
	// FOR BETTER UX...
	time.Sleep(time.Millisecond * 200)
}

func (instlr *Installer) updateProgress(prog float64) {
	glib.IdleAdd(func() {
		instlr.progressBar.SetFraction(prog)
	})
}

func (instlr *Installer) stateProcess(mesg string, success bool) {
	if success {
		instlr.displayText(mesg + "\t✓\n")
	} else {
		instlr.displayText(mesg + "\t✘\n")
	}
}
