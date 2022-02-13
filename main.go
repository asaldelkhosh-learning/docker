package main

import "strings"

func main() {
	inp := Input{}.Init()
	Init()

	for true {
		cmd := strings.Split(inp.Get(), " ")

		switch {
		case cmd[0] == "new":
			// New process function
		case cmd[0] == "kill":
			// Kill process
		case cmd[0] == "monitor":
			// Monitoring
		case cmd[0] == "terminate":
			break
		}
	}
}
