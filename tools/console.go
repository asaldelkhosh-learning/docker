package tools

import (
	"fmt"
	"github.com/hamed-yousefi/gowl"
	"time"
)

func Render(pl gowl.Pool) {
	fmt.Println(time.Now())
	mon := pl.Monitor()

	for _, prc := range mon.WorkerList() {
		fmt.Println(pl.Monitor().WorkerStatus(prc))
	}
}
