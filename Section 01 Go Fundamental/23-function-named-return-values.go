package main

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "John"
	middleName = "Paul"
	lastName = "Doe"
	return
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	println(firstName, middleName, lastName)
}
