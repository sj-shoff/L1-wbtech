package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverseRuneSlice(r []rune, left, right int) {
	for i, j := left, right-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
}

func reverseWords(s string) string {
	r := []rune(s)

	reverseRuneSlice(r, 0, len(r))

	start := 0
	for i := 0; i <= len(r); i++ {
		if i == len(r) || r[i] == ' ' {
			reverseRuneSlice(r, start, i)
			start = i + 1
		}
	}
	return string(r)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		line := scanner.Text()
		fmt.Println(reverseWords(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading error:", err)
	}
}
