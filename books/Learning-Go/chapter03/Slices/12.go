package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8}
	a := x[:1]
	b := x[1:2]
	c := x[2:3]
	d := x[3:4]
// 	Whenever you take a slice from another slice, the subslice’s capacity
// is set to the capacity of the original slice, minus the offset of the subslice within the
// original slice.
	fmt.Println(cap(x), cap(a), cap(b), cap(c), cap(d))
}