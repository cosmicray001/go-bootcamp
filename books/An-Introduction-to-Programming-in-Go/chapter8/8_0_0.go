package main

import "fmt"

func setValue(x int) {
	x = 0
}

func main() {
	x := 15
	setValue(x)
	fmt.Println(x)
}
