package installer

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"

	"github.com/itsmanjeet/installer/internal/progress"
)

func (i *Installer) Install(p progress.ProgressBar) error {
	syscall.Unmount(i.Root, 0)

	if data, err := exec.Command("mount", i.Root, i.Workdir).CombinedOutput(); err != nil {
		p.Update(100, fmt.Sprintf("failed to mount root device %s, %s", data, err))
		return err
	}
	defer syscall.Unmount(i.Workdir, 0)

	p.Update(0, "creating file hierarchy")
	for _, dir := range REQUIRED_DIR {
		if err := os.MkdirAll(path.Join(i.Workdir, dir), 0755); err != nil {
			p.Update(100, "failed to create required directory "+dir)
			return err
		}
	}

	sysroot := path.Join(i.Workdir, "sysroot")
	bootdir := path.Join(i.Workdir, "boot")

	p.Update(20, "installing system")
	if output, err := exec.Command("unsquashfs", "-f", "-d", sysroot, i.SystemImage).CombinedOutput(); err != nil {
		p.Update(100, fmt.Sprintf("failed to extract system image %s, %s", output, err))
		return err
	}

	if err := os.Chmod(sysroot, 0755); err != nil {
		p.Update(100, "failed to fixed root permission")
		return err
	}

	p.Update(50, "installing kernel")
	if output, err := exec.Command("cp", "-r", "/boot/vmlinuz-"+i.Kernel, "/boot/modules/"+i.Kernel, bootdir).CombinedOutput(); err != nil {
		p.Update(100, fmt.Sprintf("failed to install kernel %s, %s", output, err))
		return err
	}

	p.Update(60, "installing bootloader")
	bootCommand := []string{"--root-directory=" + sysroot, "--boot-directory=" + bootdir, "--recheck"}
	if i.IsEfi {
		efiPath := path.Join(bootdir, "efi")
		if err := os.MkdirAll(efiPath, 0755); err != nil {
			p.Update(100, "failed to create efi path")
			return err
		}
		if err := syscall.Mount(i.Boot, efiPath, "fat32", 0, ""); err != nil {
			p.Update(100, "failed to mount efi partition")
			return err
		}
		defer syscall.Unmount(efiPath, 0)
	} else {
		bootCommand = append(bootCommand, i.Boot)
	}

	if data, err := exec.Command("grub-install", bootCommand...).CombinedOutput(); err != nil {
		p.Update(100, fmt.Sprintf("failed to install bootloader %s, %s", string(data), err))
		return err
	}

	p.Update(80, "configuraing bootloader")
	if err := ioutil.WriteFile(path.Join(bootdir, "grub", "grub.cfg"), []byte(fmt.Sprintf(GRUB_CONFIG, i.Kernel, i.rootUUID, i.Kernel)), 0644); err != nil {
		p.Update(100, "failed to configure bootloader")
		return err
	}

	userId := 1000
	for count, user := range i.Users {
		userId += count
		passwd, err := exec.Command("openssl", "passwd", "-1", user.Password).CombinedOutput()
		if err != nil {
			p.Update(100, fmt.Sprintf("failed to calculate password hash %s, %s", passwd, err))
			return err
		}
		if data, err := exec.Command("useradd", "-R", i.Workdir, "-p", string(passwd), "-G", "adm", "-g", "user", user.Name).CombinedOutput(); err != nil {
			p.Update(100, fmt.Sprintf("failed to create new user %s, %s", data, err))
			return err
		}

		homedir := path.Join(i.Workdir, "home", user.Name)
		if err := os.MkdirAll(homedir, 0750); err != nil {
			p.Update(100, "failed to create user directory")
			return err
		}

		if err := os.Chown(homedir, userId, 999); err != nil {
			p.Update(100, "failed to fix user directory permission")
			return err
		}
	}

	p.Update(90, "generating and configuring locale")
	split := strings.Split(i.Locale, ".")
	localepath := path.Join(sysroot, "usr", "lib", "locale")
	if checkExists(localepath) != nil {
		os.MkdirAll(localepath, 0755)
	}
	if output, err := exec.Command("localedef", "-i", split[0], "-f", split[1], i.Locale, "--prefix="+localepath).CombinedOutput(); err != nil {
		p.Update(90, fmt.Sprintf("failed to generate locale %s, %s", output, err))
	}

	p.Update(95, "configurating timezone")
	if err := os.Link(path.Join("/", "usr", "share", "zoneinfo", i.TimeZone), path.Join(sysroot, "etc", "localtime")); err != nil {
		p.Update(95, fmt.Sprintf("failed to configure time zone %s", err))
	}

	p.Update(98, "generating machine id")
	uuid, err := exec.Command("uuidgen").CombinedOutput()
	if err != nil {
		p.Update(98, fmt.Sprintf("failed to generate machine id %s, %s", uuid, err))
	} else {
		ioutil.WriteFile(path.Join(sysroot, "etc", "machine-id"), uuid, 0644)
	}

	p.Update(100, "installation success")
	return nil
}
