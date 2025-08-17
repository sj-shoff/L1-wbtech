package main

import (
	"fmt"
	"sync"
	"time"
)

func timer(start *time.Time) int {
	return int(time.Since(*start).Seconds())
}

func main() {

	var timeout int
	fmt.Print("Enter timeout (seconds): ")
	_, err := fmt.Scan(&timeout)
	if err != nil || timeout <= 0 {
		fmt.Println("Invalid input")
		return
	}

	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	start := time.Now()

	go func() {
		defer wg.Done()
		i := 0
		for {
			if timer(&start) >= timeout {
				close(ch)
				return
			}

			ch <- i
			i++
		}
	}()

	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("Received:", v)
		}
		fmt.Println("Reader: channel closed – exiting")
	}()

	wg.Wait()
	fmt.Println("Timeout reached – program exiting")
}
