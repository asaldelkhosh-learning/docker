package tools

import (
	"os"
	"os/exec"
)

func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	_ = c.Run()
}
