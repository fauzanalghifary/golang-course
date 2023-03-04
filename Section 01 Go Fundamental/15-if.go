package main

func main() {
	name := "John"

	if name == "John" {
		println("Hello John")
	} else if name == "Doe" {
		println("Hello Doe")
	} else {
		println("Hello Unknown")
	}

	// Short statement
	if length := len(name); length > 5 {
		println("Too Long")
	} else {
		println("OK")
	}
}
