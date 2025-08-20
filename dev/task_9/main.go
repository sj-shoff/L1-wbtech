package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}

	ch_in := make(chan int)
	ch_out := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for n := range ch_in {
			ch_out <- n * 2
		}
		close(ch_out)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for n := range ch_out {
			fmt.Printf("%d ", n)
		}
		fmt.Println()
	}()

	test_data := []int{1, 2, 3, 4, 5}

	for _, n := range test_data {
		ch_in <- n
	}

	close(ch_in)
	wg.Wait()
}
