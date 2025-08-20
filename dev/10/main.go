package main

import (
	"fmt"
	"math"
)

const groupingStep = 10

func groupTemperatures(temperatures []float64) map[int][]float64 {
	groups := make(map[int][]float64)
	for _, temp := range temperatures {
		key := int(math.Floor(temp/groupingStep) * groupingStep)
		groups[key] = append(groups[key], temp)
	}
	return groups
}

func main() {
	test_temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	gT := groupTemperatures(test_temps)
	for group, temperatures := range gT {
		fmt.Printf("%d: %v\n", group, temperatures)
	}
}
