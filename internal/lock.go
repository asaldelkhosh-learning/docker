package internal

import "sync"

var (
	C    *sync.Cond
	Last int32
)

func initLock() {
	C = sync.NewCond(&sync.Mutex{})
	Last = -1
}
