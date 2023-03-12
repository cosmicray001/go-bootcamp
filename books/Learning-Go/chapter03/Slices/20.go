package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	d := [4]int{5, 6, 7, 8}
	y := make([]int, 2)
	copy(y, d[:])
	copy(d[:], x)
	fmt.Println(y)
	fmt.Println(d)

}