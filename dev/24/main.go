package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) Distance(other *Point) float64 {
	if p == nil || other == nil {
		return 0
	} else {
		return math.Sqrt(math.Pow(p.x-other.x, 2) + math.Pow(p.y-other.y, 2))
	}
}

func main() {
	for {
		var x1, y1, x2, y2 float64
		fmt.Print("Enter coordinates of the first point: ")
		fmt.Scan(&x1, &y1)
		fmt.Print("Enter coordinates of the second point: ")
		fmt.Scan(&x2, &y2)
		if x1 == 0 && y1 == 0 && x2 == 0 && y2 == 0 {
			break
		} else {
			p1 := NewPoint(x1, y1)
			p2 := NewPoint(x2, y2)
			fmt.Printf("Distance: %.2f\n", p1.Distance(p2))
		}
	}
}
