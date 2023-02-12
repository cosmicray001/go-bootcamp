package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			fmt.Println("Divide by 2")
		} else if i % 3 == 0 {
			fmt.Println("Divide by 3")
		} else if i % 5 == 0{
			fmt.Println("Divide by 5")
		} else {
			fmt.Println("Others")
		}
	}
}