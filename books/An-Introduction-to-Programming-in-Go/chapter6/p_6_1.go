package main

import "fmt"

func main() {
	var nums [10]int
	for i := 0; i < len(nums); i++ {
		nums[i] = i + 1
	}
	fmt.Println(nums[3])
}