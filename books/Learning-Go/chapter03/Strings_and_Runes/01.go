package main

import "fmt"

func main() {
	var s string = "Hello there"
	var s1 string = s[3:7]
	var s2 string = s[:5]
	var s3 string = s[4:]
	fmt.Printf("%s\n%s\n%s\n%s\n", s, s1, s2, s3)
}