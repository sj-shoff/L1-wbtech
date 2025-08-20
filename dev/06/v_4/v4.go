package main

import (
	"fmt"
	"sync"
	"time"
)

// Закрытый канал
func main() {

	wg := sync.WaitGroup{}
	stop := make(chan struct{}) // обычный канал, но мы будем его закрывать

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stop: // закрытый канал – остановка
				fmt.Println("Stopped by closed channel")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	close(stop) // закрытие канала - единственный способ остановки
	wg.Wait()
}
