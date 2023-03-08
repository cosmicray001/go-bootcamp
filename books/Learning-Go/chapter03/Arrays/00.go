package main

import "fmt"

func main() {
	var num [3]int
	fmt.Println(num)
	var num1 = [3]int{1, 2, 3}
	fmt.Println(num1)
	var x = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(x)
	var a = [...]int{1, 2, 3}
	var b = [3]int{1, 2, 3}
	fmt.Println(a, b)
	fmt.Println(a == b)
	fmt.Println(len(x))
}