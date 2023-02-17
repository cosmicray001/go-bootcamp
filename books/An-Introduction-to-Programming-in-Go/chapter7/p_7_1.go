package main

import "fmt"

func main() {
	nums := []int {1, 2, 3}
	sum := sumFunction(nums...)
	fmt.Println(sum)
}

func sumFunction(args ...int) int {
	sum := 0
	for _, val := range args {
		sum += val
	}
	return sum
}