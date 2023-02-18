package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func main() {
	rectangle := Rectangle{1, 2, 3, 4}
	rectangleArea := rectangle.area()
	fmt.Println(rectangleArea)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func distance(x1, y1, x2, y2 float64) (dis float64) {
	a := x2 - x1
	b := y2 - y1
	dis = math.Sqrt(a*a + b*b)
	return
}
