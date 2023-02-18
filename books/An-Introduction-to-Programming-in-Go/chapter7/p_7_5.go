package main

import "fmt"

func main() {
	fib := fibNum(5)
	fmt.Println(fib)
}

func fibNum(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return fibNum(n - 1) + fibNum(n - 2)
}