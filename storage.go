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

func (s *Storage) Kill(ID int32) {
	for index := 0; index < len(s.list); index++ {
		if proc := s.list[index]; ID == proc.PID {
			proc.Terminate = true
			s.list = append(s.list[:index], s.list[index+1:]...)
			break
		}
	}
}
