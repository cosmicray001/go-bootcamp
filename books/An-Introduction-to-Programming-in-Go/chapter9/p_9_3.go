package main

import (
	"fmt"
	"math"
)

type Shape interface {
	perimeter() float64
}

type MultipleShape struct {
	shapes []Shape
}

type Circle struct {
	R float64
}

type Rectangle struct {
	W, H float64
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.R
}

func (r *Rectangle) perimeter() float64 {
	return 2 * (r.H + r.W)
}

func (ms *MultipleShape) totalArea() float64 {
	sum := 0.0
	for _, s := range ms.shapes {
		sum += s.perimeter()
	}
	return sum
}

func main() {
	c := Circle{R: 5}
	r := Rectangle{H: 3, W: 2}
	s := new(MultipleShape)
	s.shapes = append(s.shapes, &c)
	s.shapes = append(s.shapes, &r)
	fmt.Println(s.totalArea())
}