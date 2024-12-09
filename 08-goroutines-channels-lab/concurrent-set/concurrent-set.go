package concurrentset

import "sync"

type ConcurrentHashSet struct {
	mu     sync.Mutex
	values map[string]struct{}
}

func New() *ConcurrentHashSet {
	return &ConcurrentHashSet{values: make(map[string]struct{})}
}

func (s *ConcurrentHashSet) Add(val string) {
	s.mu.Lock()
	s.values[val] = struct{}{}
	s.mu.Unlock()
}

func (s *ConcurrentHashSet) IsMember(val string) bool {
	s.mu.Lock()
	_, ok := s.values[val]
	s.mu.Unlock()
	return ok
}

func (s *ConcurrentHashSet) Remove(val string){
	s.mu.Lock()
	delete(s.values, val)
	s.mu.Unlock()
}
