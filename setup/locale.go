package setup

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func (setup *Setup) StageLocale() error {
	for {
		selected, err := setup.DialogList("Select Locale", "locale", SUPPORTED_LOCALE...)
		if err != nil {
			return err
		}
		if len(selected) != 0 {
			log.Println("LOCALE", selected)
			return setup.SetupLocale(selected)
		}
	}
}

func (setup *Setup) SetupLocale(locale string) error {
	if !setup.IsDebug() {
		if err := os.MkdirAll("/usr/lib/locale", 0755); err != nil {
			return err
		}
	}

	data := strings.Split(locale, ".")
	log.Println("Generating locale", data[0]+":"+data[1])
	if !setup.IsDebug() {
		if err := exec.Command("localedef", "-i", data[0], "-f", data[1], locale).Run(); err != nil {
			return err
		}
	}

	log.Println("writing locale to configuration")
	if !setup.IsDebug() {
		if err := ioutil.WriteFile("/etc/locale.conf", []byte("LANG="+locale), 0); err != nil {
			return err
		}
	}

	return nil
}
