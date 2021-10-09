package installer

import (
	"encoding/json"
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

func (instlr *Installer) getBlockDevices() (*BlockDevices, error) {
	data, err := exec.Command("lsblk", "-J", "-O").Output()
	if err != nil {
		return nil, err
	}
	var blockDevices BlockDevices
	err = json.Unmarshal(data, &blockDevices)
	return &blockDevices, err
}
