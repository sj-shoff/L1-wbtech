package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// context.WithTimeout / WithDeadline
func main() {

	wg := sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // гарантируем освобождение ресурсов

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done(): // либо cancel(), либо таймаут/дедлайн
				fmt.Println("Stopped by timeout:", ctx.Err())
				return
			default:
				fmt.Println("Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	wg.Wait()
}
