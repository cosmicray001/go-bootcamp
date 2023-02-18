package main

import "fmt"

func one(a *int) {
	*a = 1
}

func main() {
	a := new(int)
	fmt.Println(a)
	one(a)
	fmt.Println(*a)
}