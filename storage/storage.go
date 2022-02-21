package storage

import (
	"cmd/process"
)

type Storage struct {
	list     []*process.Process
	capacity int
}

var (
	pid int
)

func (s *Storage) Init(capacity int) {
	s.capacity = capacity
	pid = 1
}
