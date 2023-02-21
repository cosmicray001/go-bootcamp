package main

import (
	"fmt"
)


func pinger(c chan<- string) {
	c <- "ping"
}

func printer(c <-chan string) {
	msg := <- c
	fmt.Println(msg)
}

func main() {
	c := make(chan string)
	go pinger(c)
	go printer(c)
	var input string
	fmt.Scan(&input)
}