package main

import "fmt"

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
}