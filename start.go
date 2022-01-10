package main

import (
	"PM/pool"
	"fmt"
)

func main() {
	pl := pool.GetPool(5)

	defer pl.Close()
	pl.Start()

	for _, prc := range pl.Monitor().WorkerList() {
		fmt.Println(prc)
	}
}
