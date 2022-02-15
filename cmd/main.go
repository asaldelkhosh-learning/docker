package main

import (
	"cmd/input"
	"cmd/lock"
	"cmd/process"
	"cmd/storage"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	stg := storage.Storage{}
	stg.Init(2)

	inp := input.Input{}.Init()

	lock.Init()

	for true {
		fmt.Print("> ")
		cmd, err := inp.Decode(inp.Get())

		if err != nil {
			fmt.Printf(err.Error())
			continue
		}

		switch {
		case cmd["command"] == "new":
			delay, err := strconv.Atoi(cmd["--delay"])
			if err != nil {
				panic(err)
			}

			proc := stg.Add(&process.Process{
				Delay:     int32(delay),
				Task:      cmd["--task"],
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Terminate: false,
			})

			if proc == nil {
				fmt.Println("not enough capacity to create new process")
				continue
			}

			go proc.Run()
		case cmd["command"] == "kill":
			ID, err := strconv.Atoi(cmd["--id"])
			if err != nil {
				panic(err)
			}

			stg.Kill(int32(ID))
		case cmd["command"] == "monitor":
			stg.View()
		case cmd["command"] == "terminate":
			os.Exit(1)
		}
	}
}
