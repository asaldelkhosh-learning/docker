package main

import (
	"fmt"
	"time"
)

type Process struct {
	PID       int32
	Task      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Delay     int32
	Terminate bool
}

func (p Process) Run() {
	// Function
	for p.Terminate {
		// Lock
		Lock.Lock()
		Last = p.PID
		// Do
		fmt.Printf("Process %d is running now: %s\n", p.PID, p.Task)
		// Unlock
		Lock.Unlock()
		// Waiting
		time.Sleep(time.Second * time.Duration(p.Delay))
	}
}
