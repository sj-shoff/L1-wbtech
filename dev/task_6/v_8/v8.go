package main

import (
	"fmt"
	"sync"
	"time"
)

// panic + recover
func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		// Оборачиваем горутину recover‑оболочкой
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
			}
			wg.Done()
		}()

		for i := 0; i < 10; i++ {
			if i == 5 {
				panic("stop goroutine") // паника для демонстрации
			}
			fmt.Println("Working:", i)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	wg.Wait()
}
