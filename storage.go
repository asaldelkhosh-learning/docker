package main

type Storage struct {
	list     []Process
	capacity int
}

func (s *Storage) Init(capacity int) {
	s.capacity = capacity
}

func (s *Storage) Add(p Process) {
	s.list = append(s.list, p)
	go p.Run()
}
