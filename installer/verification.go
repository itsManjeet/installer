package installer

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	"unicode"

	"github.com/dustin/go-humanize"
	"github.com/gotk3/gotk3/glib"
)

// #include <unistd.h>
import "C"

func (ins *Installer) SystemMemory() uint64 {
	return uint64(C.sysconf(C._SC_PHYS_PAGES) * C.sysconf(C._SC_PAGE_SIZE))
}

func (ins *Installer) StartVerification() {
	updateProgress := func(mesg string, prog float64) {
		glib.IdleAdd(func() {
			ins.verifyProgress.SetText(mesg)
			ins.verifyProgress.SetFraction(prog)
		})

		time.Sleep(time.Millisecond * 150)
	}

	updateProgress("Verifying System Memory", 0.05)
	SystemMemory = ins.SystemMemory()
	if SystemMemory < MINIMUM_MEMORY {
		updateProgress(fmt.Sprintf("Failed, System memory is below mimimum %s requirements", humanize.Bytes(MINIMUM_MEMORY)), 1.0)
		return
	}
	updateProgress(fmt.Sprintf("System Memory Verification Pass: %s", humanize.Bytes(SystemMemory)), 0.1)

	updateProgress("Checking bootloader type", 0.15)
	if _, err := os.Stat("/sys/firmware/efi/efivars"); err == nil {
		BootLoader = BootLoaderTypeEfi
	} else {
		BootLoader = BootLoaderTypeLegacy
	}
	updateProgress(fmt.Sprintf("Found Bootloader %s Type", BootLoader.String()), 0.2)

	updateProgress("Reading paritions ", 0.25)

	getParition := func() (string, error) {
		partitionsLen := 0
		devices, err := ListPartitions()
		if err != nil {
			return "", err
		}
		for _, disks := range devices.Devices {
			for _, part := range disks.Childrens {
				partitionsLen++
				if part.Label == PARITION_LABEL {
					PartitionUUID = part.UUID
					return part.Path, nil
				}
			}
		}

		if partitionsLen == 0 {
			return "", errors.New("no parition Found")
		}

		return "", errors.New("failed to get partition with label " + PARITION_LABEL)
	}

	for {
		var err error
		PartitionPath, err = getParition()
		if err != nil {
			updateProgress(err.Error(), 0.25)
			if err := exec.Command("gparted").Run(); err != nil {
				updateProgress("Failed to launch gparted: "+err.Error(), 1.0)
				return
			}
		} else {
			break
		}
	}

	updateProgress("Found Partition: "+PartitionPath, 0.3)

	updateProgress("Checking system Image: "+IMAGE_PATH, 0.35)
	if !ins.IsDebug(APPID) {
		if _, err := os.Stat(IMAGE_PATH); err != nil {
			updateProgress("Failed to verify system Image: "+IMAGE_PATH+", "+err.Error(), 1.0)
			return
		}
	}

	if BootLoader == BootLoaderTypeEfi {
		getBootParition := func() (string, error) {
			partitionsLen := 0
			devices, err := ListPartitions()
			if err != nil {
				return "", err
			}
			for _, disks := range devices.Devices {
				for _, part := range disks.Childrens {
					partitionsLen++
					if part.PartLabel == "EFI" {
						return part.Path, nil
					}
				}
			}

			if partitionsLen == 0 {
				return "", errors.New("no parition Found")
			}

			return "", errors.New("failed to get partition with label EFI")
		}

		for {
			var err error
			BootDevice, err = getBootParition()
			if err != nil {
				updateProgress(err.Error(), 0.25)
				if err := exec.Command("gparted").Run(); err != nil {
					updateProgress("Failed to launch gparted: "+err.Error(), 1.0)
					return
				}
			} else {
				break
			}
		}

	} else if BootLoader == BootLoaderTypeLegacy {
		BootDevice = strings.TrimRightFunc(PartitionPath, func(r rune) bool {
			return unicode.IsDigit(r)
		})
	}

	updateProgress("Found System Image: "+IMAGE_PATH, 0.4)

	updateProgress("System Checkup Complete", 1.0)

	glib.IdleAdd(func() {
		ins.SetupConfirm()
		ins.Window.SetPageComplete(ins.VerifyPage, true)
	})
}
