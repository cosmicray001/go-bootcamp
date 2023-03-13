package main

import "fmt"

func main() {
	mp := make(map[string]int, 2)
	mp["hello"] = 5
	mp["world"] = 10
	fmt.Println(mp)
	delete(mp, "hello")
	fmt.Println(mp)
	delete(mp, "notFound")
	fmt.Println(mp)
}