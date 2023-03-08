package main

import "fmt"

func main() {
	var arr = [3]int{1, 2, 3}
	var arr1 = [...]int{1, 2, 3}
	var sli = []int{1, 2, 3}
	fmt.Println(arr, arr1, sli)
	var x_arr = [12]int{1, 5: 4, 6, 10: 100, 15}
	var x_sli = []int{1, 5:4, 6, 10:100, 15}
	fmt.Println(x_arr)
	fmt.Println(x_sli)
}