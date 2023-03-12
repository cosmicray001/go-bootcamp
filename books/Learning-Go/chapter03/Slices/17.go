package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 7}
	y := make([]int, 20, 100)
	num := copy(y, x)
	fmt.Println(num)
	fmt.Println(y)
}