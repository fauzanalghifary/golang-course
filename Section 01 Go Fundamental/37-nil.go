package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	}
	return map[string]string{
		"name": name,
	}
}

func main() {
	var person map[string]string = newMap("Pojan")

	if person == nil {
		println("person is nil")
	} else {
		fmt.Println(person)
	}
}
