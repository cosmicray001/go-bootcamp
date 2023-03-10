package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8}
	y := x[1:3]
	fmt.Println(x, y)
	fmt.Println(len(x), len(y))
	fmt.Println(cap(x), cap(y))
	y = append(y, 1000)
	fmt.Println(x, y)
}