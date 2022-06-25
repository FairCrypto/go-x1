package topicsdb

import (
	"math"
	"sync"
)

type (
	posCounter struct {
		pos   int
		count int
	}

	blockCounter struct {
		wait  chan struct{}
		count int
	}

	synchronizator struct {
		mu        sync.Mutex
		threads   sync.WaitGroup
		positions []*posCounter
		goNext    bool
		minBlock  uint64
		blocks    map[uint64]*blockCounter
	}
)

func newSynchronizator() *synchronizator {
	s := &synchronizator{
		positions: make([]*posCounter, 0),
		goNext:    true,
		minBlock:  0,
		blocks:    make(map[uint64]*blockCounter),
	}

	return s
}

func (s *synchronizator) Halt() {
	s.goNext = false

	s.mu.Lock()
	defer s.mu.Unlock()

	for n := range s.blocks {
		if n != s.minBlock {
			close(s.blocks[n].wait)
		}
	}
}

func (s *synchronizator) GoNext(n uint64) bool {
	if !s.goNext {
		return false
	}

	if n > s.minBlock {
		s.mu.Lock()
		s.enqueueBlock(n)
		s.dequeueBlock()
		s.mu.Unlock()
		// wait for other threads
		<-s.blocks[n].wait
	}

	return s.goNext
}

func (s *synchronizator) StartThread(pos int, num int) {
	s.threads.Add(1)

	s.mu.Lock()
	defer s.mu.Unlock()

	s.enqueueBlock(s.minBlock)

	if len(s.positions) == 0 || s.positions[len(s.positions)-1].pos != pos {
		s.positions = append(s.positions, &posCounter{pos, 1})
	} else {
		s.positions[len(s.positions)-1].count++
	}
}

func (s *synchronizator) FinishThread(pos int, num int) {
	s.threads.Done()

	s.mu.Lock()
	defer s.mu.Unlock()

	s.dequeueBlock()
}

func (s *synchronizator) enqueueBlock(n uint64) {
	if _, ok := s.blocks[n]; ok {
		s.blocks[n].count++
	} else {
		s.blocks[n] = &blockCounter{
			wait:  make(chan struct{}),
			count: 1,
		}
	}
}

func (s *synchronizator) dequeueBlock() {
	s.blocks[s.minBlock].count--
	if s.blocks[s.minBlock].count < 1 {
		delete(s.blocks, s.minBlock)
		if len(s.blocks) < 1 {
			return
		}
		// find new minBlock
		s.minBlock = math.MaxUint64
		for b := range s.blocks {
			if s.minBlock > b {
				s.minBlock = b
			}
		}
		close(s.blocks[s.minBlock].wait)
	}
}

func (s *synchronizator) WaitForThreads() {
	s.threads.Wait()
}

func (s *synchronizator) CriteriesCount() int {
	return len(s.positions)
}
