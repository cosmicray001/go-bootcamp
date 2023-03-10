package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := x[:2] // [1, 2]
	fmt.Println(cap(x), cap(y)) // 4, original array cap - offset
	y = append(y, 30) // [1, 2, 30] x = [1, 2, 30, 4] --> y[2] = 30 --> x[2] = 30
	fmt.Println(x, y) // [1, 2, 30, 4] [1, 2, 30]
}