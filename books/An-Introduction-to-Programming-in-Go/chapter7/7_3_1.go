package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(totalFunction(nums...))
}

func totalFunction(args ...int) int {
	total := 0
	for _, val := range args {
		total += val
	}
	return total
}
