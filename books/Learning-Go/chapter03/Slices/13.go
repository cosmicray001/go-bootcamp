package main

import "fmt"

func main() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4) // [1, 2, 3, 4]
	fmt.Println(x)
	y := x[:2] // [1, 2]
	z := x[2:] // [3, 4]
	fmt.Println(cap(x), cap(y), cap(z))
}