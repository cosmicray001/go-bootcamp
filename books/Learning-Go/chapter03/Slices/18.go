package main

import "fmt"

func main() {
	x := [4]int{1, 2, 3, 4} // array or slice both works
	y := make([]int, 2)
	num := copy(y, x[2:])
	fmt.Println(num, y)
}