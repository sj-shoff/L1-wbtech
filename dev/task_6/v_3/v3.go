package main

import (
	"fmt"
	"sync"
	"time"
)

// Простой канал‑сообщение
func main() {
	wg := sync.WaitGroup{}
	quit := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-quit: // остановка не из-за закрытия канала, а из-за сообщения
				fmt.Println("Stopped by channel")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	close(quit)
	wg.Wait()
}
