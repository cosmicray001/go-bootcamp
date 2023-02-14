package main

import "fmt"

func main() {
	var nums [5]float64
	nums[0] = 98
	nums[1] = 99
	nums[2] = 100
	nums[3] = 77
	nums[4] = 57

	var total float64 = 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	fmt.Println(total/float64(len(nums)))
}
