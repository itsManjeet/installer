package installer

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/itsmanjeet/installer/internal/progress"
)

func (i *Installer) Verify(p progress.ProgressBar) error {
	p.Update(10, "checking required tools")
	for _, tool := range REQUIRED_TOOLS {
		if err := checkTool(tool); err != nil {
			p.Update(100, "missing required tool "+tool)
			return err
		}
	}

	p.Update(20, "verifying system Image")
	if err := checkExists(i.SystemImage); err != nil {
		p.Update(100, "missing system image "+i.SystemImage)
		return err
	}

	p.Update(30, "verifying root device node")
	if err := checkExists(i.Root); err != nil {
		p.Update(100, "missing root device node "+i.Root)
		return err
	}

	if len(i.Kernel) == 0 {
		p.Update(40, "detecting kernel version")
		if data, err := getOutput("uname", "-r"); err != nil {
			p.Update(100, "failed to get kernel version")
			return err
		} else {
			i.Kernel = data
		}
	}

	p.Update(40, "verifying kernel and modules")
	if err := checkExists(path.Join("/boot", "modules", i.Kernel)); err != nil {
		p.Update(100, "missing required kernel "+i.Kernel)
		return err
	}

	if len(i.Workdir) == 0 {
		var err error
		i.Workdir, err = ioutil.TempDir(os.TempDir(), "installer-*")
		if err != nil {
			p.Update(100, "failed to create temporary directory")
			return err
		}
	} else {
		if err := os.MkdirAll(i.Workdir, 0755); err != nil {
			p.Update(100, "failed to create temporary directory")
			return err
		}
	}

	p.Update(50, "getting root device uuid")
	if data, err := getOutput("lsblk", "-n", "-o", "uuid", i.Root); err != nil {
		p.Update(100, fmt.Sprintf("failed to get device uuid %s, %s", data, err))
		return err
	} else {
		i.rootUUID = string(data)
	}

	p.Update(60, "verifying locale")
	if !strings.Contains(i.Locale, ".") {
		p.Update(100, "invalid locale specified")
		return fmt.Errorf("invalid locale " + i.Locale)
	}

	p.Update(100, "Verification success")
	return nil
}
