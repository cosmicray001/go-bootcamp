package main

import "fmt"

func main() {
	fmt.Println(fnc1())
}

func fnc1() int {
	return fnc2()
}

func fnc2() int {
	return 2
}
