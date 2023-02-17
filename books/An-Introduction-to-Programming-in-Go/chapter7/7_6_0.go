package main

import "fmt"

func main() {
	defer second()
	first()
}

func first() {
	fmt.Println("This is the first function")
}

func second() {
	fmt.Println("This is the second function")
}
