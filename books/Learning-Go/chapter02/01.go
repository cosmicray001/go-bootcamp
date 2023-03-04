package main

import "fmt"

func main() {
	var num1, num2 int8 = -128, 127
	fmt.Println(num1, num2)
	var num3, num4 int16 = -32768, 32767
	fmt.Println(num3, num4)
	var num5, num6 int32 = -2147483648, 2147483647
	fmt.Println(num5, num6)
	var num7, num8 = -9223372036854775808, 9223372036854775807
	fmt.Println(num7, num8)
	var num9, num10 uint8 = 0, 255
	fmt.Println(num9, num10)
	var num11, num12 uint16 = 0, 65535
	fmt.Println(num11, num12)
	var num13, num14 uint32 = 0, 4294967295
	fmt.Println(num13, num14)
	var num15, num16 uint64 = 0, 18446744073709551615
	fmt.Println(num15, num16)
}