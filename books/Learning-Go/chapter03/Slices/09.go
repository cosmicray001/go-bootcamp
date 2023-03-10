package main

import "fmt"

func main() {
	x := make([]int, 0)
	x = append(x, 1, 2, 3, 4)
	y := x[:2]
	z := x[2:]
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, 30, 40, 50)
	fmt.Println(x)
	x = append(x, 60)
	z = append(z, 70)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}