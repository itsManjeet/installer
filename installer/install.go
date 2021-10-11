package installer

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path"
)

func (in *Installer) StageInstall() error {
	root_disk := ""
	var err error
	for {

		log.Println("Searching root device with label 'rlxos'")
		root_disk, err = in.searchLabel("rlxos")
		if err != nil {
			return err
		}

		if !in.DialogAsk("Please confirm if you want to install rlxOS on " + root_disk) {
			return errors.New("user cancelled")
		}

		if len(root_disk) != 0 {
			return in.installImage("/run/iso/rootfs.img", root_disk)
		}

	}
}

func (in *Installer) searchLabel(lbl string) (string, error) {
	for _, disk := range in.devices.Devices {
		for _, part := range disk.Childrens {
			if part.Label == lbl {
				return part.Path, nil
			}
		}
	}

	if err := exec.Command("gparted").Run(); err != nil {
		return "", err
	}

	if err := in.StageSysConfig(); err != nil {
		return "", err
	}

	return "", nil
}

func (in *Installer) installImage(imagePath, deviceNode string) error {
	tempdir, err := os.MkdirTemp("", "installer-*")
	if err != nil {
		return err
	}

	in.tempdir = tempdir

	log.Printf("Mounting %s %s\n", deviceNode, tempdir)
	if err := exec.Command("mount", deviceNode, tempdir).Run(); err != nil {
		return err
	}

	log.Println("Creating directories")
	if err := os.MkdirAll(path.Join(tempdir, "rlxos", "system"), 0755); err != nil {
		return err
	}

	log.Println("Copying squash")
	if err := exec.Command("cp", imagePath, path.Join(tempdir, "rlxos", "system", VERSION)).Run(); err != nil {
		return err
	}

	return in.installBootloader(tempdir, deviceNode)
}

func (in *Installer) installBootloader(rootdir string, rootdevice string) error {
	bootdir := path.Join(rootdir, "boot")
	if err := os.MkdirAll(path.Join(bootdir, "grub"), 0755); err != nil {
		return err
	}
	args := []string{"--root-directory=" + rootdir, "--boot-directory=" + bootdir, "--recheck"}

	if _, err := os.Stat("/sys/firmware/efi/efivars"); os.IsNotExist(err) {
		BootDevice := ""
		for len(BootDevice) == 0 {
			disks := make([]string, 0)
			for _, a := range in.devices.Devices {
				disks = append(disks, a.Path)
			}
			BootDevice, err = in.DialogList("Boot Disk", "disk", disks...)
			if err != nil {
				return err
			}
		}
		args = append(args, BootDevice)
	} else {
		efidir := path.Join(bootdir, "efi")
		if err := os.MkdirAll(efidir, 0755); err != nil {
			return err
		}

		BootDevice := ""
		for len(BootDevice) == 0 {
			log.Println("Search EFI partition with LABEL 'EFI'")
			for _, disk := range in.devices.Devices {
				for _, part := range disk.Childrens {
					if part.Label == "EFI" {
						BootDevice = part.Path
						break
					}
				}
				if len(BootDevice) != 0 {
					break
				}
			}

			if len(BootDevice) == 0 {
				if err := exec.Command("gparted").Run(); err != nil {
					return err
				}

				if err := in.StageSysConfig(); err != nil {
					return err
				}
			}
		}

		log.Println("mounting boot device ", BootDevice)
		if err := exec.Command("/bin/mount", BootDevice, efidir).Run(); err != nil {
			return err
		}
		args = append(args, "--bootloader-id=rlx")
	}

	log.Println("installing grub")
	err := exec.Command("/bin/grub-install", args...).Run()
	if err != nil {
		return err
	}

	var UUID string
	log.Println("configuring bootloader")
	for _, disk := range in.devices.Devices {
		for _, part := range disk.Childrens {
			if part.Path == rootdevice {
				UUID = part.UUID
				break
			}

			if len(UUID) != 0 {
				break
			}
		}
	}

	device_config := "UUID=" + UUID
	if len(UUID) == 0 {
		device_config = rootdevice
	}

	grubcfg := `
insmod part_gpt
insmod part_msdos
insmod all_video
timeout=5
default='rlxos inital setup'
menuentry 'rlxos inital setup' {
	insmod gzio
	insmod ext2
	linux /boot/vmlinuz root=` + device_config + " system=" + VERSION + `
	initrd /boot/initrd
}`

	if err := os.WriteFile(path.Join(rootdir, "boot", "grub", "grub.cfg"), []byte(grubcfg), 0644); err != nil {
		return err
	}

	if err := exec.Command("/bin/cp", "/run/iso/boot/vmlinuz", path.Join(rootdir, "boot", "vmlinuz")).Run(); err != nil {
		return err
	}

	if err := exec.Command("/bin/cp", "/run/iso/boot/initrd", path.Join(rootdir, "boot", "initrd")).Run(); err != nil {
		return err
	}

	return nil
}
