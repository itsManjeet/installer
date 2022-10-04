package installer

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func checkTool(id string) error {
	_, err := exec.LookPath(id)
	return err
}

func checkExists(path string) error {
	_, err := os.Stat(path)
	return err
}

func getOutput(bin string, args ...string) (string, error) {
	data, err := exec.Command(bin, args...).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("output: %s, error: %s", data, err)
	}

	return strings.Trim(string(data), "\n"), nil
}
