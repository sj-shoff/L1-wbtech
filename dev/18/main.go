package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	wg := sync.WaitGroup{}

	cnt := Counter{}
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cnt.Increment()
		}()
	}

	wg.Wait()
	fmt.Println("Counter value:", cnt.Value())
}
