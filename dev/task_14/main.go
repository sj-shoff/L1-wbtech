package main

import (
	"fmt"
	"reflect"
)

func determineType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Printf("Type: int, Value: %d\n", v)
	case string:
		fmt.Printf("Type: string, Value: %s\n", v)
	case bool:
		fmt.Printf("Type: bool, Value: %t\n", v)
	default:
		if reflect.TypeOf(v).Kind() == reflect.Chan {
			fmt.Printf("Type: chan, Value: %v\n", v)
		} else {
			fmt.Printf("Unknown type: %s\n", reflect.TypeOf(v))
		}
	}
}

func main() {
	i := 10
	s := "Hello"
	b := true
	c := make(chan int)
	f := 10.5
	determineType(i)
	determineType(s)
	determineType(b)
	determineType(c)
	determineType(f)
}
