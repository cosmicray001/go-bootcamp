package main

import (
	"fmt"
)

func main() {
	var s string = "Hello 👻"
	var x []byte = []byte(s)
	fmt.Println(x)
	var y []rune = []rune(s)
	fmt.Println(y)
}