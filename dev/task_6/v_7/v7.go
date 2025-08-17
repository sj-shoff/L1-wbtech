package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// runtime.Goexit()
func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("Working:", i)
			time.Sleep(200 * time.Millisecond)
		}
		// Ничего не делать – мы просто попадаем в Goexit
		fmt.Println("Calling Goexit() – goroutine ends now")
		runtime.Goexit() // завершает лишь эту горутину
	}()

	wg.Wait()
}
