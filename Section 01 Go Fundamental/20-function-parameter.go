package main

import "fmt"

func main() {
	sayHelloTo("John", "Wick")
	sayHelloTo("Ethan", "Hunt")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
