package installer

import (
	"encoding/json"
	"log"
	"os/exec"
)

type BlockDevice struct {
	Name      string        `json:"name"`
	Type      string        `json:"fstype"`
	Path      string        `json:"Path"`
	Size      string        `json:"size"`
	Label     string        `json:"label"`
	UUID      string        `json:"uuid"`
	Childrens []BlockDevice `json:"children"`
}

type BlockDevices struct {
	Devices []BlockDevice `json:"blockdevices"`
}

func (in *Installer) StageSysConfig() error {
	log.Println("Listing block devices")
	data, err := exec.Command("lsblk", "-J", "-O").Output()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &in.devices)
}
