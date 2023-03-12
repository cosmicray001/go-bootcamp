package main

import "fmt"

func main() {
	mp := map[string][]string{
		"Orcas": {"Fred", "Ralph", "Bijou"},
		"Lions": {"Sarah", "Peter", "Billie"},
		"Kittens": {"Waldo", "Raul", "Ze"},
	}
	fmt.Println(len(mp))
	for key, strList := range mp {
		fmt.Println(key, "-->")
		for _, name := range strList {
			fmt.Print(name, ", ")
		}
		fmt.Println()
	}
}