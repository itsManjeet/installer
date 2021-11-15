package installer

const (
	MINIMUM_MEMORY = 2147483648
	PARITION_LABEL = "rlxos"
	APPID          = "SYS_INSTALLER"
	IMAGE_PATH     = "/run/iso/rootfs.img"
	VERSION        = "2200"
)

type BootLoaderType int

const (
	BootLoaderTypeEfi BootLoaderType = iota
	BootLoaderTypeLegacy
)

var (
	BootLoader    BootLoaderType
	SystemMemory  uint64
	PartitionPath string
	PartitionUUID string
	BootDevice    string
)

func (b BootLoaderType) String() string {
	switch b {
	case BootLoaderTypeEfi:
		return "EFI"
	case BootLoaderTypeLegacy:
		return "Legacy"
	}

	return "<Unsupported Bootloader Type>"
}
