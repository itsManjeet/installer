package installer

import (
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/gotk3/gotk3/glib"
)

func (ins *Installer) StartInstallation() {

	updateProgress := func(mesg string, prog float64) {
		glib.IdleAdd(func() {
			ins.installProgress.SetText(mesg)
			ins.installProgress.SetFraction(prog)
		})

		time.Sleep(time.Millisecond * 150)
	}
	updateProgress("Preparing parition", 0.05)

	workDir, err := os.MkdirTemp("", APPID+"-*")
	if err != nil {
		updateProgress("Failed to prepare work directory", 1.0)
		return
	}

	if !ins.IsDebug(APPID) {
		if err := exec.Command("mount", PartitionPath, workDir).Run(); err != nil {
			updateProgress("Failed to mount partition, "+err.Error(), 1.0)
			return
		}
	}
	updateProgress("Creating rlxos required directories", 0.1)
	systemPath := path.Join(workDir, "rlxos", "system")
	bootPath := path.Join(workDir, "boot")
	efiDir := path.Join(bootPath, "efi")
	if !ins.IsDebug(APPID) {
		if err := os.MkdirAll(systemPath, 0755); err != nil {
			updateProgress("Failed to create required dir, "+err.Error(), 1.0)
			return
		}

		if err := os.MkdirAll(bootPath, 0755); err != nil {
			updateProgress("Failed to create required dir, "+err.Error(), 1.0)
			return
		}

		if BootLoader == BootLoaderTypeEfi {

			if err := os.MkdirAll(efiDir, 0755); err != nil {
				updateProgress("Failed to create efi directory", 1.0)
				return
			}
		}

	}
	updateProgress("Installing System Image", 0.2)
	if !ins.IsDebug(APPID) {
		if err := exec.Command("cp", IMAGE_PATH, path.Join(systemPath, VERSION)).Run(); err != nil {
			updateProgress("Failed to install system image, "+err.Error(), 1.0)
			return
		}
	}
	updateProgress("Installing bootloader", 0.5)
	switch BootLoader {
	case BootLoaderTypeEfi:
		if !ins.IsDebug(APPID) {
			if err := exec.Command("mount", BootDevice, efiDir).Run(); err != nil {
				updateProgress("Failed to mount efi, "+err.Error(), 1.0)
				return
			}
			if err := exec.Command("grub-install", "--root-directory="+workDir, "--boot-directory="+bootPath, "--recheck").Run(); err != nil {
				updateProgress("Failed to install bootloader, "+err.Error(), 1.0)
				return
			}
		}

	case BootLoaderTypeLegacy:
		if !ins.IsDebug(APPID) {
			if err := exec.Command("grub-install", "--root-directory="+workDir, "--boot-directory="+bootPath, "--recheck", BootDevice).Run(); err != nil {
				updateProgress("Failed to install bootloader, "+err.Error(), 1.0)
				return
			}
		}
	default:
		updateProgress("Unsupported Bootloader: "+BootLoader.String(), 1.0)
		return
	}

	kernelVersionExec, err := exec.Command("uname", "-r").CombinedOutput()
	kernelVersion := ""
	if err == nil {
		kernelVersion = "-" + strings.TrimSpace(string(kernelVersionExec))
	}

	updateProgress("Configuring Bootloader", 0.7)
	grubConfig := `
insmod part_gpt
insmod part_msdos
insmod all_video
timeout=5
default='rlxos initial setup'

menuentry 'rlxos initial setup' {
	insmod gzio
	insmod ext2
	linux /boot/vmlinuz` + kernelVersion + ` root=UUID=` + PartitionUUID + " system=" + VERSION + `
	initrd /boot/initrd` + kernelVersion + `	
}`

	if !ins.IsDebug(APPID) {
		if err := os.WriteFile(path.Join(bootPath, "grub", "grub.cfg"), []byte(grubConfig), 0644); err != nil {
			updateProgress("Failed to write grub configuration, "+err.Error(), 1.0)
			return
		}
	}

	updateProgress("Installing kernel", 0.8)
	isoBootPath := path.Join(path.Dir(IMAGE_PATH), "boot")
	if !ins.IsDebug(APPID) {
		if err := exec.Command("cp", path.Join(isoBootPath, "vmlinuz"+kernelVersion), path.Join(bootPath, "vmlinuz"+kernelVersion)).Run(); err != nil {
			updateProgress("Failed to install linux kernel, "+err.Error(), 1.0)
			return
		}
	}

	updateProgress("Installing modules", 0.85)
	if !ins.IsDebug(APPID) {
		if err := exec.Command("cp", path.Join(isoBootPath, "modules"), path.Join(bootPath), "-a").Run(); err != nil {
			updateProgress("Failed to install kernel modules,"+err.Error(), 1.0)
			return
		}
	}

	updateProgress("Installing initrd", 0.9)
	data, err := exec.Command("zenity", "--entry", "--text=Enter Recovery Password", "--hide-text").Output()
	if err != nil {
		updateProgress("Failed to execute zentry, "+err.Error(), 1.0)
		return
	}

	if !ins.IsDebug(APPID) {
		if err := exec.Command("mkinitramfs", "-o="+path.Join(bootPath, "initrd"+kernelVersion), "-p="+string(data)).Run(); err != nil {
			updateProgress("Failed to generate initrd, "+err.Error(), 1.0)
			return
		}
	}
	updateProgress("Installation Success", 1.0)

	glib.IdleAdd(func() {
		ins.Window.SetCurrentPage(4)
	})
}
