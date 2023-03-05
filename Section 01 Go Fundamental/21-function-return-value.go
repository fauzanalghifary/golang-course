package main

import "fmt"

func main() {
	// Call function with return value
	result := getHello("John")
	fmt.Println(result)
}

// Function with return value
func getHello(name string) string {
	return "Hello " + name
}
