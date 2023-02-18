package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func main() {
	var circle Circle = Circle{1, 2, 3}
	cirArea := circleArea(circle)
	fmt.Printf("%.4f\n", cirArea)
}

func circleArea(cir Circle) float64 {
	return math.Pi * (cir.r * cir.r)
}
