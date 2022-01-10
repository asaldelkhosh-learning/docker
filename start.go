package main

import (
	"PM/pool"
	"PM/tools"
	"fmt"
)

func main() {
	pl := pool.GetPool(5)

	defer pl.Close()
	_ = pl.Start()

	for {
		tools.Clear()
		for _, prc := range pl.Monitor().WorkerList() {
			fmt.Println(prc)
		}
		tools.Delay(3)
	}
}
