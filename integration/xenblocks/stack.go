package xenblocks

import (
	"errors"
	"sync"
)

type Stack struct {
	lock    sync.Mutex // you don't have to do this if you don't want thread safety
	s       []WebSocketJob
	maxSize int
}

func NewStack() *Stack {
	return &Stack{sync.Mutex{}, make([]WebSocketJob, 0), 1}
}

func (s *Stack) Push(v WebSocketJob) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l >= s.maxSize {
		s.s = s.s[:l-1]
	}
	s.s = append(s.s, v)
}

func (s *Stack) Pop() (WebSocketJob, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l == 0 {
		return WebSocketJob{}, errors.New("empty Stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}
