package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "John",
		"address": "Indonesia",
	}

	person["title"] = "Programmer"
	fmt.Println(person)
	fmt.Println(person["name"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Learn Golang"
	book["author"] = "John"
	fmt.Println(book)
	delete(book, "author")
	fmt.Println(book)
}
