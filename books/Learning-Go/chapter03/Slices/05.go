package main

import "fmt"

func main() {
	var x []int
	fmt.Println(x == nil, x)
	y := []int{}
	fmt.Println(y == nil, y)
	z := make([]int, 2, 2)
	fmt.Println(z, len(z), cap(z))
	z = append(z, 1)
	fmt.Println(z, len(z), cap(z))
	z = append(z, 1)
	fmt.Println(z, len(z), cap(z))
	z = append(z, 1)
	fmt.Println(z, len(z), cap(z))
	z = append(z, 1)
	fmt.Println(z, len(z), cap(z))
}