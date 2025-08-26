package main

import "fmt"

func removeElement[T any](slice []T, i int) []T {
	if i < 0 || i >= len(slice) {
		return slice
	}

	copy(slice[i:], slice[i+1:])

	// Обнуляем последний элемент для избежания утечки памяти
	var zero T
	slice[len(slice)-1] = zero

	return slice[:len(slice)-1]
}

func main() {
	testSlice := []int{1, 2, 3, 4, 5}
	testSlice2 := []string{"a", "b", "c", "d", "e"}
	testSlice3 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	pslice := removeElement(testSlice, 2)
	pslice2 := removeElement(testSlice2, 2)
	pslice3 := removeElement(testSlice3, 2)
	fmt.Println(pslice, pslice2, pslice3)
}
