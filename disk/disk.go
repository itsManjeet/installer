package disk

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
