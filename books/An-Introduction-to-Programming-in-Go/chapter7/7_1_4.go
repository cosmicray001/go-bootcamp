package main

import "fmt"

func main() {
	sum, sub := myCal(20, 10)
	fmt.Println(sum)
	fmt.Println(sub)
}

func myCal(a, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}
