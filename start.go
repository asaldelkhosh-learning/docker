package main

import (
	"PM/pool"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	pl := pool.GetPool(5)

	defer pl.Close()
	pl.Start()

	for {
		for _, prc := range pl.Monitor().WorkerList() {
			fmt.Println(prc)
		}
		time.Sleep(time.Second * 5)
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
	}
}
