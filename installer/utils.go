package installer

import (
	"os"
	"os/exec"
)

func checkTool(id string) error {
	_, err := exec.LookPath(id)
	return err
}

func checkExists(path string) error {
	_, err := os.Stat(path)
	return err
}
