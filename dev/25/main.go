package main

import (
	"fmt"
	"time"
)

func Sleep(duration time.Duration) {
	ticker := time.NewTicker(duration)

	defer ticker.Stop()

	<-ticker.C
}

func main() {
	fmt.Println("Start", time.Now())
	Sleep(4 * time.Second)
	fmt.Println("End", time.Now())
}
