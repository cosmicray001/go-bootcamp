package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
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

func totalArea(shapes ...Shape) float64 {
	sum := 0.0
	for _, s := range shapes {
		sum += s.area()
	}
	return sum
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
	area := totalArea(&c, &r)
	fmt.Println(area)

}
