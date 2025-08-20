package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func worker(ctx context.Context, id int, ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case data, ok := <-ch:
			if !ok {
				return
			}
			fmt.Printf("Worker %d received: %s\n", id, data)
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	fmt.Print("Enter number of workers: ")
	var n int
	if _, err := fmt.Scan(&n); err != nil || n <= 0 {
		fmt.Println("Invalid input")
		return
	}

	ctx, done := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer done()

	ch := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(ctx, i, ch, &wg)
	}

	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		defer close(ch)

		for scanner.Scan() {
			select {
			case ch <- scanner.Text():

			case <-ctx.Done():
				return
			}
		}

	}()

	fmt.Println("Enter data:")

	wg.Wait()
	fmt.Println("All workers stopped")
}
