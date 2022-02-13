package main

import (
	"time"
)

type Process struct {
	PID       int32
	Task      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Delay     int32
	Terminate bool
	Called    int
}

func (p Process) Run() {
	// Function
	p.Called = 0
	for p.Terminate {
		// Lock
		Lock.Lock()
		Last = p.PID
		// Do
		p.Called++
		p.UpdatedAt = time.Now()
		// Unlock
		Lock.Unlock()
		// Waiting
		time.Sleep(time.Second * time.Duration(p.Delay))
	}
}
