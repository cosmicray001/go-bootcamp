package main

import "fmt"

func setValue(val *int) {
	*val = 0
}

func main() {
	num := 10
	fmt.Println(num)
	setValue(&num)

	fmt.Println(num)
}