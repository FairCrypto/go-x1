package utils

import (
	"container/list"
	"sync"
)

type Stack struct {
	dll   *list.List
	mutex sync.Mutex
}

func NewStack() *Stack {
	return &Stack{dll: list.New()}
}

func (s *Stack) Push(x interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.dll.PushBack(x)
}

func (s *Stack) Pop() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.dll.Len() == 0 {
		return nil
	}
	tail := s.dll.Back()
	val := tail.Value
	s.dll.Remove(tail)
	return val
}
