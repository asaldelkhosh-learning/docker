package main

import (
	"strconv"
	"strings"
)

func main() {
	stg := Storage{}
	stg.Init(10)

	inp := Input{}.Init()

	Init()

	for true {
		cmd := strings.Split(inp.Get(), " ")

		switch {
		case cmd[0] == "new":
			// New process function
		case cmd[0] == "kill":
			ID, err := strconv.Atoi(cmd[1])
			if err != nil {
				panic(err)
			}

			stg.Kill(int32(ID))
		case cmd[0] == "monitor":
			// Monitoring
		case cmd[0] == "terminate":
			break
		}
	}
}
