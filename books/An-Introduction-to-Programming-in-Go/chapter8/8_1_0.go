package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}
func main() {
	x := 5
	zero(&x)
	y := &x
	*y = 10
	fmt.Println(y) // x is 0
}
