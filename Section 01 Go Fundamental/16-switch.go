package main

func main() {
	names := "John"

	switch names {
	case "John":
		println("Hello John")
	case "Doe":
		println("Hello Doe")
	default:
		println("Hello Unknown")
	}

	//	short statement
	switch length := len(names); length > 5 {
	case true:
		println("Too long")
	case false:
		println("Too short")
	}
}
