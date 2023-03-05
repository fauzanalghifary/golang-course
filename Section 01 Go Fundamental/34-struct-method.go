package main

type Customer2 struct {
	Name, Address string
	Age           int
}

func (customer Customer2) sayHello(name string) {
	println("Hello", name, ", My name is", customer.Name)
}

func main() {
	var pojan Customer2
	pojan.Name = "Pojan"
	pojan.Address = "Jakarta"
	pojan.Age = 20
	pojan.sayHello("Alg")
}
