package firstlogin

import (
	"io/ioutil"
	"path"
)

func (f *FirstLogin) getTimeZoneList() []string {
	dir, err := ioutil.ReadDir(TIMEZONE_DIR)
	if err != nil {
		return nil
	}

	list := []string{}

	for _, cont := range dir {
		if cont.IsDir() {
			place_dir, err := ioutil.ReadDir(path.Join(TIMEZONE_DIR, cont.Name()))
			if err != nil {
				return nil
			}
			for _, pla := range place_dir {
				list = append(list, cont.Name()+"/"+pla.Name())
			}
		}
	}

	return list
}
