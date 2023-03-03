package main

import "fmt"

func main() {
	var num32 int32 = 1000001
	var num64 int64 = int64(num32)
	var num8 int8 = int8(num32) // error: constant 1000000 overflows int8
	fmt.Println(num32)
	fmt.Println(num64)
	fmt.Println(num8)

	name := "Pojan"
	p := name[0]
	fmt.Println(p)
	var pString string = string(p)
	fmt.Println(pString)
}
