package internal

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type storage struct {
	list []*process

	capacity int32
	pid      int32
}

func newStorage(capacity int) storage {
	return storage{
		capacity: int32(capacity),
		pid:      1,
	}
}

func (s *storage) add(p *process) *process {
	if s.pid > s.capacity {
		return nil
	}

	p.PID = s.pid
	s.list = append(s.list, p)

	s.pid++

	return p
}

func (s *storage) kill(ID int32) {
	for index := 0; index < len(s.list); index++ {
		if proc := s.list[index]; ID == proc.PID {
			proc.Terminate = true

			s.list = append(s.list[:index], s.list[index+1:]...)

			break
		}
	}
}

func (s storage) pause(ID int32, flag bool) {
	for _, p := range s.list {
		if p.PID == ID {
			p.Pause = flag

			return
		}
	}
}

func (s storage) view() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	_ = c.Run()

	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleColoredBlackOnGreenWhite)
	t.SetTitle("Monitoring Processes")
	t.AppendHeader(table.Row{"#", "PID", "Delay", "Status", "Executed", "Created At", "Last Update", "Task"})

	for i, p := range s.list {
		var status string

		if status = "RUNNING"; p.Pause {
			status = "WAITING"
		}

		t.AppendRow([]interface{}{
			i + 1,
			p.PID,
			p.Delay,
			status,
			p.Called,
			p.CreatedAt.Format(time.StampMilli),
			p.UpdatedAt.Format(time.StampMilli),
			p.Task,
		})
	}

	t.Render()

	fmt.Printf("Last PID %d\nLast Update %s\n", Last, time.Now().Format(time.StampMilli))
}
