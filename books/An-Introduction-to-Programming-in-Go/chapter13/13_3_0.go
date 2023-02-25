package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		return
	}
	defer file.Close()

	// get file size
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// read file
	bs := make([]byte, stat.Size())
	ok, err := file.Read(bs)
	if err != nil {
		return
	}
	str := string(ok)
	fmt.Println(str)
}