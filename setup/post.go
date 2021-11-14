package setup

import (
	"log"
	"os"
)

func (setup *Setup) StagePost() error {
	autologin := "/etc/lightdm/lightdm.conf.d/10-auto-login.conf"
	if _, err := os.Stat(autologin); err == nil {
		log.Println("removing autologin")
		if !setup.IsDebug() {
			if err := os.Remove(autologin); err != nil {
				return err
			}
		}

	}
	return nil
}
