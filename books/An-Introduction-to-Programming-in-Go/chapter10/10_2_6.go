package main

import "fmt"

func main() {
	c := make(chan string, 2)
	c <- "md"
	c <- "samiul"
	// c <- "islam"
	fmt.Println(<- c)
	fmt.Println(<- c)
	// fmt.Println(<- c)
}