package main

import "fmt"

func main() {
	// Closure is a function that can access the variable that declared outside the function body.
	// The function can access and assign a value to the variable that declared in the outer function.
	// The function can access the variable even after the outer function has returned.

	name := "John"
	counter := 0
	increment := func() {
		counter++
		fmt.Println(name)
	}

	increment()
	increment()
	fmt.Println(counter)
}
