package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type SafeMap struct {
	mu sync.RWMutex
	m  map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{m: make(map[string]int)}
}

func (s *SafeMap) Set(key string, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func (s *SafeMap) Get(key string) (int, bool) {
	s.mu.RLock()
	value, ok := s.m[key]
	defer s.mu.RUnlock()
	return value, ok
}

func (s *SafeMap) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.m)
}

func main() {
	const numGoroutines = 50
	const opsGoroutine = 200

	m := NewSafeMap()
	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < opsGoroutine; j++ {
				key := fmt.Sprintf("k-%d-%d", id, j)
				value := rand.Intn(1000) + id*1000
				m.Set(key, value)
			}
		}(i)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < numGoroutines*opsGoroutine; i++ {
			_, _ = m.Get(fmt.Sprintf("k-%d-%d", rand.Intn(numGoroutines), rand.Intn(opsGoroutine)))
		}
	}()

	wg.Wait()

	fmt.Println("Map size:", m.Size())
}
