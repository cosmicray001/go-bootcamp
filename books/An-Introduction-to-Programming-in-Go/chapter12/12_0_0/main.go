package main

import (
	"fmt"

	stat "github.com/cosmicray001/go-bootcamp/books/An-Introduction-to-Programming-in-Go/chapter12/12_0_0/math"
)

func main() {
	nums := []float64{1, 2, 3, 4, 5, 6}
	avg := stat.Avg(nums)
	fmt.Println(avg)
}