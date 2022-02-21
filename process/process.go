package process

import (
	"time"
)

type Process struct {
	PID       int32
	Task      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Delay     int32
	Burst     int32
	Terminate bool
	Called    int
	Pause     bool
}
