package main

import "fmt"

func main() {
	oddNum := makeOddGenerator()
	for i := 0; i < 10; i++ {
		fmt.Println(oddNum())
	}
}

func makeOddGenerator() func() (ret int) {
	i := 1
	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}