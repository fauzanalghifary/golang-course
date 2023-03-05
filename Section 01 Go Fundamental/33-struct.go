package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	Married       bool
}

func main() {
	var pojan Customer
	pojan.Name = "Pojan"
	pojan.Address = "Jakarta"
	pojan.Age = 20
	fmt.Println(pojan)

	// or
	var pojan2 = Customer{
		Name:    "Pojan 2",
		Address: "Jakarta",
		Age:     20,
	}
	fmt.Println(pojan2)

	// or
	var pojan3 = Customer{"Pojan 3", "Jakarta", 20, false}
	fmt.Println(pojan3)
}
