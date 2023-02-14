package main

import "fmt"

func main() {
	// var nums [5]float64 = [5]float64{98, 93, 77, 82, 83}
	nums := [5]float64{
		98,
		93,
		77,
		82,
		83,
		// 100,
	}
	var total float64 = 0
	for _, val := range nums {
		total += val
	}
	fmt.Println(total/float64(len(nums)))
}