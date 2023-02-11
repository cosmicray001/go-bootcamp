package main

import "fmt"

func main() {
	var ft float64
	fmt.Scanf("%f", &ft)
	meter := (ft * .3048)
	fmt.Println(meter)
}