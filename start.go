package main

import (
	"PM/pool"
	"PM/tools"
)

func main() {
	pl := pool.GetPool(5)

	defer pl.Close()
	_ = pl.Start()

	for {
		tools.Clear()
		tools.Render(pl)
		tools.Delay(3)
	}
}
