package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Сигнал ОС
func main() {
	var wg sync.WaitGroup

	// Канал для получения ОС‑сигналов
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM) // Ctrl+C, kill

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-sigCh: // получили сигнал от ОС
				fmt.Println("Received OS signal – stop goroutine")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	wg.Wait() // ждем завершения горутины
}
