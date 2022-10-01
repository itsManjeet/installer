package installer

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type User struct {
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
}

type Installer struct {
	Root        string `yaml:"root"`
	Workdir     string `yaml:"work-dir"`
	SystemImage string `yaml:"system-image"`
	Kernel      string `yaml:"kernel"`
	Boot        string `yaml:"boot"`
	IsEfi       bool   `yaml:"is-efi"`
	Locale      string `yaml:"locale"`
	TimeZone    string `yaml:"timezone"`

	rootUUID string

	Users []User `yaml:"users"`
}

func LoadConfig(filepath string) (*Installer, error) {
	buffer, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var installer Installer
	if err := yaml.Unmarshal(buffer, &installer); err != nil {
		return nil, err
	}
	return &installer, nil
}
