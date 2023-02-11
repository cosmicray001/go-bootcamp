package main

import "fmt"

func main() {
	ans := (true && false) || (false && true) || !(false && false)
	fmt.Println(ans)
}