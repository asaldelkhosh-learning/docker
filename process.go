package main

import "time"

type Process struct {
	PID       int32
	Task      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Terminate bool
}
