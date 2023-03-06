package main

import "fmt"

func main() {
	var x int = 10
	var y float64 = 30.2
	var z float64 = y + float64(x)
	var d int = x + int(y)
	fmt.Println(z, d)
}
