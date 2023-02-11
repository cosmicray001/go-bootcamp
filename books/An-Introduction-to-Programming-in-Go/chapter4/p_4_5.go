package main

import "fmt"

func main() {
	var f float64
	fmt.Scanf("%f", &f)
	c := (f - 32) * (5.0/9.0)
	fmt.Println(c)
}