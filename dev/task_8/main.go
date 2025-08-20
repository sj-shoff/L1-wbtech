package main

import "fmt"

func setBit(n int64, i uint, bit int) int64 {
	if i > 63 {
		return n
	}
	mask := int64(1) << i
	switch bit {
	case 1:
		return n | mask
	case 0:
		return n &^ mask
	default:
		return n
	}
}

func main() {
	num := int64(5)

	num = setBit(num, 0, 0)
	fmt.Println(num)

	num = setBit(num, 1, 1)
	fmt.Println(num)
}
