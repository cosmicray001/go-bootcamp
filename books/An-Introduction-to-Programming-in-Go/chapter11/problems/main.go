package main

import (
	"fmt"

	stat "github.com/cosmicray001/go-bootcamp/books/An-Introduction-to-Programming-in-Go/chapter11/problems/maxmin"
)

// Run Commad -->  godoc -http=:6060
// to generate auto gen doc
func main() {
	nums := []float64{1, 2, 3, 4, 5, 6}
	mx := stat.Max(nums)
	mn := stat.Min(nums)

	fmt.Println(mx)
	fmt.Println(mn)
}

