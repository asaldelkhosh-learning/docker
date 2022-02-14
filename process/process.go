package process

import (
	"cmd/lock"
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
	Called    int
}

func (p *Process) Run() {
	// Function
	p.Called = 0
	for !p.Terminate {
		// Lock
		lock.Lock.Lock()
		lock.Last = p.PID
		// Do
		p.Called++
		p.UpdatedAt = time.Now()
		// Unlock
		lock.Lock.Unlock()
		// Waiting
		time.Sleep(time.Second * time.Duration(p.Delay))
	}
}

func (p Process) Status(i int) string {
	return fmt.Sprintf("%d: Process %d | Task %s | Executed %d | Last Update %s\n", i+1, p.PID, p.Task, p.Called, p.UpdatedAt)
}
