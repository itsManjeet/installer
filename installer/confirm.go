package installer

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

func (ins *Installer) SetupConfirm() {
	ins.confirmBuffer.SetText(fmt.Sprintf(`
System Memory      :    %s
BootLoader Type   :    %s
Root Partition         :    %s
System Image         :    %s
Boot Device              :    %s
Version                      :    %s`, humanize.Bytes(SystemMemory), BootLoader.String(), PartitionPath, IMAGE_PATH, BootDevice, VERSION))

	ins.Window.SetPageComplete(ins.ConfirmPage, true)
}
