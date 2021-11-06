package setup

import (
	"log"
	"os/exec"
	"strings"
)

func (setup *Setup) StageAccount() error {
	for {
		username, password, err := setup.DialogUserLogin("Admin Account Details")
		if err != nil {
			return err
		}
		if len(username) != 0 && len(password) != 0 {
			log.Println("Creating user", username)
			return setup.CreateAccount(username, password)
		}
	}
}

func (setup *Setup) CreateAccount(username, password string) error {

	cmd := exec.Command("openssl", "passwd", "-1", password)
	passwordBytes, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	password = strings.TrimSpace(string(passwordBytes))

	log.Println("Creating Admin User")
	if !setup.IsDebug() {
		if err := exec.Command("useradd", "-m", "-g", "users", "-G", "adm", username, "-p", password).Run(); err != nil {
			return err
		}

	}
	return nil
}
