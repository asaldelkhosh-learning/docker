package internal

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func Execute() {
	var (
		c  int
		er error
	)

	if len(os.Args) == 1 {
		c = 10
	} else {
		c, er = strconv.Atoi(os.Args[1])
	}

	if er != nil {
		panic(fmt.Errorf("limit should be number, invalid: '%s'", os.Args[1]))
	}

	stg := Storage{}
	stg.Init(c)

	inp := Input{}.Init()

	Init()

	user, _ := os.Hostname()

	for true {
		stg.View()
		fmt.Printf("\n%s > ", user)
		cmd, err := inp.Decode(inp.Get())

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

			proc := stg.Add(&Process{
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

			go proc.Run()
		case cmd["command"] == "kill":
			ID, err := strconv.Atoi(cmd["--id"])
			if err != nil {
				panic(err)
			}

			stg.Kill(int32(ID))
		case cmd["command"] == "pause":
			ID, err := strconv.Atoi(cmd["--id"])
			if err != nil {
				panic(err)
			}

			stg.Pause(int32(ID), true)
		case cmd["command"] == "run":
			ID, err := strconv.Atoi(cmd["--id"])
			if err != nil {
				panic(err)
			}

			stg.Pause(int32(ID), false)
		case cmd["command"] == "terminate":
			os.Exit(1)
		}
	}
}
