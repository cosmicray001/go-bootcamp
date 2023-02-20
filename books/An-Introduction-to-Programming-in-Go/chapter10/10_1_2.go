package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanf("%s", &input)
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":-->", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Microsecond * amt)
	}
}