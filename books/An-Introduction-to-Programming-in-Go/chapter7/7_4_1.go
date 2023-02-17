package main

import "fmt"

func main() {
	val := 0

	increment := func() int {
		val++
		return val
	}

	fmt.Println(increment())
	fmt.Println(increment())
}