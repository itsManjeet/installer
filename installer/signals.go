package installer

import "os/exec"

func (instlr *Installer) onDestroy() {
	// TODO: check if installation is complete or not
}

func (instlr *Installer) onContinueBtnClicked() {
	instlr.stack.SetVisibleChildName("installPage")
	go instlr.startInstall()
}

func (instlr *Installer) onRebootBtnClicked() {
	go exec.Command("reboot").Run()
}
