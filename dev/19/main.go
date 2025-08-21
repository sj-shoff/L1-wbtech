package main

import "fmt"

func reverse_string(s string) string {

	r := []rune(s)
	n := len(r)

	for i := 0; i < n/2; i++ {
		r[i], r[n-1-i] = r[n-1-i], r[i]
	}

	return string(r)
}

func main() {

	fmt.Println("Enter a string:")
	var s string
	fmt.Scan(&s)
	fmt.Println(reverse_string(s))
}
