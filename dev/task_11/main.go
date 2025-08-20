package main

import "fmt"

func intersection(a, b []int) []int {
	if len(a) > len(b) {
		return intersection(b, a)
	} else {
		m := make(map[int]bool)
		for _, v := range a {
			m[v] = true
		}

		inter_result := []int{}
		for _, v := range b {
			if m[v] {
				inter_result = append(inter_result, v)
			}
		}

		return inter_result
	}
}

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	fmt.Printf("Intersection = %v\n", intersection(a, b))
}
