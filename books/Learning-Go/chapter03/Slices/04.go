package main

import "fmt"

func main() {
	x := make([]int, 5)
	fmt.Println(x)
	x = append(x, 10)
	fmt.Println(x)
	y := make([]int, 5, 10)
	fmt.Println(y, len(y), cap(y))
	z := make([]int, 0, 10)
	fmt.Println(z)
	z = append(z, 5,6,7,8)
	fmt.Println(z)
}