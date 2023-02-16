package main

import "fmt"

func main() {
	nums := []float64{98, 93, 77, 82, 83}
	avgValue := average(nums)
	fmt.Println(avgValue)
}

func average(nums []float64) float64 {
	total := 0.0
	for _, val := range nums {
		total += val
	}
	return total / float64(len(nums))
}
