package main

import "fmt"

func main() {
	mp := map[string]map[string] string {
		"BS0779": map[string]string {
			"Name": "Md Samiul Islam",
			"github": "samiul-bs23",
		},
	}
	for _, val := range mp {
		for key, info := range val {
			fmt.Println(key, ": ",info)
		}
	}
	// fmt.Println(mp)
}