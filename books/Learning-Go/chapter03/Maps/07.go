package main

import "fmt"

func main() {
	nums := []int{5,10,2,5,8,7,3,9,1,2,10}
	mp := map[int]struct{}{}
	// mp := make(map[int]struct{}, 0)

	for _, val := range nums {
		mp[val] = struct{}{}
	}

	if _, ok := mp[10]; ok {
		fmt.Println("10 exists in map")
	}
}