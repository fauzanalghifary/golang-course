package main

import "fmt"

func main() {
	type KTPNumber string
	var myKTP KTPNumber = "1234567890"
	fmt.Println(myKTP)
}
