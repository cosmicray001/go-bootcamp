package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perim() float64
}

type Rect struct {
	w, h float64
}

type Circle struct {
	r float64
}

func (r Rect) area() float64 {
	a := r.w * r.h
	return a
}

func (r Rect) perim() float64 {
	return 2 * r.w * r.h
}

func (c Circle) area() float64 {
	a := math.Pi * (c.r * c.r)
	return a
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.r
}


func measure(g Geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
	fmt.Println("Interface")
	circle := Circle{r: 5}
	rect := Rect{w: 2, h: 3}
	measure(circle)
	measure(rect)
}