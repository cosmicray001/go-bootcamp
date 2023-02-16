package main

import "fmt"

func main() {
	total := totalFunction(1, 2, 3, 4)
	fmt.Println(total)
}

func totalFunction(args ...int) int {
	total := 0
	for _, val := range args {
		total += val
	}
	return total
}
