package main

import "fmt"

func main() {
	// create map using make func when you know the size
	ages := make(map[string]int, 2)
	fmt.Println(len(ages))
	ages["alex"] = 20
	ages["Gorge"] = 30
	ages["Extra"] = 30 // also use make with 0 len when you don't know the size
	fmt.Println(ages)
	fmt.Println(len(ages))
}