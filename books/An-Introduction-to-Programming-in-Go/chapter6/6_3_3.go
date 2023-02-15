package main

import "fmt"

func main() {
	mp := map[string]string {
		"first": "01",
		"second": "02",
		"third": "03",
	}
	for _, val := range mp {
		fmt.Println(val)
	}
}