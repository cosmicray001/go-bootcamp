package main

import "fmt"

func main() {
	factorialValue := factorial(5)
	fmt.Println(factorialValue)
}

func factorial(a int) int {
	if a < 1 {
		return 1
	}
	return a * factorial(a - 1)
}