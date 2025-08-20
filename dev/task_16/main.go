package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)/2]

	less := make([]int, 0, len(arr))
	equal := make([]int, 0, len(arr))
	greater := make([]int, 0, len(arr))

	for _, num := range arr {
		switch {
		case num < pivot:
			less = append(less, num)
		case num == pivot:
			equal = append(equal, num)
		default:
			greater = append(greater, num)
		}
	}

	less = quickSort(less)
	greater = quickSort(greater)

	result := append(less, equal...)
	result = append(result, greater...)

	return result
}

func main() {
	arr := []int{3, 5, 2, 1, 4, 2, 3, 5, 1, 4}
	fmt.Println("Original:", arr)
	fmt.Println("Sorted:", quickSort(arr))
}
