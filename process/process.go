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
	Burst     int32
	Terminate bool
	Called    int
	Pause     bool
}

func (p *Process) Run() {
	// Function
	p.Called = 0
	for !p.Terminate {
		// Check for pause
		if p.Pause {
			time.Sleep(3 * time.Second)
			continue
		}
		// Lock
		lock.C.L.Lock()
		lock.Last = p.PID
		// Do
		p.Called++
		p.UpdatedAt = time.Now()
		// Burst
		time.Sleep(time.Second * time.Duration(p.Burst))
		// Unlock
		lock.C.L.Unlock()
		// Waiting
		time.Sleep(time.Second * time.Duration(p.Delay))
	}
}

func (p *Process) Status(i int) string {
	return fmt.Sprintf("%d: Process %d | Task %s | Executed %d | Last Update %s\n", i+1, p.PID, p.Task, p.Called, p.UpdatedAt)
}
