package main

import "fmt"

func main() {
	mp := map[string]int{
		"hello": 5,
		"world": 1,
	}

	if v, ok := mp["hello"]; ok {
		fmt.Println(v, ok)
	}

	if v, ok := mp["world"]; ok {
		fmt.Println(v, ok)
	}

	if v, ok := mp["empty"]; !ok {
		fmt.Println(v, ok)
	}
}