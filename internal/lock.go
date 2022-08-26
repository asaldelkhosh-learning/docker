package internal

import "sync"

var (
	C    *sync.Cond
	Last int32
)

func newLock() {
	C = sync.NewCond(&sync.Mutex{})
	Last = -1
}
