package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func worker(id int, ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range ch {
		fmt.Printf("Worker %d received: %s\n", id, data)
	}
	fmt.Printf("Worker %d stopped\n", id)
}

func main() {
	fmt.Print("Enter number of workers: ")
	var n int
	if _, err := fmt.Scan(&n); err != nil || n <= 0 {
		fmt.Println("Invalid input")
		return
	}

	ch := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChan
		fmt.Println("\nShutting down...")
		close(ch)
	}()

	fmt.Println("Enter data:")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		data := scanner.Text()
		ch <- data
	}

	close(ch)
	wg.Wait()
	fmt.Println("All workers stopped")
}
