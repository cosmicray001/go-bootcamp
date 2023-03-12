package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	num := copy(x[:len(x)-1], x[1:])
	fmt.Println(num)
	fmt.Println(x)
}