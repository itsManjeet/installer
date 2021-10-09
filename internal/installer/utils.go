package installer

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/mem"
)

// return total system memory
func (instlr Installer) verifySystemMemory() {
	instlr.startProcess("Checking system ram")
	instlr.updateProgress(0.05)

	memstat, err := mem.VirtualMemory()
	instlr.checkError(err)

	instlr.stateProcess(humanize.Bytes(memstat.Total), memstat.Total > uint64(instlr.MinimumRam))
	if memstat.Total < uint64(instlr.MinimumRam) {
		instlr.checkError(fmt.Errorf("minimum ram/memory requirement check failed"))
	}
}

func (instlr *Installer) checkBootMode() {
	instlr.updateProgress(0.17)
	instlr.startProcess("Detecting Bootloader Mode")
	if _, err := os.Stat("/sys/firmware/efi/efivars"); os.IsNotExist(err) {
		instlr.stateProcess("legacy", true)
		instlr.IsEfi = false
	} else {
		instlr.stateProcess("uefi", true)
		instlr.IsEfi = true
	}
}

func (instlr *Installer) checkBootDevice() {
	instlr.updateProgress(0.17)
	for len(instlr.BootDevice) == 0 {
		blockdevices, err := instlr.getBlockDevices()
		instlr.checkError(err)
		instlr.startProcess("Check disk for legacy bootloader")

		zenityStr := []string{"--list", "--title=\"Boot Disk\"", "--column=\"Disk\""}
		for _, i := range blockdevices.Devices {
			zenityStr = append(zenityStr, i.Path)
		}

		out, err := exec.Command("zenity", zenityStr...).Output()
		instlr.checkError(err)
		instlr.BootDevice = strings.Trim(string(out), "\n")
	}
	instlr.stateProcess(instlr.BootDevice, true)
}

// check rlxos root device
func (instlr *Installer) searchEfiDevice() {
	instlr.updateProgress(0.15)
	for len(instlr.BootDevice) == 0 {
		instlr.startProcess("Search EFI partition with LABEL '" + instlr.BootDeviceLabel + "'")
		blockdevices, err := instlr.getBlockDevices()
		instlr.checkError(err)

		for _, disk := range blockdevices.Devices {
			for _, part := range disk.Childrens {
				if part.Label == instlr.BootDeviceLabel {
					instlr.BootDevice = part.Path
					break
				}
			}
			if len(instlr.BootDevice) != 0 {
				break
			}
		}

		if len(instlr.BootDevice) == 0 {
			instlr.stateProcess("", false)
			instlr.checkError(exec.Command("gparted").Run())
		}
	}

	instlr.stateProcess(instlr.BootDevice, true)
}

// check rlxos root device
func (instlr *Installer) checkRootDevice() {
	instlr.updateProgress(0.15)
	for len(instlr.RootDevice) == 0 {
		instlr.startProcess("Search rlxos partition with LABEL '" + instlr.RootDeviceLabel + "'")
		blockdevices, err := instlr.getBlockDevices()
		instlr.checkError(err)

		for _, disk := range blockdevices.Devices {
			for _, part := range disk.Childrens {
				if part.Label == instlr.RootDeviceLabel {
					instlr.RootDevice = part.Path
					break
				}
			}
			if len(instlr.RootDevice) != 0 {
				break
			}
		}

		if len(instlr.RootDevice) == 0 {
			instlr.stateProcess("", false)
			instlr.checkError(exec.Command("gparted").Run())
		}
	}

	instlr.stateProcess(instlr.RootDevice, true)
}

func (instlr *Installer) checkPartSize() {
	instlr.startProcess("Checking root device size")
	instlr.updateProgress(0.16)
	sizefile := path.Join("/sys/class/block/", path.Base(instlr.RootDevice), "size")
	if _, err := os.Stat(sizefile); err == os.ErrNotExist {
		instlr.checkError(err)
	}
	data, err := os.ReadFile(sizefile)
	instlr.checkError(err)
	size, err := strconv.Atoi(strings.Trim(string(data), "\n"))
	instlr.checkError(err)
	size *= 512

	instlr.stateProcess(humanize.Bytes(uint64(size)), size > instlr.MinimumDisk)
	if size < instlr.MinimumDisk {
		instlr.checkError(fmt.Errorf("minimum disk space requirement check failed"))
	}

}

func (instlr *Installer) verifySystemImage() {
	instlr.startProcess("Verifying system image")
	instlr.updateProgress(0.20)
	if _, err := os.Stat(instlr.SystemImage); os.IsNotExist(err) {
		instlr.stateProcess(instlr.SystemImage, false)
		instlr.checkError(fmt.Errorf("missing system image: %s", instlr.SystemImage))
	} else {
		instlr.stateProcess(instlr.SystemImage, true)
	}

}

func (instlr *Installer) installImage(workdir string) {

	instlr.startProcess("Installing system image")
	instlr.updateProgress(0.25)

	log.Println("Creating required directories")
	sysdir := path.Join(workdir, "rlxos", "system")
	instlr.checkError(os.MkdirAll(sysdir, 0755))

	DEST := path.Join(sysdir, instlr.SystemVersion)
	log.Printf("Copying system image %s -> %s", instlr.SystemImage, DEST)

	srcfile, err := os.OpenFile(instlr.SystemImage, os.O_RDONLY, 0)
	instlr.checkError(err)
	defer srcfile.Close()

	destfile, err := os.Create(DEST)
	instlr.checkError(err)
	defer destfile.Close()

	_, err = io.Copy(destfile, srcfile)
	instlr.checkError(err)

	instlr.checkError(destfile.Sync())

	instlr.stateProcess("", true)
	instlr.updateProgress(0.6)
}

func (instlr *Installer) installBootloader(workdir string) {
	bootdir := path.Join(workdir, "boot")
	instlr.checkError(os.MkdirAll(path.Join(bootdir, "grub"), 0755))
	instlr.startProcess("Installing bootloader")

	args := []string{"--root-directory=" + workdir, "--boot-directory=" + bootdir, "--recheck"}

	if instlr.IsEfi {
		efidir := path.Join(bootdir, "efi")
		instlr.checkError(os.MkdirAll(efidir, 0755))

		log.Println("mounting boot device ", instlr.BootDevice)
		instlr.checkError(exec.Command("/bin/mount", instlr.BootDevice, efidir).Run())
		args = append(args, "--bootloader-id=rlx")
	} else {
		args = append(args, instlr.BootDevice)
	}

	instlr.checkError(exec.Command("/bin/grub-install", args...).Run())
	instlr.stateProcess("", true)
	instlr.updateProgress(0.9)
}

func (instlr *Installer) configureBootloader(workdir string) {

	instlr.startProcess("Configuring bootloader")

	blockdevices, err := instlr.getBlockDevices()
	instlr.checkError(err)

	var uuid string
	for _, disk := range blockdevices.Devices {
		for _, part := range disk.Childrens {
			if part.Path == instlr.RootDevice {
				uuid = part.UUID
				break
			}
		}
		if len(uuid) != 0 {
			break
		}
	}

	if len(uuid) == 0 {
		instlr.checkError(fmt.Errorf("internal error, failed to get parition uuid for " + instlr.RootDevice))
	}

	log.Println("UUID: ", uuid)

	grubCFG := "insmod part_gpt\ninsmod part_msdos\ninsmod all_video\ntimeout=5\ndefault='rlxos initial setup'\nmenuentry 'rlxos initial setup' {\n\tinsmod gzio\n\tinsmod ext2\n\tlinux /boot/vmlinuz root=UUID=" + uuid + " system=" + instlr.SystemVersion + "\n" + "\tinitrd /boot/initrd\n}"
	instlr.checkError(os.WriteFile(path.Join(workdir, "boot", "grub", "grub.cfg"), []byte(grubCFG), 0755))

	instlr.checkError(exec.Command("/bin/cp", path.Join(path.Dir(instlr.SystemImage), "boot", "vmlinuz"), path.Join(workdir, "boot", "vmlinuz")).Run())
	instlr.checkError(exec.Command("/bin/cp", path.Join(path.Dir(instlr.SystemImage), "boot", "initrd"), path.Join(workdir, "boot", "initrd")).Run())
	instlr.stateProcess("", true)
	instlr.updateProgress(1.0)
}
