package main

import "fmt"

func main() {
	var mp map[string]int = make(map[string]int)
	mp["first"] = 1
	mp["second"] = 2
	mp["third"] = 3

	if val, ok := mp["first"]; ok {
		fmt.Println(val)
	}
}