package main

import "fmt"

func main() {
	var a rune = 'x'
	var s string = string(a)
	fmt.Println(s)

	var b byte = 'y'
	var s2 string = string(b)
	fmt.Println(s2)
}