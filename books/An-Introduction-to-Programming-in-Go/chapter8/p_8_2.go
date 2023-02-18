package main

import "fmt"

func assignFuntion(num *int) {
	*num = 1000
}

func main() {
	var myNum int
	assignFuntion(&myNum)
	fmt.Println(myNum)
}