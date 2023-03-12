package main

import "fmt"

func main() {
	x := [4]int{5, 6, 7, 8}
	y := x[:2]
	z := x[2:]
	x[0] = 10
	// x = append(x, 1)
	y = append(y, 2)
	z = append(z, 3)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
