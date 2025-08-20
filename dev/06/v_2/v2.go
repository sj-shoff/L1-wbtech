package main

import (
	"fmt"
	"sync"
	"time"
)

// Таймаут через select
func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		timeout := time.After(2 * time.Second)

		for {
			select {
			case <-timeout:
				fmt.Println("Stopped by timeout")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
}
