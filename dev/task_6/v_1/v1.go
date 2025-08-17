package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Атомарный флаг
func main() {
	var (
		wg   sync.WaitGroup
		stop atomic.Bool // безопасный bool
	)

	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			if stop.Load() { // безопасное чтение
				fmt.Println("Stopped by atomic flag")
				return
			}
			fmt.Println("Working…")
			time.Sleep(500 * time.Millisecond)
		}
	}()

	time.Sleep(2 * time.Second)
	stop.Store(true) // безопасная запись
	wg.Wait()

	// Если использовать обычный stopFlag(bool переменная),
	// то это будет выглядеть правильно, однако на практике
	// это небезопасно – переменная stopFlag читается и записывается
	// из разных горутин без синхронизации
}
