package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	if reader == nil {
		panic(fmt.Errorf("problem in reader"))
	}

	Init()

	for true {
		cmd, _ := reader.ReadString('\n')
		cmd = strings.Trim(cmd, "\n")

		switch {
		case cmd == "new":
			// New process function
		case cmd == "kill":
			// Kill process
		case cmd == "monitor":
			// Monitoring
		case cmd == "terminate":
			break
		}
	}
}
