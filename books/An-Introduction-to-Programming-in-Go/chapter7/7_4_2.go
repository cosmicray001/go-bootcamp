package main

import "fmt"

func main() {
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}

func makeEvenGenerator() func() int {
	i := 0
	return func() int {
		res := i
		i += 2
		return res
	}
}