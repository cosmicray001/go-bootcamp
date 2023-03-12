package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	y := x[:2]
	y = append(y, 9, 8, 7, 6)
	fmt.Println(x, y)
}