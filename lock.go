package main

import "sync"

var (
	Lock sync.Mutex
	Last int32
)

func Init() {
	Lock = sync.Mutex{}
	Last = -1
}
