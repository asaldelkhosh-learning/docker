package storage

import (
	"fmt"
	"github.com/amirhnajafiz/procces-monitoring/lock"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"os/exec"
	"time"
)

func (s Storage) View() {
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

	fmt.Printf("Last PID %d\nLast Update %s\n", lock.Last, time.Now().Format(time.StampMilli))
}
