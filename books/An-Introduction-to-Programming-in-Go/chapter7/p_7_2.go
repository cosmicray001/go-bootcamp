package main

import "fmt"

func main() {
	if val, ok := half(5); ok {
		fmt.Println(val)
	}

}

func half(a int) (int, bool) {
	if a % 2 == 0 {
		return 1, true
	}
	return 0, false
}