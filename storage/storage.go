package storage

import (
	"cmd/lock"
	"cmd/process"
	"fmt"
)

type Storage struct {
	list     []*process.Process
	capacity int
}

var pid int

func (s *Storage) Init(capacity int) {
	s.capacity = capacity
	pid = 1
}

func (s *Storage) Add(p *process.Process) *process.Process {
	if pid > s.capacity {
		return nil
	}

	p.PID = int32(pid)
	pid++

	s.list = append(s.list, p)

	return p
}

func (s *Storage) Kill(ID int32) {
	for index := 0; index < len(s.list); index++ {
		if proc := s.list[index]; ID == proc.PID {
			proc.Terminate = true
			s.list = append(s.list[:index], s.list[index+1:]...)
			break
		}
	}
}

func (s Storage) View() {
	fmt.Printf("Last Process: %d\n", lock.Last)
	for i, p := range s.list {
		fmt.Print(p.Status(i))
	}
}
