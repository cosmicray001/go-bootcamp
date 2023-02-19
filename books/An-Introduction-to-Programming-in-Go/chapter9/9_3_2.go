package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type MultiShape struct {
	shapes []Shape
}

type Circle struct {
	Type string
	Radious float64
}

type Rectangle struct {
	Type string
	Height float64
	Weight float64
}

func (cir *Circle) area() float64 {
	return 2 * math.Pi * cir.Radious
}

func (rect *Rectangle) area() float64 {
	return rect.Height * rect.Weight
}

func (m *MultiShape) TotalArea() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func main() {
	c := Circle{
		Type: "Cicle",
		Radious: 5,
	}
	r := Rectangle{
		Type: "Rectangle",
		Height: 3,
		Weight: 2,
	}
	s := new(MultiShape)
	s.shapes = append(s.shapes, &c)
	s.shapes = append(s.shapes, &r)
	a := s.TotalArea()
	fmt.Println(a)
}
