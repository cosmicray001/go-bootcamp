package main

import "fmt"

func main() {
	var x []int
	x = append(x, 1)
	fmt.Println(x)
	x = append(x, 2, 3, 4)
	fmt.Println(x)
	y := []int{9, 8, 7}
	x = append(x, y...)
	fmt.Println(x)
}