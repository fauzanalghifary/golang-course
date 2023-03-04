package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "John"
	names[1] = "Wick"
	names[2] = "Bourne"
	//names[3] = "Bond" // error

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		80,
		70,
	}
	fmt.Println(values)
	fmt.Println(len(values))
}
