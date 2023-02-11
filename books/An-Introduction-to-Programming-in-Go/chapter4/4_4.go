package main

import "fmt"

func main() {
	var (
		num int = 10
		msg string
		flag bool
	)
	msg = "Hello world"
	flag = true
	fmt.Printf("%d %s %v\n", num, msg, flag)

}