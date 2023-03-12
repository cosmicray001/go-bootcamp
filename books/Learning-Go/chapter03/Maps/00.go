package main

import "fmt"

func main() {
	var nilMap map[string]int
	fmt.Println(nilMap, "is nil:", nilMap == nil)
	// nilMap["idx1"] = 10 // panic
	mp := map[string]int{}
	fmt.Println(mp, "is nil:", mp == nil)
}