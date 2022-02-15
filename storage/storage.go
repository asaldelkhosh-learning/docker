package storage

import (
	"cmd/process"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Storage struct {
	list     []*process.Process
	capacity int
}

var (
	pid    int
	format = "dd MMMM HH:mm:ss"
)

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
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleColoredBlackOnGreenWhite)
	t.SetTitle("Monitoring Processes")
	t.AppendHeader(table.Row{"#", "PID", "Delay", "Created At", "Last Update", "Task", "Number of executions"})

	for i, p := range s.list {
		t.AppendRow([]interface{}{i, p.PID, p.Delay, p.CreatedAt.Format(format), p.UpdatedAt.Format(format), p.Task, p.Called})
	}
	t.Render()
}
