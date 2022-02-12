package main

type Storage struct {
	list     []Process
	capacity int
}

func (s *Storage) Init(capacity int) {
	s.capacity = capacity
}
