package misc

import (
	"os/exec"
)

func ShellCommand(cmd string) (string, error) {
	c := exec.Command("bash", "-c", cmd)
	// 此处是windows版本
	// c := exec.Command("cmd", "/C", cmd)
	output, err := c.CombinedOutput()
	return string(output), err
}
