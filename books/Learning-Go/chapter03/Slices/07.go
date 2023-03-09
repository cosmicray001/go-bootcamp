package main

import "fmt"

func main() {
	x := []int{10, 20, 30, 40}
	y := x[:2]
	z := x[1:]
	x[0] = 0
	y[1] = 0
	z[1] = 0
	x[3] = 0
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}