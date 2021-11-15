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
	PartLabel string        `json:"partlabel"`
	Childrens []BlockDevice `json:"children"`
}

type BlockDevices struct {
	Devices []BlockDevice `json:"blockdevices"`
}

func ListPartitions() (*BlockDevices, error) {
	log.Println("Listing block devices")
	data, err := exec.Command("lsblk", "-J", "-O").Output()
	if err != nil {
		return nil, err
	}
	var devices BlockDevices
	if err := json.Unmarshal(data, &devices); err != nil {
		return nil, err
	}

	return &devices, nil
}
