package main

import "fmt"

func main() {
	a := 10
	b := 5

	fmt.Printf("Before swap: a = %d, b = %d\n", a, b)

	// Способ 1: Сложение и вычитание
	a = a + b // a = a + b
	b = a - b // b = старое a
	a = a - b // a = старое b

	fmt.Printf("After swap (add/sub): a = %d, b = %d\n", a, b)

	a = 10 // Сброс значений
	b = 5

	fmt.Printf("Before swap: a = %d, b = %d\n", a, b)

	// Способ 2: XOR-обмен
	a = a ^ b // a = a XOR b
	b = a ^ b // b = старое a
	a = a ^ b // a = старое b

	fmt.Printf("After swap (XOR): a = %d, b = %d\n", a, b)
}
