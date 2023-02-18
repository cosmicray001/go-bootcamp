package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

func main() {
	circle := new(Circle)
	circle.r = 123.321
	cirArea := circleArea(circle)
	fmt.Printf("%.4f\n", cirArea)
}

func circleArea(c *Circle) float64 {
	return math.Pi * (c.r * c.r)
}