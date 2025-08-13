package main

import (
	"fmt"
	"sync"
)

type Concurrency struct {
	data []int
}

func NewConcurrency(data []int) *Concurrency {
	c := Concurrency{
		data: make([]int, len(data)),
	}
	copy(c.data, data)
	return &c
}

func (c *Concurrency) calculateSquares() {
	var wg sync.WaitGroup
	for i := range c.data {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			fmt.Printf("start  idx=%d data=%d\n", idx, c.data[idx])
			c.data[idx] = c.data[idx] * c.data[idx]
			fmt.Printf("end    idx=%d data=%d\n", idx, c.data[idx])
		}(i)
	}
	wg.Wait()
}

func main() {
	data := []int{2, 4, 6, 8, 10}
	c := NewConcurrency(data)
	c.calculateSquares()

	for _, value := range c.data {
		fmt.Printf("%d ", value)
	}
	fmt.Println()
}
