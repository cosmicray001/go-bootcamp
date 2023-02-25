package main

import "fmt"

func main() {
	// string to byte of slice
	arr := []byte("test")
	fmt.Println(arr)
	// byte og slice to string
	str := string([]byte{'a', 'b', 'c', 'd'})
	fmt.Println(str)

}