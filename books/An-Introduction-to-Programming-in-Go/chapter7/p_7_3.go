package main

import "fmt"

func main() {
	nums := []int{100, 1, 2, 3, 4, 5, 6, 7}
	mx := fincMax(nums...)
	fmt.Println(mx)
}

func fincMax(args ...int) int {
	if len(args) == 0 {
		return -1
	}
	mx := args[0]
	for i := 1; i < len(args); i++ {
		mx = max(mx, args[i])
	}
	return mx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
