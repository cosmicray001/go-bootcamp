package main

import "fmt"
import m "github.com/cosmicray001/go-bootcamp/books/An-Introduction-to-Programming-in-Go/chapter11/math"

func main() {
	nums := []float64{1, 2, 3, 4, 5, 6}
	avg := m.Avg(nums)

	fmt.Println(avg)
}