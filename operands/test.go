package operands

import (
	"log"
	"os"
	"os/exec"
	"time"
)

var DIR = "./scripts/"

func RunTest() {
	cmd := exec.Command(DIR + "script.sh")
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Just ran subprocess %d, exiting\n", cmd.Process.Pid)
	time.Sleep(2 * time.Second)
}
