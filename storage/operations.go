package storage

import "cmd/process"

func (s *Storage) Add(p *process.Process) *process.Process {
	if pid > s.capacity {
		return nil
	}

	p.PID = int32(pid)
	pid++

	s.list = append(s.list, p)

	return p
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

func (s Storage) Pause(ID int32, flag bool) {
	for _, p := range s.list {
		if p.PID == ID {
			p.Pause = flag
			return
		}
	}
}
