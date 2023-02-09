package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// number
	fmt.Println("One = ", 1)
	fmt.Println("Two = ", 2)
	fmt.Println("Three point fiove = ", 3.5)
	
	// boolean
	fmt.Println("True = ", true)
	fmt.Println("False = ", false)
	
	// string
	fmt.Println("Fauzan")
	fmt.Println("Fauzan Al G")
	fmt.Println(len("Fauzan Al G"))
	fmt.Println("Fauzan Al G"[1])
	
	// variable
	var name string
	name = "Fauzan Al G"
	fmt.Println(name)

	name = "Fauzan Al G"
	fmt.Println(name)

	var age = 20
	fmt.Println(age)

	newAge := 21
	fmt.Println(newAge)

	var (
		firstName = "Fauzan"
		lastName = "Al G"
	)
	fmt.Println(firstName, lastName)

	// constant

	const pi = 3.14
	fmt.Println(pi)
	
	const (
		pi1 = 3.14
		pi2 = 3.15
	)
	fmt.Println(pi1, pi2)

}