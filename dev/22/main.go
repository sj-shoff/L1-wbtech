package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)
	operationStr := ""

	for {
		fmt.Printf("Enter first number (a): ")
		var inputA string
		fmt.Scan(&inputA)
		if _, ok := a.SetString(inputA, 10); ok {
			break
		}
		fmt.Println("Invalid input. Please enter a valid integer.")
	}

	for {
		fmt.Printf("Enter second number (b): ")
		var inputB string
		fmt.Scan(&inputB)
		if _, ok := b.SetString(inputB, 10); ok {
			break
		}
		fmt.Println("Invalid input. Please enter a valid integer.")
	}

	result := new(big.Int)
	for {
		fmt.Printf("Select an operation (+, -, *, /): ")
		fmt.Scan(&operationStr)

		zero := big.NewInt(0)

		switch operationStr {
		case "+":
			result.Add(a, b)
			fmt.Printf("%s + %s = %s\n", a.String(), b.String(), result.String())
			return
		case "-":
			result.Sub(a, b)
			fmt.Printf("%s - %s = %s\n", a.String(), b.String(), result.String())
			return
		case "*":
			result.Mul(a, b)
			fmt.Printf("%s * %s = %s\n", a.String(), b.String(), result.String())
			return
		case "/":
			if b.Cmp(zero) == 0 {
				fmt.Println("Error: Division by zero. Please choose another operation.")
			} else {
				result.Div(a, b)
				fmt.Printf("%s / %s = %s\n", a.String(), b.String(), result.String())
				return
			}
		default:
			fmt.Println("Invalid operation. Please choose from +, -, *, /.")
		}
	}
}
