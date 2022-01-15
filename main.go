package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	cmd := exec.Command("./script.sh")
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Just ran subprocess %d, exiting\n", cmd.Process.Pid)
	time.Sleep(5 * time.Second)
}
