package internal

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func Run() {
	var (
		c  int
		er error
	)

	if len(os.Args) == 1 {
		c = 10
	} else {
		if c, er = strconv.Atoi(os.Args[1]); er != nil {
			panic(fmt.Errorf("limit should be number, invalid: '%s'", os.Args[1]))
		}
	}

	stg := newStorage(c)
	inp := newInput()

	newLock()

	user, _ := os.Hostname()

	for true {
		stg.view()

		fmt.Printf("\n%s > ", user)

		cmd, err := inp.decode(inp.get())

		if err != nil {
			fmt.Println(err.Error())

			continue
		}

		switch {
		case cmd["command"] == "new":
			delay, err := strconv.Atoi(cmd["--delay"])
			if err != nil {
				panic(err)
			}

			burst, err := strconv.Atoi(cmd["--burst"])
			if err != nil {
				panic(err)
			}

			proc := stg.add(&process{
				Delay:     int32(delay),
				Task:      cmd["--task"],
				Burst:     int32(burst),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Terminate: false,
				Pause:     false,
			})

			if proc == nil {
				fmt.Println("not enough capacity to create new process")
				continue
			}

			go proc.run()
		case cmd["command"] == "kill":
			ID, err := strconv.Atoi(cmd["--id"])
			if err != nil {
				panic(err)
			}

			stg.kill(int32(ID))
		case cmd["command"] == "pause":
			ID, err := strconv.Atoi(cmd["--id"])
			if err != nil {
				panic(err)
			}

			stg.pause(int32(ID), true)
		case cmd["command"] == "run":
			ID, err := strconv.Atoi(cmd["--id"])
			if err != nil {
				panic(err)
			}

			stg.pause(int32(ID), false)
		case cmd["command"] == "terminate":
			os.Exit(1)
		}
	}
}
