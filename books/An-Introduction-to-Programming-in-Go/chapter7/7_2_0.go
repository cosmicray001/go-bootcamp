package main

import "fmt"

func main() {
	sum, sub := multipleReturnFuntion(20, 10)
	fmt.Println(sum, sub)
}

func multipleReturnFuntion(a, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}