package main

func main() {
	// Call function with return value
	result, length := getHelloTo("John")
	println(result, length)
}

// Function with return value
func getHelloTo(name string) (string, int) {
	return "Hello " + name, len(name)
}
