package setup

import (
	"io/ioutil"
	"log"
	"os"
)

func (setup *Setup) StageTimezone() error {

	zoneDir, err := ioutil.ReadDir("/usr/share/zoneinfo")
	if err != nil {
		return err
	}

	timezones := make([]string, 0)
	for _, a := range zoneDir {
		if !a.IsDir() {
			continue
		}
		dir, err := ioutil.ReadDir("/usr/share/zoneinfo/" + a.Name())
		if err != nil {
			continue
		}

		for _, d := range dir {
			timezones = append(timezones, a.Name()+"/"+d.Name())
		}
	}

	for {
		selected, err := setup.DialogList("Select Timezone", "timezone", timezones...)
		if err != nil {
			return err
		}
		if len(selected) != 0 {
			log.Println("TIMEZONE", selected)
			if !setup.IsDebug() {
				return setup.SetupTimeZone(selected)
			} else {
				return nil
			}

		}
	}
}

func (setup *Setup) SetupTimeZone(zone string) error {
	os.Remove("/etc/localtime")
	return os.Symlink("/usr/share/zoneinfo/"+zone, "/etc/localtime")
}
