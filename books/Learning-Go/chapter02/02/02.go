package main

import "fmt"

func main() {
	var a, b int
	a = 20
	b = 10
	sum := a + b
	sub := a - b
	mul := a * b
	div := a / b
	rem := a % b
	fmt.Println(sum, sub, mul, div, rem)
}