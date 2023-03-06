package main

import "fmt"

func main() {
	a := 30
	b := 10
	fmt.Println(a << 1)
	fmt.Println(a >> 1)
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)

	var num int = 1
	fmt.Println(num &^ 0)
}