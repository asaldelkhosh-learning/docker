package lock

import "sync"

var (
	C    *sync.Cond
	Last int32
)

func Init() {
	C = sync.NewCond(&sync.Mutex{})
	Last = -1
}
