package main

import "fmt"

func main() {
	data := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]bool)

	for _, value := range data {
		set[value] = true
	}

	for key := range set {
		fmt.Printf("%s ", key)
	}
	fmt.Println()
}
