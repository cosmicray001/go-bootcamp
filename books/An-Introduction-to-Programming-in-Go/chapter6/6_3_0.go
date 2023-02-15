package main

import "fmt"

func main() {
	mp := make(map[string]int)
	mp["first"] = 1
	fmt.Println(mp["first"])
}