package main

import "fmt"

func binarySearch(slice []int, target int) int {
	low := 0
	high := len(slice) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := slice[mid]
		if guess == target {
			return mid
		} else if guess > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func main() {
	test_slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	search := binarySearch(test_slice, 8)
	fmt.Println(search)

	search = binarySearch(test_slice, 12)
	fmt.Println(search)
}
